package dynamic

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/waku-org/go-waku/waku/v2/protocol/rln/contracts"
	"github.com/waku-org/go-waku/waku/v2/protocol/rln/group_manager"
	"github.com/waku-org/go-waku/waku/v2/protocol/rln/keystore"
	"github.com/waku-org/go-zerokit-rln/rln"
	om "github.com/wk8/go-ordered-map"
	"go.uber.org/zap"
)

var RLNAppInfo = keystore.AppInfo{
	Application:   "waku-rln-relay",
	AppIdentifier: "01234567890abcdef",
	Version:       "0.1",
}

type DynamicGroupManager struct {
	rln *rln.RLN
	log *zap.Logger

	cancel context.CancelFunc
	wg     sync.WaitGroup

	identityCredential *rln.IdentityCredential
	membershipIndex    *rln.MembershipIndex

	membershipContractAddress common.Address
	membershipGroupIndex      uint
	ethClientAddress          string
	ethClient                 *ethclient.Client

	lastBlockProcessed uint64

	eventHandler RegistrationEventHandler

	chainId     *big.Int
	rlnContract *contracts.RLN

	saveKeystore     bool
	keystorePath     string
	keystorePassword string
	keystoreIndex    uint

	rootTracker *group_manager.MerkleRootTracker
}

func handler(gm *DynamicGroupManager, events []*contracts.RLNMemberRegistered) error {
	toRemoveTable := om.New()
	toInsertTable := om.New()

	lastBlockProcessed := gm.lastBlockProcessed
	for _, event := range events {
		if event.Raw.Removed {
			var indexes []uint
			i_idx, ok := toRemoveTable.Get(event.Raw.BlockNumber)
			if ok {
				indexes = i_idx.([]uint)
			}
			indexes = append(indexes, uint(event.Index.Uint64()))
			toRemoveTable.Set(event.Raw.BlockNumber, indexes)
		} else {
			var eventsPerBlock []*contracts.RLNMemberRegistered
			i_evt, ok := toInsertTable.Get(event.Raw.BlockNumber)
			if ok {
				eventsPerBlock = i_evt.([]*contracts.RLNMemberRegistered)
			}
			eventsPerBlock = append(eventsPerBlock, event)
			toInsertTable.Set(event.Raw.BlockNumber, eventsPerBlock)

			if event.Raw.BlockNumber > lastBlockProcessed {
				lastBlockProcessed = event.Raw.BlockNumber
			}
		}
	}

	err := gm.RemoveMembers(toRemoveTable)
	if err != nil {
		return err
	}

	err = gm.InsertMembers(toInsertTable)
	if err != nil {
		return err
	}

	gm.lastBlockProcessed = lastBlockProcessed
	err = gm.SetMetadata(RLNMetadata{
		LastProcessedBlock: gm.lastBlockProcessed,
	})
	if err != nil {
		// this is not a fatal error, hence we don't raise an exception
		gm.log.Warn("failed to persist rln metadata", zap.Error(err))
	} else {
		gm.log.Debug("rln metadata persisted", zap.Uint64("lastProcessedBlock", gm.lastBlockProcessed))
	}

	return nil
}

type RegistrationHandler = func(tx *types.Transaction)

func NewDynamicGroupManager(
	ethClientAddr string,
	memContractAddr common.Address,
	membershipGroupIndex uint,
	keystorePath string,
	keystorePassword string,
	keystoreIndex uint,
	saveKeystore bool,
	log *zap.Logger,
) (*DynamicGroupManager, error) {
	log = log.Named("rln-dynamic")

	path := keystorePath
	if path == "" {
		log.Warn("keystore: no credentials path set, using default path", zap.String("path", keystore.RLN_CREDENTIALS_FILENAME))
		path = keystore.RLN_CREDENTIALS_FILENAME
	}

	password := keystorePassword
	if password == "" {
		log.Warn("keystore: no credentials password set, using default password", zap.String("password", keystore.RLN_CREDENTIALS_PASSWORD))
		password = keystore.RLN_CREDENTIALS_PASSWORD
	}

	return &DynamicGroupManager{
		membershipGroupIndex:      membershipGroupIndex,
		membershipContractAddress: memContractAddr,
		ethClientAddress:          ethClientAddr,
		eventHandler:              handler,
		saveKeystore:              saveKeystore,
		keystorePath:              path,
		keystorePassword:          password,
		keystoreIndex:             keystoreIndex,
		log:                       log,
	}, nil
}

func (gm *DynamicGroupManager) getMembershipFee(ctx context.Context) (*big.Int, error) {
	return gm.rlnContract.MEMBERSHIPDEPOSIT(&bind.CallOpts{Context: ctx})
}

