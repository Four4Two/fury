// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/liquid/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgMintDerivative defines the Msg/MintDerivative request type.
type MsgMintDerivative struct {
	// sender is the owner of the delegation to be converted
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// validator is the validator of the delegation to be converted
	Validator string `protobuf:"bytes,2,opt,name=validator,proto3" json:"validator,omitempty"`
	// amount is the quantity of staked assets to be converted
	Amount types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgMintDerivative) Reset()         { *m = MsgMintDerivative{} }
func (m *MsgMintDerivative) String() string { return proto.CompactTextString(m) }
func (*MsgMintDerivative) ProtoMessage()    {}
func (*MsgMintDerivative) Descriptor() ([]byte, []int) {
	return fileDescriptor_928a8dc767d67df5, []int{0}
}
func (m *MsgMintDerivative) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintDerivative) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintDerivative.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintDerivative) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintDerivative.Merge(m, src)
}
func (m *MsgMintDerivative) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintDerivative) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintDerivative.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintDerivative proto.InternalMessageInfo

func (m *MsgMintDerivative) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgMintDerivative) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *MsgMintDerivative) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

// MsgMintDerivativeResponse defines the Msg/MintDerivative response type.
type MsgMintDerivativeResponse struct {
	// received is the amount of staking derivative minted and sent to the sender
	Received types.Coin `protobuf:"bytes,1,opt,name=received,proto3" json:"received"`
}

func (m *MsgMintDerivativeResponse) Reset()         { *m = MsgMintDerivativeResponse{} }
func (m *MsgMintDerivativeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMintDerivativeResponse) ProtoMessage()    {}
func (*MsgMintDerivativeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_928a8dc767d67df5, []int{1}
}
func (m *MsgMintDerivativeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintDerivativeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintDerivativeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintDerivativeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintDerivativeResponse.Merge(m, src)
}
func (m *MsgMintDerivativeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintDerivativeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintDerivativeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintDerivativeResponse proto.InternalMessageInfo

func (m *MsgMintDerivativeResponse) GetReceived() types.Coin {
	if m != nil {
		return m.Received
	}
	return types.Coin{}
}

// MsgBurnDerivative defines the Msg/BurnDerivative request type.
type MsgBurnDerivative struct {
	// sender is the owner of the derivatives to be converted
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// validator is the validator of the derivatives to be converted
	Validator string `protobuf:"bytes,2,opt,name=validator,proto3" json:"validator,omitempty"`
	// amount is the quantity of derivatives to be converted
	Amount types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgBurnDerivative) Reset()         { *m = MsgBurnDerivative{} }
func (m *MsgBurnDerivative) String() string { return proto.CompactTextString(m) }
func (*MsgBurnDerivative) ProtoMessage()    {}
func (*MsgBurnDerivative) Descriptor() ([]byte, []int) {
	return fileDescriptor_928a8dc767d67df5, []int{2}
}
func (m *MsgBurnDerivative) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurnDerivative) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurnDerivative.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurnDerivative) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurnDerivative.Merge(m, src)
}
func (m *MsgBurnDerivative) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurnDerivative) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurnDerivative.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurnDerivative proto.InternalMessageInfo

func (m *MsgBurnDerivative) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgBurnDerivative) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *MsgBurnDerivative) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

