// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/lightclients/harmony/v1/harmony.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/ibc-go/modules/core/02-client/types"
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

type ClientState struct {
	ShardId         uint32        `protobuf:"varint,1,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	ContractAddress []byte        `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	LatestEpoch     uint64        `protobuf:"varint,3,opt,name=latest_epoch,json=latestEpoch,proto3" json:"latest_epoch,omitempty"`
	LatestCommittee []byte        `protobuf:"bytes,4,opt,name=latest_committee,json=latestCommittee,proto3" json:"latest_committee,omitempty"`
	LatestHeight    *types.Height `protobuf:"bytes,5,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height,omitempty"`
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c699a4b541c7b8b1, []int{0}
}
func (m *ClientState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientState.Merge(m, src)
}
func (m *ClientState) XXX_Size() int {
	return m.Size()
}
func (m *ClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientState.DiscardUnknown(m)
}

var xxx_messageInfo_ClientState proto.InternalMessageInfo

type ConsensusState struct {
	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Root      []byte `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
}

func (m *ConsensusState) Reset()         { *m = ConsensusState{} }
func (m *ConsensusState) String() string { return proto.CompactTextString(m) }
func (*ConsensusState) ProtoMessage()    {}
func (*ConsensusState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c699a4b541c7b8b1, []int{1}
}
func (m *ConsensusState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsensusState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsensusState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsensusState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusState.Merge(m, src)
}
func (m *ConsensusState) XXX_Size() int {
	return m.Size()
}
func (m *ConsensusState) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusState.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusState proto.InternalMessageInfo

type EpochState struct {
	Committee []byte `protobuf:"bytes,1,opt,name=committee,proto3" json:"committee,omitempty"`
}

func (m *EpochState) Reset()         { *m = EpochState{} }
func (m *EpochState) String() string { return proto.CompactTextString(m) }
func (*EpochState) ProtoMessage()    {}
func (*EpochState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c699a4b541c7b8b1, []int{2}
}
func (m *EpochState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochState.Merge(m, src)
}
func (m *EpochState) XXX_Size() int {
	return m.Size()
}
func (m *EpochState) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochState.DiscardUnknown(m)
}

var xxx_messageInfo_EpochState proto.InternalMessageInfo

type Header struct {
	BeaconHeader   []byte `protobuf:"bytes,1,opt,name=beacon_header,json=beaconHeader,proto3" json:"beacon_header,omitempty"`
	ShardHeader    []byte `protobuf:"bytes,2,opt,name=shard_header,json=shardHeader,proto3" json:"shard_header,omitempty"`
	Signature      []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Bitmap         []byte `protobuf:"bytes,4,opt,name=bitmap,proto3" json:"bitmap,omitempty"`
	CrossLinkIndex uint32 `protobuf:"varint,5,opt,name=cross_link_index,json=crossLinkIndex,proto3" json:"cross_link_index,omitempty"`
	AccountProof   []byte `protobuf:"bytes,6,opt,name=account_proof,json=accountProof,proto3" json:"account_proof,omitempty"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_c699a4b541c7b8b1, []int{3}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClientState)(nil), "ibc.lightclients.harmony.v1.ClientState")
	proto.RegisterType((*ConsensusState)(nil), "ibc.lightclients.harmony.v1.ConsensusState")
	proto.RegisterType((*EpochState)(nil), "ibc.lightclients.harmony.v1.EpochState")
	proto.RegisterType((*Header)(nil), "ibc.lightclients.harmony.v1.Header")
}

func init() {
	proto.RegisterFile("ibc/lightclients/harmony/v1/harmony.proto", fileDescriptor_c699a4b541c7b8b1)
}

