// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/committee/v1beta1/committee.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// TallyOption enumerates the valid types of a tally.
type TallyOption int32

const (
	// TALLY_OPTION_UNSPECIFIED defines a null tally option.
	TALLY_OPTION_UNSPECIFIED TallyOption = 0
	// Votes are tallied each block and the proposal passes as soon as the vote threshold is reached
	TALLY_OPTION_FIRST_PAST_THE_POST TallyOption = 1
	// Votes are tallied exactly once, when the deadline time is reached
	TALLY_OPTION_DEADLINE TallyOption = 2
)

var TallyOption_name = map[int32]string{
	0: "TALLY_OPTION_UNSPECIFIED",
	1: "TALLY_OPTION_FIRST_PAST_THE_POST",
	2: "TALLY_OPTION_DEADLINE",
}

var TallyOption_value = map[string]int32{
	"TALLY_OPTION_UNSPECIFIED":         0,
	"TALLY_OPTION_FIRST_PAST_THE_POST": 1,
	"TALLY_OPTION_DEADLINE":            2,
}

func (x TallyOption) String() string {
	return proto.EnumName(TallyOption_name, int32(x))
}

func (TallyOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c873432765d1f05e, []int{0}
}

// BaseCommittee is a common type shared by all Committees
type BaseCommittee struct {
	ID          uint64                                          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string                                          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Members     []github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,rep,name=members,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"members,omitempty"`
	Permissions []*types.Any                                    `protobuf:"bytes,4,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// Smallest percentage that must vote for a proposal to pass
	VoteThreshold github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=vote_threshold,json=voteThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"vote_threshold"`
	// The length of time a proposal remains active for. Proposals will close earlier if they get enough votes.
	ProposalDuration time.Duration `protobuf:"bytes,6,opt,name=proposal_duration,json=proposalDuration,proto3,stdduration" json:"proposal_duration"`
	TallyOption      TallyOption   `protobuf:"varint,7,opt,name=tally_option,json=tallyOption,proto3,enum=fury.committee.v1beta1.TallyOption" json:"tally_option,omitempty"`
}

func (m *BaseCommittee) Reset()      { *m = BaseCommittee{} }
func (*BaseCommittee) ProtoMessage() {}
func (*BaseCommittee) Descriptor() ([]byte, []int) {
	return fileDescriptor_c873432765d1f05e, []int{0}
}
func (m *BaseCommittee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseCommittee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseCommittee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseCommittee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseCommittee.Merge(m, src)
}
func (m *BaseCommittee) XXX_Size() int {
	return m.Size()
}
func (m *BaseCommittee) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseCommittee.DiscardUnknown(m)
}

var xxx_messageInfo_BaseCommittee proto.InternalMessageInfo

// MemberCommittee is an alias of BaseCommittee
type MemberCommittee struct {
	*BaseCommittee `protobuf:"bytes,1,opt,name=base_committee,json=baseCommittee,proto3,embedded=base_committee" json:"base_committee,omitempty"`
}

func (m *MemberCommittee) Reset()      { *m = MemberCommittee{} }
func (*MemberCommittee) ProtoMessage() {}
func (*MemberCommittee) Descriptor() ([]byte, []int) {
	return fileDescriptor_c873432765d1f05e, []int{1}
}
func (m *MemberCommittee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MemberCommittee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberCommittee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MemberCommittee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberCommittee.Merge(m, src)
}
func (m *MemberCommittee) XXX_Size() int {
	return m.Size()
}
func (m *MemberCommittee) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberCommittee.DiscardUnknown(m)
}

var xxx_messageInfo_MemberCommittee proto.InternalMessageInfo

// TokenCommittee supports voting on proposals by token holders
type TokenCommittee struct {
	*BaseCommittee `protobuf:"bytes,1,opt,name=base_committee,json=baseCommittee,proto3,embedded=base_committee" json:"base_committee,omitempty"`
	Quorum         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=quorum,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"quorum"`
	TallyDenom     string                                 `protobuf:"bytes,3,opt,name=tally_denom,json=tallyDenom,proto3" json:"tally_denom,omitempty"`
}

func (m *TokenCommittee) Reset()      { *m = TokenCommittee{} }
func (*TokenCommittee) ProtoMessage() {}
func (*TokenCommittee) Descriptor() ([]byte, []int) {
	return fileDescriptor_c873432765d1f05e, []int{2}
}
func (m *TokenCommittee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenCommittee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenCommittee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenCommittee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenCommittee.Merge(m, src)
}
func (m *TokenCommittee) XXX_Size() int {
	return m.Size()
}
func (m *TokenCommittee) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenCommittee.DiscardUnknown(m)
}

