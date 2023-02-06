// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pairing/params.proto

package types

import (
	fmt "fmt"
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

// Params defines the parameters for the module.
type Params struct {
	MintCoinsPerCU                      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=mintCoinsPerCU,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"mintCoinsPerCU" yaml:"mint_coins_per_cu"`
	BurnCoinsPerCU                      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=burnCoinsPerCU,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"burnCoinsPerCU" yaml:"burn_coins_per_cu"`
	FraudStakeSlashingFactor            github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=fraudStakeSlashingFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fraudStakeSlashingFactor" yaml:"fraud_stake_slashing_factor"`
	FraudSlashingAmount                 uint64                                 `protobuf:"varint,6,opt,name=fraudSlashingAmount,proto3" json:"fraudSlashingAmount,omitempty" yaml:"fraud_slashing_amount"`
	ServicersToPairCount                uint64                                 `protobuf:"varint,7,opt,name=servicersToPairCount,proto3" json:"servicersToPairCount,omitempty" yaml:"servicers_to_pair_count"`
	EpochBlocksOverlap                  uint64                                 `protobuf:"varint,8,opt,name=epochBlocksOverlap,proto3" json:"epochBlocksOverlap,omitempty" yaml:"epoch_blocks_overlap"`
	StakeToMaxCUList                    StakeToMaxCUList                       `protobuf:"bytes,9,opt,name=stakeToMaxCUList,proto3,customtype=StakeToMaxCUList" json:"stakeToMaxCUList" yaml:"stake_to_computeunits_list"`
	UnpayLimit                          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=unpayLimit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"unpayLimit" yaml:"unpay_limit"`
	SlashLimit                          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=slashLimit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slashLimit" yaml:"slash_limit"`
	DataReliabilityReward               github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,12,opt,name=dataReliabilityReward,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"dataReliabilityReward" yaml:"data_reliability_reward"`
	QoSWeight                           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,13,opt,name=QoSWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"QoSWeight" yaml:"data_reliability_reward"`
	RecommendedEpochNumToCollectPayment uint64                                 `protobuf:"varint,14,opt,name=recommendedEpochNumToCollectPayment,proto3" json:"recommendedEpochNumToCollectPayment,omitempty" yaml:"recommended_epoch_num_to_collect_payment"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_72cc734580d3bc3a, []int{0}
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

func (m *Params) GetFraudSlashingAmount() uint64 {
	if m != nil {
		return m.FraudSlashingAmount
	}
	return 0
}

func (m *Params) GetServicersToPairCount() uint64 {
	if m != nil {
		return m.ServicersToPairCount
	}
	return 0
}

func (m *Params) GetEpochBlocksOverlap() uint64 {
	if m != nil {
		return m.EpochBlocksOverlap
	}
	return 0
}

func (m *Params) GetRecommendedEpochNumToCollectPayment() uint64 {
	if m != nil {
		return m.RecommendedEpochNumToCollectPayment
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "lavanet.lava.pairing.Params")
}

func init() { proto.RegisterFile("pairing/params.proto", fileDescriptor_72cc734580d3bc3a) }

var fileDescriptor_72cc734580d3bc3a = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x41, 0x6f, 0xd3, 0x3c,
	0x18, 0xc7, 0x9b, 0xad, 0xef, 0xde, 0xcd, 0xef, 0xcb, 0x54, 0x85, 0x22, 0x45, 0x80, 0x92, 0x11,
	0x24, 0xd8, 0x85, 0xe6, 0xb0, 0xdb, 0x24, 0x0e, 0x6b, 0x07, 0x87, 0x69, 0xb0, 0x92, 0x76, 0x20,
	0x71, 0xb1, 0x5c, 0xc7, 0x6b, 0xad, 0xc5, 0x76, 0x64, 0x3b, 0x63, 0xfd, 0x00, 0x70, 0xde, 0x91,
	0x23, 0x1f, 0x67, 0xe2, 0xb4, 0x23, 0xe2, 0x10, 0xa1, 0xed, 0x1b, 0xf4, 0x13, 0xa0, 0xd8, 0xd9,
	0xd6, 0x8d, 0x22, 0x31, 0x4d, 0x9c, 0x5c, 0xd5, 0xff, 0xe7, 0xf7, 0x8b, 0x9e, 0xe7, 0x91, 0x41,
	0x33, 0x43, 0x54, 0x52, 0x3e, 0x8c, 0x32, 0x24, 0x11, 0x53, 0xad, 0x4c, 0x0a, 0x2d, 0xdc, 0x66,
	0x8a, 0x0e, 0x10, 0x27, 0xba, 0x55, 0x9e, 0xad, 0x2a, 0x72, 0xbf, 0x39, 0x14, 0x43, 0x61, 0x02,
	0x51, 0xf9, 0xcb, 0x66, 0xc3, 0xaf, 0x00, 0x2c, 0x74, 0x4d, 0xb1, 0x2b, 0xc1, 0x32, 0xa3, 0x5c,
	0x77, 0x04, 0xe5, 0xaa, 0x4b, 0x64, 0x67, 0xd7, 0x9b, 0x5f, 0x71, 0x56, 0x97, 0xda, 0x5b, 0xc7,
	0x45, 0x50, 0xfb, 0x5e, 0x04, 0x4f, 0x86, 0x54, 0x8f, 0xf2, 0x41, 0x0b, 0x0b, 0x16, 0x61, 0xa1,
	0x98, 0x50, 0xd5, 0xf1, 0x4c, 0x25, 0xfb, 0x91, 0x1e, 0x67, 0x44, 0xb5, 0x36, 0x09, 0x9e, 0x14,
	0x81, 0x37, 0x46, 0x2c, 0x5d, 0x0f, 0x4b, 0x1a, 0xc4, 0x25, 0x0e, 0x66, 0x44, 0x42, 0x9c, 0x87,
	0xf1, 0x35, 0x43, 0xe9, 0x1c, 0xe4, 0x92, 0x4f, 0x39, 0xeb, 0xb7, 0x73, 0x96, 0xb4, 0xeb, 0xce,
	0xab, 0x06, 0xf7, 0xc8, 0x01, 0xde, 0x9e, 0x44, 0x79, 0xd2, 0xd3, 0x68, 0x9f, 0xf4, 0x52, 0xa4,
	0x46, 0x94, 0x0f, 0x5f, 0x22, 0xac, 0x85, 0xf4, 0xfe, 0x31, 0xfa, 0xfe, 0x8d, 0xf5, 0xa1, 0xd5,
	0x1b, 0x2e, 0x54, 0x25, 0x18, 0xaa, 0x8a, 0x0c, 0xf7, 0x0c, 0x3a, 0x8c, 0x7f, 0x6b, 0x75, 0x63,
	0x70, 0xd7, 0xde, 0x55, 0x7f, 0x6f, 0x30, 0x91, 0x73, 0xed, 0x2d, 0xac, 0x38, 0xab, 0xf5, 0xf6,
	0xca, 0xa4, 0x08, 0x1e, 0x5e, 0xc1, 0x9f, 0x83, 0x91, 0x89, 0x85, 0xf1, 0xac, 0x62, 0xf7, 0x2d,
	0x68, 0x2a, 0x22, 0x0f, 0x28, 0x26, 0x52, 0xf5, 0x45, 0x17, 0x51, 0xd9, 0x31, 0xd0, 0x7f, 0x0d,
	0x34, 0x9c, 0x14, 0x81, 0x6f, 0xa1, 0x17, 0x29, 0xa8, 0x05, 0x2c, 0xb7, 0x05, 0x62, 0x8b, 0x9d,
	0x59, 0xef, 0xee, 0x00, 0x97, 0x64, 0x02, 0x8f, 0xda, 0xa9, 0xc0, 0xfb, 0x6a, 0xe7, 0x80, 0xc8,
	0x14, 0x65, 0xde, 0xa2, 0xa1, 0x06, 0x93, 0x22, 0x78, 0x60, 0xa9, 0x26, 0x03, 0x07, 0x26, 0x04,
	0x85, 0x4d, 0x85, 0xf1, 0x8c, 0x52, 0x97, 0x82, 0x86, 0x69, 0x58, 0x5f, 0xbc, 0x42, 0x87, 0x9d,
	0xdd, 0x6d, 0xaa, 0xb4, 0xb7, 0x64, 0xc6, 0xf0, 0xbc, 0x1a, 0x43, 0xa3, 0x77, 0xed, 0x7e, 0x52,
	0x04, 0x8f, 0xaa, 0x8f, 0x37, 0xad, 0xd6, 0x02, 0x62, 0xc1, 0xb2, 0x5c, 0x93, 0x9c, 0x53, 0xad,
	0x60, 0x4a, 0x95, 0x0e, 0xe3, 0x5f, 0xb0, 0x6e, 0x02, 0x40, 0xce, 0x33, 0x34, 0xde, 0xa6, 0x8c,
	0x6a, 0x0f, 0x18, 0xc9, 0xe6, 0x8d, 0x67, 0xed, 0x5a, 0xb5, 0x21, 0xc1, 0xb4, 0x44, 0x85, 0xf1,
	0x14, 0xb7, 0xb4, 0x98, 0x11, 0x59, 0xcb, 0x7f, 0xb7, 0xb3, 0x18, 0xd2, 0x85, 0xe5, 0x92, 0xeb,
	0x7e, 0x72, 0xc0, 0xbd, 0x04, 0x69, 0x14, 0x93, 0x94, 0xa2, 0x01, 0x4d, 0xa9, 0x1e, 0xc7, 0xe4,
	0x03, 0x92, 0x89, 0xf7, 0xbf, 0x31, 0x76, 0x6f, 0x6c, 0xac, 0xf6, 0xa1, 0x84, 0x42, 0x79, 0x49,
	0x85, 0xd2, 0x60, 0xc3, 0x78, 0xb6, 0xce, 0xe5, 0x60, 0xe9, 0x8d, 0xe8, 0xbd, 0x23, 0x74, 0x38,
	0xd2, 0xde, 0x9d, 0xbf, 0xe4, 0xbe, 0x54, 0xb8, 0x1f, 0x1d, 0xf0, 0x58, 0x12, 0x2c, 0x18, 0x23,
	0x3c, 0x21, 0xc9, 0x8b, 0x72, 0xa3, 0x5e, 0xe7, 0xac, 0x2f, 0x3a, 0x22, 0x4d, 0x09, 0xd6, 0x5d,
	0x34, 0x66, 0x84, 0x6b, 0x6f, 0xd9, 0xac, 0xe4, 0xda, 0xa4, 0x08, 0x22, 0x0b, 0x9f, 0x2a, 0x82,
	0x76, 0x3d, 0x79, 0xce, 0xec, 0xee, 0x98, 0x42, 0x98, 0xd9, 0xca, 0x30, 0xfe, 0x13, 0xfe, 0x7a,
	0xfd, 0xf3, 0x97, 0xa0, 0xb6, 0x55, 0x5f, 0x74, 0x1a, 0x73, 0x5b, 0xf5, 0xc5, 0xb9, 0xc6, 0x7c,
	0x7b, 0xe3, 0xf8, 0xd4, 0x77, 0x4e, 0x4e, 0x7d, 0xe7, 0xc7, 0xa9, 0xef, 0x1c, 0x9d, 0xf9, 0xb5,
	0x93, 0x33, 0xbf, 0xf6, 0xed, 0xcc, 0xaf, 0xbd, 0x7f, 0x3a, 0xd5, 0x88, 0xea, 0x75, 0x36, 0x67,
	0x74, 0x18, 0x9d, 0x3f, 0xe1, 0xa6, 0x1b, 0x83, 0x05, 0xf3, 0x2c, 0xaf, 0xfd, 0x0c, 0x00, 0x00,
	0xff, 0xff, 0xa8, 0x01, 0x54, 0x8a, 0xda, 0x05, 0x00, 0x00,
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
	if m.RecommendedEpochNumToCollectPayment != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RecommendedEpochNumToCollectPayment))
		i--
		dAtA[i] = 0x70
	}
	{
		size := m.QoSWeight.Size()
		i -= size
		if _, err := m.QoSWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.DataReliabilityReward.Size()
		i -= size
		if _, err := m.DataReliabilityReward.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.SlashLimit.Size()
		i -= size
		if _, err := m.SlashLimit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.UnpayLimit.Size()
		i -= size
		if _, err := m.UnpayLimit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.StakeToMaxCUList.Size()
		i -= size
		if _, err := m.StakeToMaxCUList.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.EpochBlocksOverlap != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EpochBlocksOverlap))
		i--
		dAtA[i] = 0x40
	}
	if m.ServicersToPairCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ServicersToPairCount))
		i--
		dAtA[i] = 0x38
	}
	if m.FraudSlashingAmount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.FraudSlashingAmount))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.FraudStakeSlashingFactor.Size()
		i -= size
		if _, err := m.FraudStakeSlashingFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.BurnCoinsPerCU.Size()
		i -= size
		if _, err := m.BurnCoinsPerCU.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.MintCoinsPerCU.Size()
		i -= size
		if _, err := m.MintCoinsPerCU.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MintCoinsPerCU.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.BurnCoinsPerCU.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.FraudStakeSlashingFactor.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.FraudSlashingAmount != 0 {
		n += 1 + sovParams(uint64(m.FraudSlashingAmount))
	}
	if m.ServicersToPairCount != 0 {
		n += 1 + sovParams(uint64(m.ServicersToPairCount))
	}
	if m.EpochBlocksOverlap != 0 {
		n += 1 + sovParams(uint64(m.EpochBlocksOverlap))
	}
	l = m.StakeToMaxCUList.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.UnpayLimit.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashLimit.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.DataReliabilityReward.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.QoSWeight.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.RecommendedEpochNumToCollectPayment != 0 {
		n += 1 + sovParams(uint64(m.RecommendedEpochNumToCollectPayment))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintCoinsPerCU", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MintCoinsPerCU.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnCoinsPerCU", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurnCoinsPerCU.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FraudStakeSlashingFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FraudStakeSlashingFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FraudSlashingAmount", wireType)
			}
			m.FraudSlashingAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FraudSlashingAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServicersToPairCount", wireType)
			}
			m.ServicersToPairCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServicersToPairCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochBlocksOverlap", wireType)
			}
			m.EpochBlocksOverlap = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochBlocksOverlap |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeToMaxCUList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakeToMaxCUList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnpayLimit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UnpayLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashLimit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataReliabilityReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DataReliabilityReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QoSWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.QoSWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecommendedEpochNumToCollectPayment", wireType)
			}
			m.RecommendedEpochNumToCollectPayment = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RecommendedEpochNumToCollectPayment |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