var fileDescriptor_c699a4b541c7b8b1 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0x63, 0x08, 0x81, 0x3a, 0x49, 0xa9, 0x2c, 0x84, 0x96, 0x50, 0x2d, 0x69, 0xb8, 0xa4,
	0x48, 0xd9, 0x55, 0xe0, 0x01, 0x10, 0x8d, 0x90, 0x5a, 0x09, 0x21, 0xb4, 0xdc, 0xb8, 0xac, 0xbc,
	0x5e, 0x93, 0xb5, 0xba, 0x6b, 0xaf, 0xec, 0xd9, 0x88, 0xbe, 0x05, 0x8f, 0xd5, 0x63, 0x2f, 0x48,
	0x1c, 0x21, 0x79, 0x11, 0xe4, 0x3f, 0x21, 0x88, 0xdb, 0xcc, 0xcf, 0xdf, 0x8c, 0xe7, 0x1b, 0x0d,
	0x3e, 0x17, 0x05, 0x4b, 0x6b, 0xb1, 0xae, 0x80, 0xd5, 0x82, 0x4b, 0x30, 0x69, 0x45, 0x75, 0xa3,
	0xe4, 0x4d, 0xba, 0x59, 0xee, 0xc3, 0xa4, 0xd5, 0x0a, 0x14, 0x79, 0x2e, 0x0a, 0x96, 0xfc, 0x2b,
	0x4d, 0xf6, 0xef, 0x9b, 0xe5, 0xe4, 0xc9, 0x5a, 0xad, 0x95, 0xd3, 0xa5, 0x36, 0xf2, 0x25, 0x93,
	0x17, 0xb6, 0x3b, 0x53, 0x9a, 0xa7, 0xbe, 0xc4, 0x36, 0xf5, 0x91, 0x17, 0xcc, 0xb6, 0x08, 0x0f,
	0x57, 0x0e, 0x7c, 0x06, 0x0a, 0x9c, 0x3c, 0xc3, 0x8f, 0x4c, 0x45, 0x75, 0x99, 0x8b, 0x32, 0x42,
	0x53, 0x34, 0x1f, 0x67, 0x0f, 0x5d, 0x7e, 0x55, 0x92, 0x73, 0x7c, 0xc2, 0x94, 0x04, 0x4d, 0x19,
	0xe4, 0xb4, 0x2c, 0x35, 0x37, 0x26, 0xba, 0x37, 0x45, 0xf3, 0x51, 0xf6, 0x78, 0xcf, 0xdf, 0x79,
	0x4c, 0xce, 0xf0, 0xa8, 0xa6, 0xc0, 0x0d, 0xe4, 0xbc, 0x55, 0xac, 0x8a, 0xee, 0x4f, 0xd1, 0xbc,
	0x9f, 0x0d, 0x3d, 0x7b, 0x6f, 0x91, 0xed, 0x16, 0x24, 0x4c, 0x35, 0x8d, 0x00, 0xe0, 0x3c, 0xea,
	0xfb, 0x6e, 0x9e, 0xaf, 0xf6, 0x98, 0xbc, 0xc5, 0xe3, 0x20, 0xad, 0xb8, 0x75, 0x1f, 0x3d, 0x98,
	0xa2, 0xf9, 0xf0, 0xf5, 0x24, 0xb1, 0xfb, 0xb0, 0xe6, 0x92, 0x60, 0x69, 0xb3, 0x4c, 0x2e, 0x9d,
	0x22, 0x0b, 0xdf, 0xfb, 0x6c, 0x76, 0x81, 0x8f, 0x57, 0x4a, 0x1a, 0x2e, 0x4d, 0x67, 0xbc, 0xcd,
	0x53, 0x7c, 0x04, 0xa2, 0xe1, 0x06, 0x68, 0xd3, 0x3a, 0x9f, 0xfd, 0xec, 0x00, 0x08, 0xc1, 0x7d,
	0xad, 0x14, 0x04, 0x77, 0x2e, 0x9e, 0xbd, 0xc2, 0xd8, 0x0d, 0xfe, 0xb7, 0xfe, 0x30, 0x36, 0x72,
	0xb2, 0x03, 0x98, 0xfd, 0x40, 0x78, 0x70, 0xc9, 0x69, 0xc9, 0x35, 0x79, 0x89, 0xc7, 0x05, 0xa7,
	0x4c, 0xc9, 0xbc, 0x72, 0x20, 0x88, 0x47, 0x1e, 0x06, 0xd1, 0x19, 0x1e, 0xf9, 0xa5, 0x07, 0x8d,
	0xff, 0x77, 0xe8, 0x58, 0x90, 0x9c, 0xe2, 0x23, 0x23, 0xd6, 0x92, 0x42, 0xa7, 0xb9, 0x5b, 0xe7,
	0x28, 0x3b, 0x00, 0xf2, 0x14, 0x0f, 0x0a, 0x01, 0x0d, 0x6d, 0xc3, 0x0a, 0x43, 0x46, 0xe6, 0xf8,
	0x84, 0x69, 0x65, 0x4c, 0x5e, 0x0b, 0x79, 0x9d, 0x0b, 0x59, 0xf2, 0x6f, 0x6e, 0x79, 0xe3, 0xec,
	0xd8, 0xf1, 0x0f, 0x42, 0x5e, 0x5f, 0x59, 0x6a, 0xe7, 0xa4, 0x8c, 0xa9, 0x4e, 0x42, 0xde, 0x6a,
	0xa5, 0xbe, 0x46, 0x03, 0x3f, 0x67, 0x80, 0x9f, 0x2c, 0xbb, 0xa8, 0x6f, 0x7f, 0xc7, 0xbd, 0xdb,
	0x6d, 0x8c, 0xee, 0xb6, 0x31, 0xfa, 0xb5, 0x8d, 0xd1, 0xf7, 0x5d, 0xdc, 0xbb, 0xdb, 0xc5, 0xbd,
	0x9f, 0xbb, 0xb8, 0xf7, 0xe5, 0xe3, 0x5a, 0x40, 0xd5, 0x15, 0x09, 0x53, 0x4d, 0x5a, 0x52, 0xa0,
	0xac, 0xa2, 0x42, 0xd6, 0xb4, 0x48, 0x45, 0xc1, 0x16, 0xe1, 0x52, 0x17, 0xe1, 0x0c, 0x1b, 0x55,
	0x76, 0x35, 0x37, 0xfe, 0xf0, 0x17, 0xff, 0x5f, 0x3e, 0xdc, 0xb4, 0xdc, 0x14, 0x03, 0x77, 0xa1,
	0x6f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x10, 0xc1, 0xa6, 0x22, 0x03, 0x00, 0x00,
}