var xxx_messageInfo_TokenCommittee proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("fury.committee.v1beta1.TallyOption", TallyOption_name, TallyOption_value)
	proto.RegisterType((*BaseCommittee)(nil), "fury.committee.v1beta1.BaseCommittee")
	proto.RegisterType((*MemberCommittee)(nil), "fury.committee.v1beta1.MemberCommittee")
	proto.RegisterType((*TokenCommittee)(nil), "fury.committee.v1beta1.TokenCommittee")
}

func init() {
	proto.RegisterFile("fury/committee/v1beta1/committee.proto", fileDescriptor_c873432765d1f05e)
}

var fileDescriptor_c873432765d1f05e = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xb5, 0x93, 0x7c, 0xe9, 0xd7, 0x71, 0x1b, 0xd2, 0xa1, 0x54, 0x4e, 0x85, 0x6c, 0xab, 0x40,
	0x15, 0x55, 0x8a, 0xad, 0x06, 0x56, 0xec, 0xe2, 0x3a, 0x51, 0x23, 0x95, 0x26, 0x72, 0xdc, 0x05,
	0x6c, 0x2c, 0x3b, 0x9e, 0xa6, 0x56, 0xe3, 0x4c, 0xf0, 0x8c, 0x4b, 0xf3, 0x06, 0x2c, 0x59, 0x76,
	0x89, 0xc4, 0x2b, 0xf4, 0x21, 0xaa, 0xae, 0x2a, 0x56, 0x88, 0x45, 0x28, 0xe9, 0x53, 0xc0, 0x0a,
	0xf9, 0xaf, 0x49, 0x21, 0x48, 0xb0, 0x60, 0x65, 0xcf, 0xb9, 0xe7, 0xde, 0x3b, 0xe7, 0xde, 0xa3,
	0x01, 0x9b, 0x87, 0x81, 0x3f, 0x52, 0xba, 0xd8, 0xf3, 0x5c, 0x4a, 0x11, 0x52, 0x4e, 0xb6, 0x6d,
	0x44, 0xad, 0xed, 0x29, 0x22, 0x0f, 0x7d, 0x4c, 0x31, 0x5c, 0x0b, 0x79, 0xf2, 0x14, 0x4d, 0x78,
	0xeb, 0xa5, 0x2e, 0x26, 0x1e, 0x26, 0x66, 0xc4, 0x52, 0xe2, 0x43, 0x9c, 0xb2, 0xbe, 0xda, 0xc3,
	0x3d, 0x1c, 0xe3, 0xe1, 0x5f, 0x82, 0x96, 0x7a, 0x18, 0xf7, 0xfa, 0x48, 0x89, 0x4e, 0x76, 0x70,
	0xa8, 0x58, 0x83, 0x51, 0x12, 0x12, 0x7e, 0x0e, 0x39, 0x81, 0x6f, 0x51, 0x17, 0x0f, 0xe2, 0xf8,
	0xc6, 0xb7, 0x2c, 0x58, 0x56, 0x2d, 0x82, 0x76, 0xd2, 0x5b, 0xc0, 0x35, 0x90, 0x71, 0x1d, 0x9e,
	0x95, 0xd8, 0x72, 0x4e, 0xcd, 0x4f, 0xc6, 0x62, 0xa6, 0xa9, 0xe9, 0x19, 0xd7, 0x81, 0x12, 0xe0,
	0x1c, 0x44, 0xba, 0xbe, 0x3b, 0x0c, 0xd3, 0xf9, 0x8c, 0xc4, 0x96, 0x17, 0xf5, 0x59, 0x08, 0xda,
	0x60, 0xc1, 0x43, 0x9e, 0x8d, 0x7c, 0xc2, 0x67, 0xa5, 0x6c, 0x79, 0x49, 0xdd, 0xfd, 0x3e, 0x16,
	0x2b, 0x3d, 0x97, 0x1e, 0x05, 0x76, 0x28, 0x33, 0x91, 0x92, 0x7c, 0x2a, 0xc4, 0x39, 0x56, 0xe8,
	0x68, 0x88, 0x88, 0x5c, 0xeb, 0x76, 0x6b, 0x8e, 0xe3, 0x23, 0x42, 0x3e, 0x9e, 0x57, 0xee, 0x27,
	0x82, 0x13, 0x44, 0x1d, 0x51, 0x44, 0xf4, 0xb4, 0x30, 0x6c, 0x00, 0x6e, 0x88, 0x7c, 0xcf, 0x25,
	0xc4, 0xc5, 0x03, 0xc2, 0xe7, 0xa4, 0x6c, 0x99, 0xab, 0xae, 0xca, 0xb1, 0x4a, 0x39, 0x55, 0x29,
	0xd7, 0x06, 0x23, 0xb5, 0x70, 0x79, 0x5e, 0x01, 0xed, 0x5b, 0xb2, 0x3e, 0x9b, 0x08, 0x0f, 0x40,
	0xe1, 0x04, 0x53, 0x64, 0xd2, 0x23, 0x1f, 0x91, 0x23, 0xdc, 0x77, 0xf8, 0xff, 0x42, 0x41, 0xaa,
	0x7c, 0x31, 0x16, 0x99, 0xcf, 0x63, 0x71, 0xf3, 0x0f, 0xae, 0xad, 0xa1, 0xae, 0xbe, 0x1c, 0x56,
	0x31, 0xd2, 0x22, 0xb0, 0x0d, 0x56, 0x86, 0x3e, 0x1e, 0x62, 0x62, 0xf5, 0xcd, 0x74, 0xd2, 0x7c,
	0x5e, 0x62, 0xcb, 0x5c, 0xb5, 0xf4, 0xcb, 0x25, 0xb5, 0x84, 0xa0, 0xfe, 0x1f, 0x36, 0x3d, 0xfb,
	0x22, 0xb2, 0x7a, 0x31, 0xcd, 0x4e, 0x63, 0xb0, 0x01, 0x96, 0xa8, 0xd5, 0xef, 0x8f, 0x4c, 0x1c,
	0xcf, 0x7d, 0x41, 0x62, 0xcb, 0x85, 0xea, 0x23, 0x79, 0xbe, 0x77, 0x64, 0x23, 0xe4, 0xb6, 0x22,
	0xaa, 0xce, 0xd1, 0xe9, 0xe1, 0xf9, 0xca, 0xd9, 0x7b, 0x91, 0xb9, 0x3c, 0xaf, 0x2c, 0xde, 0x6e,
	0x7a, 0xe3, 0x14, 0xdc, 0x7b, 0x11, 0x8d, 0x75, 0xba, 0x7c, 0x1d, 0x14, 0x6c, 0x8b, 0x20, 0xf3,
	0xb6, 0x70, 0x64, 0x04, 0xae, 0xfa, 0xe4, 0x77, 0xfd, 0xee, 0x78, 0x47, 0xcd, 0x5d, 0x8d, 0x45,
	0x56, 0x5f, 0xb6, 0x67, 0xc1, 0x79, 0x9d, 0xaf, 0x59, 0x50, 0x30, 0xf0, 0x31, 0x1a, 0xfc, 0xd3,
	0xce, 0xb0, 0x01, 0xf2, 0xaf, 0x03, 0xec, 0x07, 0x5e, 0xec, 0xd6, 0xbf, 0x5e, 0x6e, 0x92, 0x0d,
	0x45, 0x10, 0x8f, 0xd2, 0x74, 0xd0, 0x00, 0x7b, 0x7c, 0x36, 0xb2, 0x3e, 0x88, 0x20, 0x2d, 0x44,
	0xe6, 0x48, 0xdc, 0xf2, 0x01, 0x37, 0xb3, 0x0b, 0xf8, 0x10, 0xf0, 0x46, 0x6d, 0x6f, 0xef, 0xa5,
	0xd9, 0x6a, 0x1b, 0xcd, 0xd6, 0xbe, 0x79, 0xb0, 0xdf, 0x69, 0xd7, 0x77, 0x9a, 0x8d, 0x66, 0x5d,
	0x2b, 0x32, 0xf0, 0x31, 0x90, 0xee, 0x44, 0x1b, 0x4d, 0xbd, 0x63, 0x98, 0xed, 0x5a, 0xc7, 0x30,
	0x8d, 0xdd, 0xba, 0xd9, 0x6e, 0x75, 0x8c, 0x22, 0x0b, 0x4b, 0xe0, 0xc1, 0x1d, 0x96, 0x56, 0xaf,
	0x69, 0x7b, 0xcd, 0xfd, 0x7a, 0x31, 0xb3, 0x9e, 0x7b, 0xfb, 0x41, 0x60, 0xd4, 0xdd, 0x8b, 0xaf,
	0x02, 0x73, 0x31, 0x11, 0xd8, 0xab, 0x89, 0xc0, 0x5e, 0x4f, 0x04, 0xf6, 0xdd, 0x8d, 0xc0, 0x5c,
	0xdd, 0x08, 0xcc, 0xa7, 0x1b, 0x81, 0x79, 0xb5, 0x35, 0xa3, 0xfa, 0x10, 0x07, 0xfe, 0x33, 0xfa,
	0x06, 0x2b, 0xd1, 0x53, 0x75, 0x3a, 0xf3, 0x58, 0x45, 0xea, 0xed, 0x7c, 0x64, 0xd2, 0xa7, 0x3f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xef, 0x97, 0xc8, 0xf1, 0xcb, 0x04, 0x00, 0x00,
}

