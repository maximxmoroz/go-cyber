// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cyber/resources/v1beta1/types.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	MaxSlots                   uint32                                  `protobuf:"varint,1,opt,name=max_slots,json=maxSlots,proto3" json:"max_slots,omitempty"`
	HalvingPeriodVoltBlocks    uint32                                  `protobuf:"varint,2,opt,name=halving_period_volt_blocks,json=halvingPeriodVoltBlocks,proto3" json:"halving_period_volt_blocks,omitempty"`
	HalvingPeriodAmpereBlocks  uint32                                  `protobuf:"varint,3,opt,name=halving_period_ampere_blocks,json=halvingPeriodAmpereBlocks,proto3" json:"halving_period_ampere_blocks,omitempty"`
	BaseInvestmintPeriodVolt   uint32                                  `protobuf:"varint,4,opt,name=base_investmint_period_volt,json=baseInvestmintPeriodVolt,proto3" json:"base_investmint_period_volt,omitempty"`
	BaseInvestmintPeriodAmpere uint32                                  `protobuf:"varint,5,opt,name=base_investmint_period_ampere,json=baseInvestmintPeriodAmpere,proto3" json:"base_investmint_period_ampere,omitempty"`
	MinInvestmintPeriod        uint32                                  `protobuf:"varint,6,opt,name=min_investmint_period,json=minInvestmintPeriod,proto3" json:"min_investmint_period,omitempty"`
	BaseInvestmintAmountVolt   github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,7,opt,name=base_investmint_amount_volt,json=baseInvestmintAmountVolt,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"base_investmint_amount_volt"`
	BaseInvestmintAmountAmpere github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,8,opt,name=base_investmint_amount_ampere,json=baseInvestmintAmountAmpere,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"base_investmint_amount_ampere"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_3be852646b47c447, []int{0}
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

func (m *Params) GetMaxSlots() uint32 {
	if m != nil {
		return m.MaxSlots
	}
	return 0
}

func (m *Params) GetHalvingPeriodVoltBlocks() uint32 {
	if m != nil {
		return m.HalvingPeriodVoltBlocks
	}
	return 0
}

func (m *Params) GetHalvingPeriodAmpereBlocks() uint32 {
	if m != nil {
		return m.HalvingPeriodAmpereBlocks
	}
	return 0
}

func (m *Params) GetBaseInvestmintPeriodVolt() uint32 {
	if m != nil {
		return m.BaseInvestmintPeriodVolt
	}
	return 0
}

func (m *Params) GetBaseInvestmintPeriodAmpere() uint32 {
	if m != nil {
		return m.BaseInvestmintPeriodAmpere
	}
	return 0
}

func (m *Params) GetMinInvestmintPeriod() uint32 {
	if m != nil {
		return m.MinInvestmintPeriod
	}
	return 0
}

func (m *Params) GetBaseInvestmintAmountVolt() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.BaseInvestmintAmountVolt
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *Params) GetBaseInvestmintAmountAmpere() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.BaseInvestmintAmountAmpere
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func init() {
	proto.RegisterType((*Params)(nil), "cyber.resources.v1beta1.Params")
}

func init() {
	proto.RegisterFile("cyber/resources/v1beta1/types.proto", fileDescriptor_3be852646b47c447)
}

var fileDescriptor_3be852646b47c447 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0x63, 0x38, 0xca, 0x61, 0xc4, 0x12, 0x40, 0x97, 0xcb, 0x41, 0xee, 0x04, 0x03, 0xb7,
	0x5c, 0xac, 0xbb, 0x1b, 0x11, 0x42, 0x3d, 0x26, 0x06, 0xa4, 0x0a, 0x24, 0x06, 0x96, 0xc8, 0x49,
	0xad, 0xd4, 0x6a, 0xec, 0x37, 0xb2, 0xdd, 0xa8, 0xfd, 0x08, 0x0c, 0x48, 0x7c, 0xac, 0x8e, 0x1d,
	0x99, 0x2a, 0xd4, 0x7e, 0x06, 0x16, 0x26, 0x14, 0xdb, 0x2d, 0xa5, 0x2d, 0xdb, 0x4d, 0x49, 0xe4,
	0xdf, 0xf3, 0xf8, 0xc9, 0xfb, 0x07, 0xbf, 0x2c, 0x26, 0x39, 0x53, 0x44, 0x31, 0x0d, 0x23, 0x55,
	0x30, 0x4d, 0x9a, 0xcb, 0x9c, 0x19, 0x7a, 0x49, 0xcc, 0xa4, 0x66, 0x3a, 0xad, 0x15, 0x18, 0x08,
	0x8f, 0x2c, 0x94, 0xae, 0xa1, 0xd4, 0x43, 0xf1, 0x93, 0x12, 0x4a, 0xb0, 0x0c, 0x69, 0xdf, 0x1c,
	0x1e, 0x27, 0x05, 0x68, 0x01, 0x9a, 0xe4, 0x54, 0xb3, 0xb5, 0x5f, 0x01, 0x5c, 0xba, 0xf3, 0x17,
	0xbf, 0x0e, 0x70, 0xa7, 0x47, 0x15, 0x15, 0x3a, 0x3c, 0xc1, 0x0f, 0x04, 0x1d, 0x67, 0xba, 0x02,
	0xa3, 0x23, 0x74, 0x86, 0xce, 0x1f, 0x7d, 0x3c, 0x14, 0x74, 0xfc, 0xa9, 0xfd, 0x0e, 0x5f, 0xe3,
	0x78, 0x40, 0xab, 0x86, 0xcb, 0x32, 0xab, 0x99, 0xe2, 0xd0, 0xcf, 0x1a, 0xa8, 0x4c, 0x96, 0x57,
	0x50, 0x0c, 0x75, 0x74, 0xc7, 0xd2, 0x47, 0x9e, 0xe8, 0x59, 0xe0, 0x33, 0x54, 0xe6, 0xc6, 0x1e,
	0x87, 0x6f, 0xf1, 0xb3, 0x2d, 0x31, 0x15, 0x35, 0x53, 0x6c, 0x25, 0xbf, 0x6b, 0xe5, 0xc7, 0xff,
	0xc8, 0xbb, 0x96, 0xf0, 0x06, 0x6f, 0xf0, 0x49, 0xfb, 0x03, 0x19, 0x97, 0x0d, 0xd3, 0x46, 0x70,
	0x69, 0x36, 0x53, 0x44, 0x07, 0x56, 0x1f, 0xb5, 0xc8, 0xfb, 0x35, 0xf1, 0x37, 0x45, 0xd8, 0xc5,
	0xcf, 0xff, 0x23, 0x77, 0x39, 0xa2, 0x7b, 0xd6, 0x20, 0xde, 0x67, 0xe0, 0x72, 0x84, 0x57, 0xf8,
	0xa9, 0xe0, 0x72, 0xd7, 0x21, 0xea, 0x58, 0xe9, 0x63, 0xc1, 0xe5, 0xb6, 0x32, 0xfc, 0x8a, 0x76,
	0x63, 0x53, 0x01, 0x23, 0x69, 0x5c, 0xec, 0xfb, 0x67, 0xe8, 0xfc, 0xe1, 0xd5, 0x71, 0xea, 0x5a,
	0x94, 0xb6, 0xe8, 0xaa, 0x9b, 0xe9, 0x3b, 0xe0, 0xf2, 0x86, 0x4c, 0xe7, 0xa7, 0xc1, 0xef, 0xf9,
	0xe9, 0xab, 0x92, 0x9b, 0xc1, 0x28, 0x4f, 0x0b, 0x10, 0xc4, 0xf7, 0xd3, 0x3d, 0x2e, 0x74, 0x7f,
	0xe8, 0xa7, 0xa3, 0x15, 0x6c, 0x97, 0xa0, 0x6b, 0x2f, 0xb3, 0x25, 0xf8, 0x86, 0x76, 0x6b, 0xe0,
	0xb3, 0xf8, 0x1a, 0x1c, 0xde, 0x7a, 0x9a, 0x78, 0x5f, 0x1a, 0xdf, 0xd7, 0x0f, 0xd3, 0x45, 0x82,
	0x66, 0x8b, 0x04, 0xfd, 0x5c, 0x24, 0xe8, 0xfb, 0x32, 0x09, 0x66, 0xcb, 0x24, 0xf8, 0xb1, 0x4c,
	0x82, 0x2f, 0xd7, 0x9b, 0xf6, 0xed, 0xac, 0x17, 0x20, 0x4b, 0xc5, 0xb4, 0x26, 0x25, 0x5c, 0xb8,
	0x0d, 0x19, 0x6f, 0xec, 0x88, 0xbd, 0x2f, 0xef, 0xd8, 0x69, 0xbe, 0xfe, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x42, 0x8b, 0x46, 0xd6, 0x43, 0x03, 0x00, 0x00,
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
	{
		size, err := m.BaseInvestmintAmountAmpere.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size, err := m.BaseInvestmintAmountVolt.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.MinInvestmintPeriod != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MinInvestmintPeriod))
		i--
		dAtA[i] = 0x30
	}
	if m.BaseInvestmintPeriodAmpere != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.BaseInvestmintPeriodAmpere))
		i--
		dAtA[i] = 0x28
	}
	if m.BaseInvestmintPeriodVolt != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.BaseInvestmintPeriodVolt))
		i--
		dAtA[i] = 0x20
	}
	if m.HalvingPeriodAmpereBlocks != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.HalvingPeriodAmpereBlocks))
		i--
		dAtA[i] = 0x18
	}
	if m.HalvingPeriodVoltBlocks != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.HalvingPeriodVoltBlocks))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxSlots != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MaxSlots))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
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
	if m.MaxSlots != 0 {
		n += 1 + sovTypes(uint64(m.MaxSlots))
	}
	if m.HalvingPeriodVoltBlocks != 0 {
		n += 1 + sovTypes(uint64(m.HalvingPeriodVoltBlocks))
	}
	if m.HalvingPeriodAmpereBlocks != 0 {
		n += 1 + sovTypes(uint64(m.HalvingPeriodAmpereBlocks))
	}
	if m.BaseInvestmintPeriodVolt != 0 {
		n += 1 + sovTypes(uint64(m.BaseInvestmintPeriodVolt))
	}
	if m.BaseInvestmintPeriodAmpere != 0 {
		n += 1 + sovTypes(uint64(m.BaseInvestmintPeriodAmpere))
	}
	if m.MinInvestmintPeriod != 0 {
		n += 1 + sovTypes(uint64(m.MinInvestmintPeriod))
	}
	l = m.BaseInvestmintAmountVolt.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.BaseInvestmintAmountAmpere.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSlots", wireType)
			}
			m.MaxSlots = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSlots |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HalvingPeriodVoltBlocks", wireType)
			}
			m.HalvingPeriodVoltBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HalvingPeriodVoltBlocks |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HalvingPeriodAmpereBlocks", wireType)
			}
			m.HalvingPeriodAmpereBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HalvingPeriodAmpereBlocks |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseInvestmintPeriodVolt", wireType)
			}
			m.BaseInvestmintPeriodVolt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BaseInvestmintPeriodVolt |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseInvestmintPeriodAmpere", wireType)
			}
			m.BaseInvestmintPeriodAmpere = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BaseInvestmintPeriodAmpere |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInvestmintPeriod", wireType)
			}
			m.MinInvestmintPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinInvestmintPeriod |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseInvestmintAmountVolt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseInvestmintAmountVolt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseInvestmintAmountAmpere", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseInvestmintAmountAmpere.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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

func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
