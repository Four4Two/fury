// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/evmutil/v1beta1/genesis.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the evmutil module's genesis state.
type GenesisState struct {
	Accounts []Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts"`
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea75ced7569751e4, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

// BalanceAccount defines an account in the evmutil module.
type Account struct {
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
	// balance indicates the amount of afury owned by the address.
	Balance github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=balance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"balance"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea75ced7569751e4, []int{1}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

// Params defines the evmutil module params
type Params struct {
	// enabled_conversion_pairs defines the list of conversion pairs allowed to be
	// converted between Fury ERC20 and sdk.Coin
	EnabledConversionPairs ConversionPairs `protobuf:"bytes,4,rep,name=enabled_conversion_pairs,json=enabledConversionPairs,proto3,castrepeated=ConversionPairs" json:"enabled_conversion_pairs"`
	// allowed_cosmos_denoms is a list of denom & erc20 token metadata pairs.
	// if a denom is in the list, it is allowed to be converted to an erc20 in the evm.
	AllowedCosmosDenoms AllowedCosmosCoinERC20Tokens `protobuf:"bytes,1,rep,name=allowed_cosmos_denoms,json=allowedCosmosDenoms,proto3,castrepeated=AllowedCosmosCoinERC20Tokens" json:"allowed_cosmos_denoms"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea75ced7569751e4, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEnabledConversionPairs() ConversionPairs {
	if m != nil {
		return m.EnabledConversionPairs
	}
	return nil
}

func (m *Params) GetAllowedCosmosDenoms() AllowedCosmosCoinERC20Tokens {
	if m != nil {
		return m.AllowedCosmosDenoms
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "fury.evmutil.v1beta1.GenesisState")
	proto.RegisterType((*Account)(nil), "fury.evmutil.v1beta1.Account")
	proto.RegisterType((*Params)(nil), "fury.evmutil.v1beta1.Params")
}

func init() {
	proto.RegisterFile("fury/evmutil/v1beta1/genesis.proto", fileDescriptor_ea75ced7569751e4)
}

var fileDescriptor_ea75ced7569751e4 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xf5, 0x96, 0x28, 0x81, 0x6d, 0x25, 0x24, 0xb7, 0x40, 0xa8, 0x8a, 0x5d, 0x45, 0x15, 0x8a,
	0x90, 0x62, 0x93, 0xc0, 0xa9, 0x42, 0x42, 0x75, 0x40, 0x50, 0x4e, 0x95, 0x41, 0x1c, 0xb8, 0x44,
	0x6b, 0x7b, 0x1b, 0xac, 0xda, 0x3b, 0x91, 0x77, 0x9d, 0x90, 0x3f, 0x40, 0xe2, 0x00, 0xfc, 0x01,
	0x47, 0xc4, 0xb9, 0x1f, 0x51, 0x89, 0x4b, 0xd5, 0x13, 0xe2, 0x10, 0x4a, 0xf2, 0x17, 0x9c, 0x90,
	0x77, 0x37, 0x51, 0x89, 0x0c, 0xe2, 0x64, 0x7b, 0xfc, 0xde, 0x9b, 0x37, 0x6f, 0x06, 0x37, 0x0e,
	0xf3, 0x6c, 0xec, 0xd2, 0x61, 0x9a, 0x8b, 0x38, 0x71, 0x87, 0xed, 0x80, 0x0a, 0xd2, 0x76, 0xfb,
	0x94, 0x51, 0x1e, 0x73, 0x67, 0x90, 0x81, 0x00, 0x73, 0xa3, 0xc0, 0x38, 0x1a, 0xe3, 0x68, 0xcc,
	0xe6, 0xcd, 0x10, 0x78, 0x0a, 0xbc, 0x27, 0x31, 0xae, 0xfa, 0x50, 0x84, 0xcd, 0x3b, 0xa5, 0xa2,
	0x21, 0xb0, 0x21, 0xcd, 0x78, 0x0c, 0xac, 0x37, 0x20, 0x71, 0xa6, 0xb1, 0x1b, 0x7d, 0xe8, 0x83,
	0xd2, 0x28, 0xde, 0x54, 0xb5, 0xf1, 0x11, 0xe1, 0xb5, 0x27, 0xca, 0xc4, 0x73, 0x41, 0x04, 0x35,
	0x1f, 0xe2, 0xcb, 0x24, 0x0c, 0x21, 0x67, 0x82, 0xd7, 0xd1, 0xf6, 0xa5, 0xe6, 0x6a, 0xe7, 0x96,
	0x53, 0x66, 0xcb, 0xd9, 0x53, 0x28, 0xaf, 0x72, 0x32, 0xb1, 0x0d, 0x7f, 0x41, 0x32, 0x77, 0x71,
	0x75, 0x40, 0x32, 0x92, 0xf2, 0xfa, 0xca, 0x36, 0x6a, 0xae, 0x76, 0xb6, 0xca, 0xe9, 0x07, 0x12,
	0xa3, 0xd9, 0x9a, 0xb1, 0x5b, 0x79, 0xfb, 0xc9, 0x36, 0x1a, 0x5f, 0x11, 0xae, 0x69, 0x75, 0x33,
	0xc0, 0x35, 0x12, 0x45, 0x19, 0xe5, 0x85, 0x1b, 0xd4, 0x5c, 0xf3, 0x9e, 0xfe, 0x9a, 0xd8, 0xad,
	0x7e, 0x2c, 0x5e, 0xe7, 0x81, 0x13, 0x42, 0xaa, 0xf3, 0xd0, 0x8f, 0x16, 0x8f, 0x8e, 0x5c, 0x31,
	0x1e, 0x50, 0x5e, 0xd8, 0xdb, 0x53, 0xc4, 0xb3, 0xe3, 0xd6, 0xba, 0x4e, 0x4d, 0x57, 0xbc, 0xb1,
	0xa0, 0xdc, 0x9f, 0x0b, 0x9b, 0x2f, 0x71, 0x2d, 0x20, 0x09, 0x61, 0x21, 0x95, 0x96, 0xaf, 0x78,
	0x0f, 0x0a, 0x53, 0xdf, 0x27, 0xf6, 0xed, 0xff, 0xe8, 0xb3, 0xcf, 0xc4, 0xd9, 0x71, 0x0b, 0xeb,
	0x06, 0xfb, 0x4c, 0xf8, 0x73, 0x31, 0x3d, 0xcd, 0xfb, 0x15, 0x5c, 0x55, 0xc3, 0x9a, 0x23, 0x5c,
	0xa7, 0x8c, 0x04, 0x09, 0x8d, 0x7a, 0x4b, 0x3b, 0xe2, 0xf5, 0x8a, 0xcc, 0x7a, 0xa7, 0x3c, 0xac,
	0xee, 0x02, 0x7d, 0x40, 0xe2, 0xcc, 0xbb, 0x51, 0xf8, 0xfb, 0xf2, 0xc3, 0xbe, 0xfa, 0x67, 0x9d,
	0xfb, 0xd7, 0xb5, 0xfc, 0x52, 0xdd, 0x7c, 0x87, 0xf0, 0x35, 0x92, 0x24, 0x30, 0x92, 0x9d, 0xe5,
	0x35, 0x45, 0x94, 0x41, 0x3a, 0x5f, 0x71, 0xfb, 0x2f, 0x2b, 0x56, 0x94, 0xae, 0x64, 0x74, 0x21,
	0x66, 0x8f, 0xfd, 0x6e, 0xe7, 0xee, 0x0b, 0x38, 0xa2, 0xcc, 0xdb, 0xd1, 0x1e, 0xb6, 0xfe, 0x01,
	0xe2, 0xfe, 0x3a, 0xb9, 0xf8, 0xf7, 0x91, 0xec, 0xe9, 0x3d, 0x3b, 0xff, 0x69, 0xa1, 0xcf, 0x53,
	0x0b, 0x9d, 0x4c, 0x2d, 0x74, 0x3a, 0xb5, 0xd0, 0xf9, 0xd4, 0x42, 0x1f, 0x66, 0x96, 0x71, 0x3a,
	0xb3, 0x8c, 0x6f, 0x33, 0xcb, 0x78, 0xd5, 0xbc, 0x10, 0xfc, 0x21, 0xe4, 0xd9, 0x7d, 0x31, 0x02,
	0x57, 0xde, 0xfa, 0x9b, 0xc5, 0xb5, 0xcb, 0xf8, 0x83, 0xaa, 0x3c, 0xe3, 0x7b, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x91, 0x8b, 0xbf, 0x5a, 0x5f, 0x03, 0x00, 0x00,
}

func (this *GenesisState) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*GenesisState)
	if !ok {
		that2, ok := that.(GenesisState)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *GenesisState")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *GenesisState but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *GenesisState but is not nil && this == nil")
	}
	if len(this.Accounts) != len(that1.Accounts) {
		return fmt.Errorf("Accounts this(%v) Not Equal that(%v)", len(this.Accounts), len(that1.Accounts))
	}
	for i := range this.Accounts {
		if !this.Accounts[i].Equal(&that1.Accounts[i]) {
			return fmt.Errorf("Accounts this[%v](%v) Not Equal that[%v](%v)", i, this.Accounts[i], i, that1.Accounts[i])
		}
	}
	if !this.Params.Equal(&that1.Params) {
		return fmt.Errorf("Params this(%v) Not Equal that(%v)", this.Params, that1.Params)
	}
	return nil
}
func (this *GenesisState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisState)
	if !ok {
		that2, ok := that.(GenesisState)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Accounts) != len(that1.Accounts) {
		return false
	}
	for i := range this.Accounts {
		if !this.Accounts[i].Equal(&that1.Accounts[i]) {
			return false
		}
	}
	if !this.Params.Equal(&that1.Params) {
		return false
	}
	return true
}
func (this *Account) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Account)
	if !ok {
		that2, ok := that.(Account)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Account")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Account but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Account but is not nil && this == nil")
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return fmt.Errorf("Address this(%v) Not Equal that(%v)", this.Address, that1.Address)
	}
	if !this.Balance.Equal(that1.Balance) {
		return fmt.Errorf("Balance this(%v) Not Equal that(%v)", this.Balance, that1.Balance)
	}
	return nil
}
func (this *Account) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Account)
	if !ok {
		that2, ok := that.(Account)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if !this.Balance.Equal(that1.Balance) {
		return false
	}
	return true
}
func (this *Params) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Params")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Params but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Params but is not nil && this == nil")
	}
	if len(this.EnabledConversionPairs) != len(that1.EnabledConversionPairs) {
		return fmt.Errorf("EnabledConversionPairs this(%v) Not Equal that(%v)", len(this.EnabledConversionPairs), len(that1.EnabledConversionPairs))
	}
	for i := range this.EnabledConversionPairs {
		if !this.EnabledConversionPairs[i].Equal(&that1.EnabledConversionPairs[i]) {
			return fmt.Errorf("EnabledConversionPairs this[%v](%v) Not Equal that[%v](%v)", i, this.EnabledConversionPairs[i], i, that1.EnabledConversionPairs[i])
		}
	}
	if len(this.AllowedCosmosDenoms) != len(that1.AllowedCosmosDenoms) {
		return fmt.Errorf("AllowedCosmosDenoms this(%v) Not Equal that(%v)", len(this.AllowedCosmosDenoms), len(that1.AllowedCosmosDenoms))
	}
	for i := range this.AllowedCosmosDenoms {
		if !this.AllowedCosmosDenoms[i].Equal(&that1.AllowedCosmosDenoms[i]) {
			return fmt.Errorf("AllowedCosmosDenoms this[%v](%v) Not Equal that[%v](%v)", i, this.AllowedCosmosDenoms[i], i, that1.AllowedCosmosDenoms[i])
		}
	}
	return nil
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.EnabledConversionPairs) != len(that1.EnabledConversionPairs) {
		return false
	}
	for i := range this.EnabledConversionPairs {
		if !this.EnabledConversionPairs[i].Equal(&that1.EnabledConversionPairs[i]) {
			return false
		}
	}
	if len(this.AllowedCosmosDenoms) != len(that1.AllowedCosmosDenoms) {
		return false
	}
	for i := range this.AllowedCosmosDenoms {
		if !this.AllowedCosmosDenoms[i].Equal(&that1.AllowedCosmosDenoms[i]) {
			return false
		}
	}
	return true
}
func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Accounts) > 0 {
		for iNdEx := len(m.Accounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Accounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Balance.Size()
		i -= size
		if _, err := m.Balance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EnabledConversionPairs) > 0 {
		for iNdEx := len(m.EnabledConversionPairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EnabledConversionPairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AllowedCosmosDenoms) > 0 {
		for iNdEx := len(m.AllowedCosmosDenoms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AllowedCosmosDenoms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for _, e := range m.Accounts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Balance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AllowedCosmosDenoms) > 0 {
		for _, e := range m.AllowedCosmosDenoms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.EnabledConversionPairs) > 0 {
		for _, e := range m.EnabledConversionPairs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accounts = append(m.Accounts, Account{})
			if err := m.Accounts[len(m.Accounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedCosmosDenoms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedCosmosDenoms = append(m.AllowedCosmosDenoms, AllowedCosmosCoinERC20Token{})
			if err := m.AllowedCosmosDenoms[len(m.AllowedCosmosDenoms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnabledConversionPairs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnabledConversionPairs = append(m.EnabledConversionPairs, ConversionPair{})
			if err := m.EnabledConversionPairs[len(m.EnabledConversionPairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