func (m *BaseCommittee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseCommittee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseCommittee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TallyOption != 0 {
		i = encodeVarintCommittee(dAtA, i, uint64(m.TallyOption))
		i--
		dAtA[i] = 0x38
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.ProposalDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.ProposalDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintCommittee(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	{
		size := m.VoteThreshold.Size()
		i -= size
		if _, err := m.VoteThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCommittee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Permissions) > 0 {
		for iNdEx := len(m.Permissions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Permissions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommittee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Members[iNdEx])
			copy(dAtA[i:], m.Members[iNdEx])
			i = encodeVarintCommittee(dAtA, i, uint64(len(m.Members[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCommittee(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintCommittee(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MemberCommittee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberCommittee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberCommittee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BaseCommittee != nil {
		{
			size, err := m.BaseCommittee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommittee(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenCommittee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenCommittee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenCommittee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TallyDenom) > 0 {
		i -= len(m.TallyDenom)
		copy(dAtA[i:], m.TallyDenom)
		i = encodeVarintCommittee(dAtA, i, uint64(len(m.TallyDenom)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Quorum.Size()
		i -= size
		if _, err := m.Quorum.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCommittee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.BaseCommittee != nil {
		{
			size, err := m.BaseCommittee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommittee(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommittee(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommittee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseCommittee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovCommittee(uint64(m.ID))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCommittee(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, b := range m.Members {
			l = len(b)
			n += 1 + l + sovCommittee(uint64(l))
		}
	}
	if len(m.Permissions) > 0 {
		for _, e := range m.Permissions {
			l = e.Size()
			n += 1 + l + sovCommittee(uint64(l))
		}
	}
	l = m.VoteThreshold.Size()
	n += 1 + l + sovCommittee(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.ProposalDuration)
	n += 1 + l + sovCommittee(uint64(l))
	if m.TallyOption != 0 {
		n += 1 + sovCommittee(uint64(m.TallyOption))
	}
	return n
}

func (m *MemberCommittee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseCommittee != nil {
		l = m.BaseCommittee.Size()
		n += 1 + l + sovCommittee(uint64(l))
	}
	return n
}

func (m *TokenCommittee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseCommittee != nil {
		l = m.BaseCommittee.Size()
		n += 1 + l + sovCommittee(uint64(l))
	}
	l = m.Quorum.Size()
	n += 1 + l + sovCommittee(uint64(l))
	l = len(m.TallyDenom)
	if l > 0 {
		n += 1 + l + sovCommittee(uint64(l))
	}
	return n
}

func sovCommittee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommittee(x uint64) (n int) {
	return sovCommittee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseCommittee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommittee
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
			return fmt.Errorf("proto: BaseCommittee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseCommittee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, make([]byte, postIndex-iNdEx))
			copy(m.Members[len(m.Members)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permissions = append(m.Permissions, &types.Any{})
			if err := m.Permissions[len(m.Permissions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VoteThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.ProposalDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TallyOption", wireType)
			}
			m.TallyOption = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TallyOption |= TallyOption(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommittee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommittee
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
func (m *MemberCommittee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommittee
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
			return fmt.Errorf("proto: MemberCommittee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberCommittee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseCommittee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseCommittee == nil {
				m.BaseCommittee = &BaseCommittee{}
			}
			if err := m.BaseCommittee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommittee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommittee
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
func (m *TokenCommittee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommittee
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
			return fmt.Errorf("proto: TokenCommittee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenCommittee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseCommittee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseCommittee == nil {
				m.BaseCommittee = &BaseCommittee{}
			}
			if err := m.BaseCommittee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quorum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quorum.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TallyDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TallyDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommittee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommittee
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
func skipCommittee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommittee
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
					return 0, ErrIntOverflowCommittee
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
					return 0, ErrIntOverflowCommittee
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
				return 0, ErrInvalidLengthCommittee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommittee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommittee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommittee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommittee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommittee = fmt.Errorf("proto: unexpected end of group")
)