func (gm *DynamicGroupManager) Start(ctx context.Context, rlnInstance *rln.RLN, rootTracker *group_manager.MerkleRootTracker) error {
	if gm.cancel != nil {
		return errors.New("already started")
	}

	ctx, cancel := context.WithCancel(ctx)
	gm.cancel = cancel

	gm.log.Info("mounting rln-relay in on-chain/dynamic mode")

	backend, err := ethclient.Dial(gm.ethClientAddress)
	if err != nil {
		return err
	}
	gm.ethClient = backend

	gm.rln = rlnInstance
	gm.rootTracker = rootTracker

	gm.chainId, err = backend.ChainID(ctx)
	if err != nil {
		return err
	}

	gm.rlnContract, err = contracts.NewRLN(gm.membershipContractAddress, backend)
	if err != nil {
		return err
	}

	// check if the contract exists by calling a static function
	_, err = gm.getMembershipFee(ctx)
	if err != nil {
		return err
	}

	if gm.identityCredential == nil && gm.keystorePassword != "" && gm.keystorePath != "" {
		credentials, err := keystore.GetMembershipCredentials(gm.log,
			gm.keystorePath,
			gm.keystorePassword,
			RLNAppInfo,
			nil,
			[]keystore.MembershipContract{{
				ChainId: fmt.Sprintf("0x%X", gm.chainId),
				Address: gm.membershipContractAddress.Hex(),
			}})
		if err != nil {
			return err
		}

		if len(credentials) != 0 {
			if int(gm.keystoreIndex) <= len(credentials)-1 {
				credential := credentials[gm.keystoreIndex]
				gm.identityCredential = credential.IdentityCredential
				if int(gm.membershipGroupIndex) <= len(credential.MembershipGroups)-1 {
					gm.membershipIndex = &credential.MembershipGroups[gm.membershipGroupIndex].TreeIndex
				} else {
					return errors.New("invalid membership group index")
				}
			} else {
				return errors.New("invalid keystore index")
			}
		}
	}

	if gm.identityCredential == nil || gm.membershipIndex == nil {
		return errors.New("no credentials available")
	}

	if err = gm.HandleGroupUpdates(ctx, gm.eventHandler); err != nil {
		return err
	}

	return nil
}

func (gm *DynamicGroupManager) InsertMembers(toInsert *om.OrderedMap) error {
	for pair := toInsert.Oldest(); pair != nil; pair = pair.Next() {
		events := pair.Value.([]*contracts.RLNMemberRegistered) // TODO: should these be sortered by index? we assume all members arrive in order
		var idCommitments []rln.IDCommitment
		var oldestIndexInBlock *big.Int
		for _, evt := range events {
			if oldestIndexInBlock == nil {
				oldestIndexInBlock = evt.Index
			}
			idCommitments = append(idCommitments, rln.BigIntToBytes32(evt.IdCommitment))
		}

		if len(idCommitments) == 0 {
			continue
		}

		// TODO: should we track indexes to identify missing?
		startIndex := rln.MembershipIndex(uint(oldestIndexInBlock.Int64()))
		err := gm.rln.InsertMembers(startIndex, idCommitments)
		if err != nil {
			gm.log.Error("inserting members into merkletree", zap.Error(err))
			return err
		}

		_, err = gm.rootTracker.UpdateLatestRoot(pair.Key.(uint64))
		if err != nil {
			return err
		}
	}
	return nil
}

func (gm *DynamicGroupManager) RemoveMembers(toRemove *om.OrderedMap) error {
	for pair := toRemove.Newest(); pair != nil; pair = pair.Prev() {
		memberIndexes := pair.Value.([]uint)
		err := gm.rln.DeleteMembers(memberIndexes)
		if err != nil {
			gm.log.Error("deleting members", zap.Error(err))
			return err
		}
		gm.rootTracker.Backfill(pair.Key.(uint64))
	}

	return nil
}

func (gm *DynamicGroupManager) IdentityCredentials() (rln.IdentityCredential, error) {
	if gm.identityCredential == nil {
		return rln.IdentityCredential{}, errors.New("identity credential has not been setup")
	}

	return *gm.identityCredential, nil
}

func (gm *DynamicGroupManager) MembershipIndex() (rln.MembershipIndex, error) {
	if gm.membershipIndex == nil {
		return 0, errors.New("membership index has not been setup")
	}

	return *gm.membershipIndex, nil
}

// Stop stops all go-routines, eth client and closes the rln database
func (gm *DynamicGroupManager) Stop() error {
	if gm.cancel == nil {
		return nil
	}

	gm.cancel()

	err := gm.rln.Flush()
	if err != nil {
		return err
	}
	gm.ethClient.Close()

	gm.wg.Wait()

	return nil
}
