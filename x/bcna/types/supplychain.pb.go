// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bcna/supplychain.proto

package types

import (
	fmt "fmt"
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

type Supplychain struct {
	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Product    string `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Info       string `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	Supplyinfo string `protobuf:"bytes,4,opt,name=supplyinfo,proto3" json:"supplyinfo,omitempty"`
	Creator    string `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Supplychain) Reset()         { *m = Supplychain{} }
func (m *Supplychain) String() string { return proto.CompactTextString(m) }
func (*Supplychain) ProtoMessage()    {}
func (*Supplychain) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6988f475c438d71, []int{0}
}
func (m *Supplychain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Supplychain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Supplychain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Supplychain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Supplychain.Merge(m, src)
}
func (m *Supplychain) XXX_Size() int {
	return m.Size()
}
func (m *Supplychain) XXX_DiscardUnknown() {
	xxx_messageInfo_Supplychain.DiscardUnknown(m)
}

var xxx_messageInfo_Supplychain proto.InternalMessageInfo

func (m *Supplychain) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Supplychain) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *Supplychain) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *Supplychain) GetSupplyinfo() string {
	if m != nil {
		return m.Supplyinfo
	}
	return ""
}

func (m *Supplychain) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Supplychain)(nil), "BitCannaGlobal.bcna.bcna.Supplychain")
}

func init() { proto.RegisterFile("bcna/supplychain.proto", fileDescriptor_b6988f475c438d71) }

var fileDescriptor_b6988f475c438d71 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x4a, 0xce, 0x4b,
	0xd4, 0x2f, 0x2e, 0x2d, 0x28, 0xc8, 0xa9, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x92, 0x70, 0xca, 0x2c, 0x71, 0x4e, 0xcc, 0xcb, 0x4b, 0x74, 0xcf, 0xc9, 0x4f,
	0x4a, 0xcc, 0xd1, 0x03, 0x29, 0x03, 0x13, 0x4a, 0xad, 0x8c, 0x5c, 0xdc, 0xc1, 0x08, 0xf5, 0x42,
	0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x4c, 0x99, 0x29, 0x42,
	0x12, 0x5c, 0xec, 0x05, 0x45, 0xf9, 0x29, 0xa5, 0xc9, 0x25, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x30, 0xae, 0x90, 0x10, 0x17, 0x4b, 0x66, 0x5e, 0x5a, 0xbe, 0x04, 0x33, 0x58, 0x18, 0xcc,
	0x16, 0x92, 0xe3, 0xe2, 0x82, 0x58, 0x0e, 0x96, 0x61, 0x01, 0xcb, 0x20, 0x89, 0x80, 0x4c, 0x4b,
	0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x85, 0x98, 0x06, 0xe5, 0x3a, 0xb9, 0x9e, 0x78,
	0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c,
	0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x76, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92,
	0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xaa, 0x37, 0xf4, 0xc1, 0xbe, 0xad, 0x80, 0x50, 0x25, 0x95, 0x05,
	0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xff, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xed, 0x20, 0xee,
	0xa2, 0x09, 0x01, 0x00, 0x00,
}

func (m *Supplychain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Supplychain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Supplychain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSupplychain(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Supplyinfo) > 0 {
		i -= len(m.Supplyinfo)
		copy(dAtA[i:], m.Supplyinfo)
		i = encodeVarintSupplychain(dAtA, i, uint64(len(m.Supplyinfo)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintSupplychain(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Product) > 0 {
		i -= len(m.Product)
		copy(dAtA[i:], m.Product)
		i = encodeVarintSupplychain(dAtA, i, uint64(len(m.Product)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSupplychain(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSupplychain(dAtA []byte, offset int, v uint64) int {
	offset -= sovSupplychain(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Supplychain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSupplychain(uint64(m.Id))
	}
	l = len(m.Product)
	if l > 0 {
		n += 1 + l + sovSupplychain(uint64(l))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovSupplychain(uint64(l))
	}
	l = len(m.Supplyinfo)
	if l > 0 {
		n += 1 + l + sovSupplychain(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSupplychain(uint64(l))
	}
	return n
}

func sovSupplychain(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSupplychain(x uint64) (n int) {
	return sovSupplychain(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Supplychain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSupplychain
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
			return fmt.Errorf("proto: Supplychain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Supplychain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplychain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Product", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplychain
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
				return ErrInvalidLengthSupplychain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSupplychain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Product = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplychain
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
				return ErrInvalidLengthSupplychain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSupplychain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supplyinfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplychain
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
				return ErrInvalidLengthSupplychain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSupplychain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Supplyinfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplychain
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
				return ErrInvalidLengthSupplychain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSupplychain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSupplychain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSupplychain
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
func skipSupplychain(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSupplychain
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
					return 0, ErrIntOverflowSupplychain
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
					return 0, ErrIntOverflowSupplychain
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
				return 0, ErrInvalidLengthSupplychain
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSupplychain
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSupplychain
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSupplychain        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSupplychain          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSupplychain = fmt.Errorf("proto: unexpected end of group")
)
