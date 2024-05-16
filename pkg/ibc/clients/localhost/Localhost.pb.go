// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: clients/localhost/Localhost.proto

package localhost

import (
	fmt "fmt"
	_ "github.com/datachainlab/solidity-protobuf/protobuf-solidity/src/protoc/go"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	client "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/core/client"
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
	// the latest block height
	LatestHeight client.Height `protobuf:"bytes,1,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height"`
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3c78b2970c484fc, []int{0}
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

func init() {
	proto.RegisterType((*ClientState)(nil), "ibc.lightclients.localhost.v2.ClientState")
}

func init() { proto.RegisterFile("clients/localhost/Localhost.proto", fileDescriptor_b3c78b2970c484fc) }

var fileDescriptor_b3c78b2970c484fc = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x3f, 0x4e, 0xf4, 0x30,
	0x10, 0xc5, 0x63, 0xe9, 0xd3, 0x87, 0x94, 0x85, 0x66, 0x45, 0xb1, 0x5a, 0x84, 0x81, 0xad, 0x68,
	0x6c, 0xa3, 0xc0, 0x01, 0xf8, 0x23, 0x21, 0x0a, 0x2a, 0xe8, 0x68, 0x50, 0xec, 0x35, 0x8e, 0x85,
	0xc9, 0x44, 0xf1, 0x04, 0x91, 0x96, 0x13, 0x70, 0x00, 0xae, 0xc1, 0x1d, 0x52, 0x6e, 0x49, 0x85,
	0x20, 0xb9, 0x08, 0x5a, 0x27, 0x4b, 0x43, 0x41, 0xe5, 0x79, 0xf6, 0xcf, 0xef, 0x3d, 0x4d, 0xbc,
	0xa7, 0x9c, 0xd5, 0x39, 0x7a, 0xe1, 0x40, 0xa5, 0x2e, 0x03, 0x8f, 0xe2, 0x72, 0x35, 0xf1, 0xa2,
	0x04, 0x84, 0xf1, 0xb6, 0x95, 0x8a, 0x3b, 0x6b, 0x32, 0x1c, 0x58, 0xfe, 0xc3, 0xf2, 0xc7, 0x64,
	0xba, 0x69, 0xc0, 0x40, 0x20, 0xc5, 0x72, 0xea, 0x3f, 0x4d, 0x67, 0x1e, 0x9c, 0x9d, 0x5b, 0xac,
	0x59, 0xd0, 0xb2, 0xba, 0x63, 0xfa, 0x09, 0x75, 0xee, 0x2d, 0xe4, 0x7e, 0x60, 0xb6, 0x14, 0x94,
	0x5a, 0x1c, 0x24, 0xac, 0xf7, 0x15, 0x67, 0xe1, 0xe8, 0x1f, 0x67, 0x27, 0xf1, 0xa8, 0xd7, 0xd7,
	0x98, 0xa2, 0x1e, 0x27, 0xf1, 0x86, 0x4b, 0x51, 0x7b, 0xbc, 0xcd, 0xf4, 0xb2, 0xca, 0x84, 0xec,
	0x92, 0xfd, 0x51, 0xb2, 0xc6, 0x2f, 0x82, 0x3c, 0xfd, 0xd7, 0x7c, 0xec, 0x44, 0x57, 0xeb, 0x3d,
	0x33, 0xdc, 0xbd, 0x92, 0xe7, 0xb7, 0xc9, 0x51, 0x9c, 0x1c, 0x67, 0x75, 0xa1, 0x4b, 0xa7, 0xe7,
	0x46, 0x97, 0xcc, 0xa5, 0xd2, 0x8b, 0xba, 0xb2, 0xcc, 0x4a, 0xc5, 0x56, 0x25, 0x85, 0x82, 0x1c,
	0xcb, 0x54, 0xa1, 0x17, 0x21, 0xbd, 0xf9, 0xa2, 0x51, 0xd3, 0x52, 0xb2, 0x68, 0x29, 0xf9, 0x6c,
	0x29, 0x79, 0xe9, 0x68, 0xb4, 0xe8, 0x68, 0xf4, 0xde, 0xd1, 0xe8, 0xe6, 0xdc, 0x58, 0xcc, 0x2a,
	0xc9, 0x15, 0x3c, 0x88, 0xbf, 0x9d, 0x8b, 0x7b, 0x23, 0xac, 0x54, 0xe2, 0xd7, 0x9e, 0xe5, 0xff,
	0x10, 0x75, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xc2, 0x51, 0xc7, 0x83, 0x01, 0x00, 0x00,
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
	{
		size, err := m.LatestHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLocalhost(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintLocalhost(dAtA []byte, offset int, v uint64) int {
	offset -= sovLocalhost(v)
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
	l = m.LatestHeight.Size()
	n += 1 + l + sovLocalhost(uint64(l))
	return n
}

func sovLocalhost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLocalhost(x uint64) (n int) {
	return sovLocalhost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocalhost
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalhost
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
				return ErrInvalidLengthLocalhost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLocalhost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LatestHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLocalhost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLocalhost
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
func skipLocalhost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLocalhost
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
					return 0, ErrIntOverflowLocalhost
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
					return 0, ErrIntOverflowLocalhost
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
				return 0, ErrInvalidLengthLocalhost
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLocalhost
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLocalhost
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLocalhost        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLocalhost          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLocalhost = fmt.Errorf("proto: unexpected end of group")
)
