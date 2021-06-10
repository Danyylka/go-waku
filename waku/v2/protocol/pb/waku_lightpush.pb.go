// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: waku_lightpush.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PushRequest struct {
	PubsubTopic          string       `protobuf:"bytes,1,opt,name=pubsub_topic,json=pubsubTopic,proto3" json:"pubsub_topic,omitempty"`
	Message              *WakuMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PushRequest) Reset()         { *m = PushRequest{} }
func (m *PushRequest) String() string { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()    {}
func (*PushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0edfa2f8ec212684, []int{0}
}
func (m *PushRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRequest.Merge(m, src)
}
func (m *PushRequest) XXX_Size() int {
	return m.Size()
}
func (m *PushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushRequest proto.InternalMessageInfo

func (m *PushRequest) GetPubsubTopic() string {
	if m != nil {
		return m.PubsubTopic
	}
	return ""
}

func (m *PushRequest) GetMessage() *WakuMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

type PushResponse struct {
	IsSuccess bool `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	// Error messages, etc
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushResponse) Reset()         { *m = PushResponse{} }
func (m *PushResponse) String() string { return proto.CompactTextString(m) }
func (*PushResponse) ProtoMessage()    {}
func (*PushResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0edfa2f8ec212684, []int{1}
}
func (m *PushResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushResponse.Merge(m, src)
}
func (m *PushResponse) XXX_Size() int {
	return m.Size()
}
func (m *PushResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushResponse proto.InternalMessageInfo

func (m *PushResponse) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *PushResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type PushRPC struct {
	RequestId            string        `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Query                *PushRequest  `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Response             *PushResponse `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PushRPC) Reset()         { *m = PushRPC{} }
func (m *PushRPC) String() string { return proto.CompactTextString(m) }
func (*PushRPC) ProtoMessage()    {}
func (*PushRPC) Descriptor() ([]byte, []int) {
	return fileDescriptor_0edfa2f8ec212684, []int{2}
}
func (m *PushRPC) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushRPC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushRPC.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushRPC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRPC.Merge(m, src)
}
func (m *PushRPC) XXX_Size() int {
	return m.Size()
}
func (m *PushRPC) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRPC.DiscardUnknown(m)
}

var xxx_messageInfo_PushRPC proto.InternalMessageInfo

func (m *PushRPC) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *PushRPC) GetQuery() *PushRequest {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *PushRPC) GetResponse() *PushResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*PushRequest)(nil), "pb.PushRequest")
	proto.RegisterType((*PushResponse)(nil), "pb.PushResponse")
	proto.RegisterType((*PushRPC)(nil), "pb.PushRPC")
}

func init() { proto.RegisterFile("waku_lightpush.proto", fileDescriptor_0edfa2f8ec212684) }

