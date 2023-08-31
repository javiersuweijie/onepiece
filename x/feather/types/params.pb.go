// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feather/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/terra-money/alliance/x/alliance/types"
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

type Params struct {
	HaltIfNoChannel    bool                            `protobuf:"varint,1,opt,name=halt_if_no_channel,json=haltIfNoChannel,proto3" json:"halt_if_no_channel,omitempty"`
	BaseDenom          string                          `protobuf:"bytes,2,opt,name=base_denom,json=baseDenom,proto3" json:"base_denom,omitempty"`
	BaseChainId        string                          `protobuf:"bytes,3,opt,name=base_chain_id,json=baseChainId,proto3" json:"base_chain_id,omitempty"`
	AllianceBondHeight int64                           `protobuf:"varint,4,opt,name=alliance_bond_height,json=allianceBondHeight,proto3" json:"alliance_bond_height,omitempty"`
	Alliance           types.MsgCreateAllianceProposal `protobuf:"bytes,5,opt,name=alliance,proto3" json:"alliance"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_533ed41b0bf25026, []int{0}
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

func (m *Params) GetHaltIfNoChannel() bool {
	if m != nil {
		return m.HaltIfNoChannel
	}
	return false
}

func (m *Params) GetBaseDenom() string {
	if m != nil {
		return m.BaseDenom
	}
	return ""
}

func (m *Params) GetBaseChainId() string {
	if m != nil {
		return m.BaseChainId
	}
	return ""
}

func (m *Params) GetAllianceBondHeight() int64 {
	if m != nil {
		return m.AllianceBondHeight
	}
	return 0
}

func (m *Params) GetAlliance() types.MsgCreateAllianceProposal {
	if m != nil {
		return m.Alliance
	}
	return types.MsgCreateAllianceProposal{}
}

func init() {
	proto.RegisterType((*Params)(nil), "feather.Params")
}

func init() { proto.RegisterFile("feather/params.proto", fileDescriptor_533ed41b0bf25026) }

var fileDescriptor_533ed41b0bf25026 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xd0, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0x07, 0xf0, 0xc6, 0xcd, 0xb9, 0x65, 0x88, 0x10, 0x76, 0x28, 0x03, 0x6b, 0x99, 0x97, 0x82,
	0xac, 0x11, 0xf7, 0x04, 0x6e, 0x0a, 0xee, 0x30, 0x19, 0x3d, 0x7a, 0x09, 0x69, 0x9b, 0x35, 0x85,
	0x36, 0x5f, 0x49, 0xa3, 0xb8, 0xb7, 0xf0, 0xb1, 0x76, 0xdc, 0xd1, 0x93, 0xc8, 0xf6, 0x04, 0xbe,
	0x81, 0xb4, 0x5b, 0x77, 0xfb, 0xf2, 0xfb, 0x27, 0x21, 0xf9, 0xe3, 0xc1, 0x4a, 0x70, 0x23, 0x85,
	0xa6, 0x05, 0xd7, 0x3c, 0x2f, 0xfd, 0x42, 0x83, 0x01, 0x72, 0x71, 0xd4, 0x21, 0xe1, 0x59, 0x96,
	0x72, 0x15, 0x09, 0x9a, 0xc0, 0xc7, 0x21, 0x1c, 0x0e, 0x12, 0x48, 0xa0, 0x1e, 0x69, 0x35, 0x1d,
	0x74, 0xf4, 0x87, 0x70, 0x67, 0x59, 0xdf, 0x41, 0xee, 0x30, 0x91, 0x3c, 0x33, 0x2c, 0x5d, 0x31,
	0x05, 0x2c, 0x92, 0x5c, 0x29, 0x91, 0xd9, 0xc8, 0x45, 0x5e, 0x37, 0xb8, 0xaa, 0x92, 0xf9, 0xea,
	0x15, 0x66, 0x07, 0x26, 0xd7, 0x18, 0x87, 0xbc, 0x14, 0x2c, 0x16, 0x0a, 0x72, 0xfb, 0xcc, 0x45,
	0x5e, 0x2f, 0xe8, 0x55, 0xf2, 0x54, 0x01, 0x19, 0xe1, 0xcb, 0x3a, 0x8e, 0x24, 0x4f, 0x15, 0x4b,
	0x63, 0xbb, 0x55, 0xef, 0xe8, 0x57, 0x38, 0xab, 0x6c, 0x1e, 0x93, 0x7b, 0x3c, 0x68, 0x9e, 0xc9,
	0x42, 0x50, 0x31, 0x93, 0x22, 0x4d, 0xa4, 0xb1, 0xdb, 0x2e, 0xf2, 0x5a, 0xc1, 0xe9, 0x0b, 0x53,
	0x50, 0xf1, 0x4b, 0x9d, 0x90, 0x67, 0xdc, 0x6d, 0xd4, 0x3e, 0x77, 0x91, 0xd7, 0x7f, 0xb8, 0xf5,
	0x1b, 0xf0, 0x17, 0x65, 0x32, 0xd3, 0x82, 0x1b, 0xf1, 0x78, 0x94, 0xa5, 0x86, 0x02, 0x4a, 0x9e,
	0x4d, 0xdb, 0x9b, 0x9f, 0x1b, 0x2b, 0x38, 0x1d, 0x9d, 0x2e, 0x36, 0x3b, 0x07, 0x6d, 0x77, 0x0e,
	0xfa, 0xdd, 0x39, 0xe8, 0x6b, 0xef, 0x58, 0xdb, 0xbd, 0x63, 0x7d, 0xef, 0x1d, 0xeb, 0x6d, 0x92,
	0xa4, 0x46, 0xbe, 0x87, 0x7e, 0x04, 0x39, 0x35, 0x42, 0x6b, 0x3e, 0xce, 0x41, 0x89, 0x35, 0x3d,
	0xf6, 0x3a, 0x8e, 0x40, 0x0b, 0xfa, 0xd9, 0x2c, 0xa9, 0x59, 0x17, 0xa2, 0x0c, 0x3b, 0x75, 0x93,
	0x93, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x4d, 0x54, 0x15, 0x94, 0x01, 0x00, 0x00,
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
		size, err := m.Alliance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.AllianceBondHeight != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AllianceBondHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.BaseChainId) > 0 {
		i -= len(m.BaseChainId)
		copy(dAtA[i:], m.BaseChainId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.BaseChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BaseDenom) > 0 {
		i -= len(m.BaseDenom)
		copy(dAtA[i:], m.BaseDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.BaseDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.HaltIfNoChannel {
		i--
		if m.HaltIfNoChannel {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
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
	if m.HaltIfNoChannel {
		n += 2
	}
	l = len(m.BaseDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.BaseChainId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.AllianceBondHeight != 0 {
		n += 1 + sovParams(uint64(m.AllianceBondHeight))
	}
	l = m.Alliance.Size()
	n += 1 + l + sovParams(uint64(l))
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
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HaltIfNoChannel", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.HaltIfNoChannel = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenom", wireType)
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
			m.BaseDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseChainId", wireType)
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
			m.BaseChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllianceBondHeight", wireType)
			}
			m.AllianceBondHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllianceBondHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alliance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Alliance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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