// MsgBurnDerivativeResponse defines the Msg/BurnDerivative response type.
type MsgBurnDerivativeResponse struct {
	// received is the number of delegation shares sent to the sender
	Received github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=received,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"received"`
}

func (m *MsgBurnDerivativeResponse) Reset()         { *m = MsgBurnDerivativeResponse{} }
func (m *MsgBurnDerivativeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBurnDerivativeResponse) ProtoMessage()    {}
func (*MsgBurnDerivativeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_928a8dc767d67df5, []int{3}
}
func (m *MsgBurnDerivativeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurnDerivativeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurnDerivativeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurnDerivativeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurnDerivativeResponse.Merge(m, src)
}
func (m *MsgBurnDerivativeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurnDerivativeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurnDerivativeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurnDerivativeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgMintDerivative)(nil), "fury.liquid.v1beta1.MsgMintDerivative")
	proto.RegisterType((*MsgMintDerivativeResponse)(nil), "fury.liquid.v1beta1.MsgMintDerivativeResponse")
	proto.RegisterType((*MsgBurnDerivative)(nil), "fury.liquid.v1beta1.MsgBurnDerivative")
	proto.RegisterType((*MsgBurnDerivativeResponse)(nil), "fury.liquid.v1beta1.MsgBurnDerivativeResponse")
}

func init() { proto.RegisterFile("fury/liquid/v1beta1/tx.proto", fileDescriptor_928a8dc767d67df5) }

var fileDescriptor_928a8dc767d67df5 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xcd, 0xae, 0xd2, 0x40,
	0x14, 0xee, 0x78, 0x0d, 0x91, 0x31, 0xb9, 0x89, 0xf5, 0x2e, 0x0a, 0x21, 0x85, 0xb0, 0x40, 0x36,
	0x4c, 0x05, 0x4d, 0x5c, 0xe8, 0x86, 0xca, 0x96, 0x4d, 0xdd, 0x10, 0x37, 0xa6, 0x3f, 0x43, 0x99,
	0x08, 0x33, 0x38, 0x33, 0xad, 0xf0, 0x16, 0x3e, 0x80, 0x8f, 0xc1, 0x43, 0xb0, 0x24, 0xac, 0xd4,
	0x05, 0x31, 0xf0, 0x22, 0xa6, 0xed, 0x50, 0xe4, 0x2f, 0x61, 0x79, 0x57, 0xed, 0x9c, 0xef, 0x7c,
	0x67, 0xce, 0xf7, 0x9d, 0x39, 0xb0, 0x32, 0x8c, 0xf8, 0xdc, 0x1a, 0x93, 0x6f, 0x11, 0x09, 0xac,
	0xb8, 0xed, 0x61, 0xe9, 0xb6, 0x2d, 0x39, 0x43, 0x53, 0xce, 0x24, 0xd3, 0x5f, 0x26, 0x28, 0xca,
	0x50, 0xa4, 0xd0, 0xb2, 0xe9, 0x33, 0x31, 0x61, 0xc2, 0xf2, 0x5c, 0x81, 0x73, 0x8a, 0xcf, 0x08,
	0xcd, 0x48, 0xe5, 0x52, 0x86, 0x7f, 0x49, 0x4f, 0x56, 0x76, 0x50, 0xd0, 0x43, 0xc8, 0x42, 0x96,
	0xc5, 0x93, 0xbf, 0x2c, 0x5a, 0xff, 0x09, 0xe0, 0x8b, 0xbe, 0x08, 0xfb, 0x84, 0xca, 0x1e, 0xe6,
	0x24, 0x76, 0x25, 0x89, 0xb1, 0xfe, 0x1a, 0x16, 0x04, 0xa6, 0x01, 0xe6, 0x06, 0xa8, 0x81, 0x66,
	0xd1, 0x36, 0xd6, 0x8b, 0xd6, 0x83, 0xaa, 0xd6, 0x0d, 0x02, 0x8e, 0x85, 0xf8, 0x24, 0x39, 0xa1,
	0xa1, 0xa3, 0xf2, 0xf4, 0x0a, 0x2c, 0xc6, 0xee, 0x98, 0x04, 0xae, 0x64, 0xdc, 0x78, 0x92, 0x90,
	0x9c, 0x43, 0x40, 0x7f, 0x07, 0x0b, 0xee, 0x84, 0x45, 0x54, 0x1a, 0x77, 0x35, 0xd0, 0x7c, 0xde,
	0x29, 0x21, 0x55, 0x2c, 0xd1, 0xb1, 0x17, 0x87, 0x3e, 0x32, 0x42, 0xed, 0xa7, 0xcb, 0x4d, 0x55,
	0x73, 0x54, 0x7a, 0x7d, 0x00, 0x4b, 0x67, 0xdd, 0x39, 0x58, 0x4c, 0x19, 0x15, 0x58, 0x7f, 0x0f,
	0x9f, 0x71, 0xec, 0x63, 0x12, 0xe3, 0x20, 0xed, 0xf3, 0x86, 0xba, 0x39, 0x61, 0x2f, 0xdc, 0x8e,
	0x38, 0x7d, 0x8c, 0xc2, 0xa3, 0x54, 0xf8, 0x71, 0x77, 0xb9, 0xf0, 0xc1, 0x89, 0xf0, 0xa2, 0xfd,
	0x21, 0x21, 0xff, 0xd9, 0x54, 0x1b, 0x21, 0x91, 0xa3, 0xc8, 0x43, 0x3e, 0x9b, 0xa8, 0xe9, 0xab,
	0x4f, 0x4b, 0x04, 0x5f, 0x2d, 0x39, 0x9f, 0x62, 0x81, 0x7a, 0xd8, 0x5f, 0x2f, 0x5a, 0x50, 0x35,
	0xd2, 0xc3, 0xfe, 0xc1, 0x95, 0xce, 0x6f, 0x00, 0xef, 0xfa, 0x22, 0xd4, 0x47, 0xf0, 0xfe, 0xe4,
	0x49, 0x34, 0xd0, 0x85, 0xf7, 0x88, 0xce, 0x86, 0x53, 0x46, 0xb7, 0xe5, 0xe5, 0x5a, 0x46, 0xf0,
	0xfe, 0x64, 0x06, 0x57, 0x6f, 0x3a, 0xce, 0xbb, 0x7e, 0xd3, 0x65, 0xd7, 0xec, 0xee, 0x72, 0x6b,
	0x82, 0xd5, 0xd6, 0x04, 0x7f, 0xb7, 0x26, 0xf8, 0xb1, 0x33, 0xb5, 0xd5, 0xce, 0xd4, 0x7e, 0xed,
	0x4c, 0xed, 0xf3, 0xab, 0xff, 0x5c, 0x1b, 0xb2, 0x88, 0xbf, 0x95, 0xdf, 0x99, 0x95, 0x2e, 0xe7,
	0x6c, 0xbf, 0x9e, 0xa9, 0x75, 0x5e, 0x21, 0x5d, 0x9a, 0x37, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xc2, 0xd5, 0x43, 0x28, 0xba, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// MintDerivative defines a method for converting a delegation into staking deriviatives.
	MintDerivative(ctx context.Context, in *MsgMintDerivative, opts ...grpc.CallOption) (*MsgMintDerivativeResponse, error)
	// BurnDerivative defines a method for converting staking deriviatives into a delegation.
	BurnDerivative(ctx context.Context, in *MsgBurnDerivative, opts ...grpc.CallOption) (*MsgBurnDerivativeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) MintDerivative(ctx context.Context, in *MsgMintDerivative, opts ...grpc.CallOption) (*MsgMintDerivativeResponse, error) {
	out := new(MsgMintDerivativeResponse)
	err := c.cc.Invoke(ctx, "/fury.liquid.v1beta1.Msg/MintDerivative", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BurnDerivative(ctx context.Context, in *MsgBurnDerivative, opts ...grpc.CallOption) (*MsgBurnDerivativeResponse, error) {
	out := new(MsgBurnDerivativeResponse)
	err := c.cc.Invoke(ctx, "/fury.liquid.v1beta1.Msg/BurnDerivative", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// MintDerivative defines a method for converting a delegation into staking deriviatives.
	MintDerivative(context.Context, *MsgMintDerivative) (*MsgMintDerivativeResponse, error)
	// BurnDerivative defines a method for converting staking deriviatives into a delegation.
	BurnDerivative(context.Context, *MsgBurnDerivative) (*MsgBurnDerivativeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) MintDerivative(ctx context.Context, req *MsgMintDerivative) (*MsgMintDerivativeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintDerivative not implemented")
}
func (*UnimplementedMsgServer) BurnDerivative(ctx context.Context, req *MsgBurnDerivative) (*MsgBurnDerivativeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BurnDerivative not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_MintDerivative_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintDerivative)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintDerivative(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fury.liquid.v1beta1.Msg/MintDerivative",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintDerivative(ctx, req.(*MsgMintDerivative))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BurnDerivative_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurnDerivative)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BurnDerivative(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fury.liquid.v1beta1.Msg/BurnDerivative",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BurnDerivative(ctx, req.(*MsgBurnDerivative))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fury.liquid.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MintDerivative",
			Handler:    _Msg_MintDerivative_Handler,
		},
		{
			MethodName: "BurnDerivative",
			Handler:    _Msg_BurnDerivative_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fury/liquid/v1beta1/tx.proto",
}

func (m *MsgMintDerivative) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintDerivative) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintDerivative) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMintDerivativeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintDerivativeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintDerivativeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Received.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgBurnDerivative) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurnDerivative) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurnDerivative) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBurnDerivativeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurnDerivativeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurnDerivativeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Received.Size()
		i -= size
		if _, err := m.Received.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgMintDerivative) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgMintDerivativeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Received.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgBurnDerivative) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgBurnDerivativeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Received.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgMintDerivative) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMintDerivative: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintDerivative: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgMintDerivativeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMintDerivativeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintDerivativeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Received", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Received.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgBurnDerivative) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBurnDerivative: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurnDerivative: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgBurnDerivativeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBurnDerivativeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurnDerivativeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Received", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Received.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