var fileDescriptor_0edfa2f8ec212684 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x4f, 0xcc, 0x2e,
	0x8d, 0xcf, 0xc9, 0x4c, 0xcf, 0x28, 0x29, 0x28, 0x2d, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x2a, 0x48, 0x92, 0x12, 0x02, 0xcb, 0xe4, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x42,
	0xc4, 0x95, 0xa2, 0xb9, 0xb8, 0x03, 0x4a, 0x8b, 0x33, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x84, 0x14, 0xb9, 0x78, 0x0a, 0x4a, 0x93, 0x8a, 0x4b, 0x93, 0xe2, 0x4b, 0xf2, 0x0b, 0x32, 0x93,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xb8, 0x21, 0x62, 0x21, 0x20, 0x21, 0x21, 0x4d, 0x2e,
	0x76, 0xa8, 0x11, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xfc, 0x7a, 0x05, 0x49, 0x7a, 0xe1,
	0x89, 0xd9, 0xa5, 0xbe, 0x10, 0xe1, 0x20, 0x98, 0xbc, 0x92, 0x23, 0x17, 0x0f, 0xc4, 0xf0, 0xe2,
	0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x59, 0x2e, 0xae, 0xcc, 0xe2, 0xf8, 0xe2, 0xd2, 0xe4, 0xe4,
	0xd4, 0xe2, 0x62, 0xb0, 0xd9, 0x1c, 0x41, 0x9c, 0x99, 0xc5, 0xc1, 0x10, 0x01, 0x21, 0x21, 0x2e,
	0x96, 0xcc, 0xbc, 0xb4, 0x7c, 0xb0, 0xb1, 0x9c, 0x41, 0x60, 0xb6, 0x52, 0x2d, 0x17, 0x3b, 0xd8,
	0x88, 0x00, 0x67, 0x90, 0xee, 0x22, 0x88, 0x33, 0xe3, 0x33, 0x53, 0xa0, 0x2e, 0xe3, 0x84, 0x8a,
	0x78, 0xa6, 0x08, 0xa9, 0x72, 0xb1, 0x16, 0x96, 0xa6, 0x16, 0x55, 0x22, 0xbb, 0x0a, 0xc9, 0x6b,
	0x41, 0x10, 0x59, 0x21, 0x1d, 0x2e, 0x8e, 0x22, 0xa8, 0x7b, 0x24, 0x98, 0xc1, 0x2a, 0x05, 0x10,
	0x2a, 0x21, 0xe2, 0x41, 0x70, 0x15, 0x4e, 0x02, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7,
	0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0xe0, 0x70, 0x33, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x76, 0x20, 0x2e, 0xed, 0x67, 0x01, 0x00, 0x00,
}

func (m *PushRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Message != nil {
		{
			size, err := m.Message.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWakuLightpush(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.PubsubTopic) > 0 {
		i -= len(m.PubsubTopic)
		copy(dAtA[i:], m.PubsubTopic)
		i = encodeVarintWakuLightpush(dAtA, i, uint64(len(m.PubsubTopic)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PushResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintWakuLightpush(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x12
	}
	if m.IsSuccess {
		i--
		if m.IsSuccess {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PushRPC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushRPC) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushRPC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Response != nil {
		{
			size, err := m.Response.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWakuLightpush(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Query != nil {
		{
			size, err := m.Query.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWakuLightpush(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RequestId) > 0 {
		i -= len(m.RequestId)
		copy(dAtA[i:], m.RequestId)
		i = encodeVarintWakuLightpush(dAtA, i, uint64(len(m.RequestId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWakuLightpush(dAtA []byte, offset int, v uint64) int {
	offset -= sovWakuLightpush(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PushRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PubsubTopic)
	if l > 0 {
		n += 1 + l + sovWakuLightpush(uint64(l))
	}
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovWakuLightpush(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PushResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsSuccess {
		n += 2
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovWakuLightpush(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PushRPC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestId)
	if l > 0 {
		n += 1 + l + sovWakuLightpush(uint64(l))
	}
	if m.Query != nil {
		l = m.Query.Size()
		n += 1 + l + sovWakuLightpush(uint64(l))
	}
	if m.Response != nil {
		l = m.Response.Size()
		n += 1 + l + sovWakuLightpush(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWakuLightpush(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWakuLightpush(x uint64) (n int) {
	return sovWakuLightpush(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PushRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWakuLightpush
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PushRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubsubTopic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWakuLightpush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubsubTopic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWakuLightpush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &WakuMessage{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWakuLightpush(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PushResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWakuLightpush
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PushResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSuccess", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWakuLightpush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsSuccess = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWakuLightpush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWakuLightpush(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PushRPC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWakuLightpush
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PushRPC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushRPC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWakuLightpush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWakuLightpush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Query == nil {
				m.Query = &PushRequest{}
			}
			if err := m.Query.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWakuLightpush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Response == nil {
				m.Response = &PushResponse{}
			}
			if err := m.Response.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWakuLightpush(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWakuLightpush
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWakuLightpush(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWakuLightpush
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWakuLightpush
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWakuLightpush
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthWakuLightpush
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWakuLightpush
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWakuLightpush
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWakuLightpush        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWakuLightpush          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWakuLightpush = fmt.Errorf("proto: unexpected end of group")
)