func (m *ClientState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LatestHeight != nil {
		{
			size, err := m.LatestHeight.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHarmony(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.LatestCommittee) > 0 {
		i -= len(m.LatestCommittee)
		copy(dAtA[i:], m.LatestCommittee)
		i = encodeVarintHarmony(dAtA, i, uint64(len(m.LatestCommittee)))
		i--
		dAtA[i] = 0x22
	}
	if m.LatestEpoch != 0 {
		i = encodeVarintHarmony(dAtA, i, uint64(m.LatestEpoch))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintHarmony(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.ShardId != 0 {
		i = encodeVarintHarmony(dAtA, i, uint64(m.ShardId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ConsensusState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConsensusState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Root) > 0 {
		i -= len(m.Root)
		copy(dAtA[i:], m.Root)
		i = encodeVarintHarmony(dAtA, i, uint64(len(m.Root)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != 0 {
		i = encodeVarintHarmony(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EpochState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Committee) > 0 {
		i -= len(m.Committee)
		copy(dAtA[i:], m.Committee)
		i = encodeVarintHarmony(dAtA, i, uint64(len(m.Committee)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccountProof) > 0 {
		i -= len(m.AccountProof)
		copy(dAtA[i:], m.AccountProof)
		i = encodeVarintHarmony(dAtA, i, uint64(len(m.AccountProof)))
		i--
		dAtA[i] = 0x32
	}
	if m.CrossLinkIndex != 0 {
		i = encodeVarintHarmony(dAtA, i, uint64(m.CrossLinkIndex))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Bitmap) > 0 {
		i -= len(m.Bitmap)
		copy(dAtA[i:], m.Bitmap)
		i = encodeVarintHarmony(dAtA, i, uint64(len(m.Bitmap)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintHarmony(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ShardHeader) > 0 {
		i -= len(m.ShardHeader)
		copy(dAtA[i:], m.ShardHeader)
		i = encodeVarintHarmony(dAtA, i, uint64(len(m.ShardHeader)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BeaconHeader) > 0 {
		i -= len(m.BeaconHeader)
		copy(dAtA[i:], m.BeaconHeader)
		i = encodeVarintHarmony(dAtA, i, uint64(len(m.BeaconHeader)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHarmony(dAtA []byte, offset int, v uint64) int {
	offset -= sovHarmony(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ShardId != 0 {
		n += 1 + sovHarmony(uint64(m.ShardId))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovHarmony(uint64(l))
	}
	if m.LatestEpoch != 0 {
		n += 1 + sovHarmony(uint64(m.LatestEpoch))
	}
	l = len(m.LatestCommittee)
	if l > 0 {
		n += 1 + l + sovHarmony(uint64(l))
	}
	if m.LatestHeight != nil {
		l = m.LatestHeight.Size()
		n += 1 + l + sovHarmony(uint64(l))
	}
	return n
}

func (m *ConsensusState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovHarmony(uint64(m.Timestamp))
	}
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovHarmony(uint64(l))
	}
	return n
}

func (m *EpochState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Committee)
	if l > 0 {
		n += 1 + l + sovHarmony(uint64(l))
	}
	return n
}

func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BeaconHeader)
	if l > 0 {
		n += 1 + l + sovHarmony(uint64(l))
	}
	l = len(m.ShardHeader)
	if l > 0 {
		n += 1 + l + sovHarmony(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovHarmony(uint64(l))
	}
	l = len(m.Bitmap)
	if l > 0 {
		n += 1 + l + sovHarmony(uint64(l))
	}
	if m.CrossLinkIndex != 0 {
		n += 1 + sovHarmony(uint64(m.CrossLinkIndex))
	}
	l = len(m.AccountProof)
	if l > 0 {
		n += 1 + l + sovHarmony(uint64(l))
	}
	return n
}

func sovHarmony(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHarmony(x uint64) (n int) {
	return sovHarmony(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHarmony
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
			return fmt.Errorf("proto: ClientState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardId", wireType)
			}
			m.ShardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHarmony
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShardId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHarmony
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
				return ErrInvalidLengthHarmony
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHarmony
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = append(m.ContractAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ContractAddress == nil {
				m.ContractAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestEpoch", wireType)
			}
			m.LatestEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHarmony
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestCommittee", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHarmony
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
				return ErrInvalidLengthHarmony
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHarmony
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LatestCommittee = append(m.LatestCommittee[:0], dAtA[iNdEx:postIndex]...)
			if m.LatestCommittee == nil {
				m.LatestCommittee = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHarmony
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
				return ErrInvalidLengthHarmony
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHarmony
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LatestHeight == nil {
				m.LatestHeight = &types.Height{}
			}
			if err := m.LatestHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHarmony(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHarmony
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
func (m *ConsensusState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHarmony
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
			return fmt.Errorf("proto: ConsensusState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsensusState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHarmony
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHarmony
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
				return ErrInvalidLengthHarmony
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHarmony
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = append(m.Root[:0], dAtA[iNdEx:postIndex]...)
			if m.Root == nil {
				m.Root = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHarmony(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHarmony
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
func (m *EpochState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHarmony
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
			return fmt.Errorf("proto: EpochState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Committee", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHarmony
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
				return ErrInvalidLengthHarmony
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHarmony
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Committee = append(m.Committee[:0], dAtA[iNdEx:postIndex]...)
			if m.Committee == nil {
				m.Committee = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHarmony(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHarmony
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
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHarmony
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeaconHeader", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHarmony
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
				return ErrInvalidLengthHarmony
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHarmony
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BeaconHeader = append(m.BeaconHeader[:0], dAtA[iNdEx:postIndex]...)
			if m.BeaconHeader == nil {
				m.BeaconHeader = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardHeader", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHarmony
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
				return ErrInvalidLengthHarmony
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHarmony
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShardHeader = append(m.ShardHeader[:0], dAtA[iNdEx:postIndex]...)
			if m.ShardHeader == nil {
				m.ShardHeader = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHarmony
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
				return ErrInvalidLengthHarmony
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHarmony
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bitmap", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHarmony
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
				return ErrInvalidLengthHarmony
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHarmony
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bitmap = append(m.Bitmap[:0], dAtA[iNdEx:postIndex]...)
			if m.Bitmap == nil {
				m.Bitmap = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrossLinkIndex", wireType)
			}
			m.CrossLinkIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHarmony
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CrossLinkIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountProof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHarmony
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
				return ErrInvalidLengthHarmony
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHarmony
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountProof = append(m.AccountProof[:0], dAtA[iNdEx:postIndex]...)
			if m.AccountProof == nil {
				m.AccountProof = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHarmony(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHarmony
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
func skipHarmony(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHarmony
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
					return 0, ErrIntOverflowHarmony
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
					return 0, ErrIntOverflowHarmony
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
				return 0, ErrInvalidLengthHarmony
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHarmony
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHarmony
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHarmony        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHarmony          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHarmony = fmt.Errorf("proto: unexpected end of group")
)