// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: connection/Connection.proto

package connection

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

// State defines if a connection is in one of the following states:
// INIT, TRYOPEN, OPEN or UNINITIALIZED.
type ConnectionEnd_State int32

const (
	// Default State
	ConnectionEnd_STATE_UNINITIALIZED_UNSPECIFIED ConnectionEnd_State = 0
	// A connection end has just started the opening handshake.
	ConnectionEnd_STATE_INIT ConnectionEnd_State = 1
	// A connection end has acknowledged the handshake step on the counterparty
	// chain.
	ConnectionEnd_STATE_TRYOPEN ConnectionEnd_State = 2
	// A connection end has completed the handshake.
	ConnectionEnd_STATE_OPEN ConnectionEnd_State = 3
)

var ConnectionEnd_State_name = map[int32]string{
	0: "STATE_UNINITIALIZED_UNSPECIFIED",
	1: "STATE_INIT",
	2: "STATE_TRYOPEN",
	3: "STATE_OPEN",
}

var ConnectionEnd_State_value = map[string]int32{
	"STATE_UNINITIALIZED_UNSPECIFIED": 0,
	"STATE_INIT":                      1,
	"STATE_TRYOPEN":                   2,
	"STATE_OPEN":                      3,
}

func (x ConnectionEnd_State) String() string {
	return proto.EnumName(ConnectionEnd_State_name, int32(x))
}

func (ConnectionEnd_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e47af75a3093bd8, []int{0, 0}
}

// ConnectionEnd defines a stateful object on a chain connected to another
// separate one.
// NOTE: there must only be 2 defined ConnectionEnds to establish
// a connection between two chains.
type ConnectionEnd struct {
	// client associated with this connection.
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// IBC version which can be utilised to determine encodings or protocols for
	// channels or packets utilising this connection.
	Versions []*Version `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
	// current state of the connection end.
	State ConnectionEnd_State `protobuf:"varint,3,opt,name=state,proto3,enum=ConnectionEnd_State" json:"state,omitempty"`
	// counterparty chain associated with this connection.
	Counterparty *Counterparty `protobuf:"bytes,4,opt,name=counterparty,proto3" json:"counterparty,omitempty"`
	// delay period that must pass before a consensus state can be used for packet-verification
	// NOTE: delay period logic is only implemented by some clients.
	DelayPeriod uint64 `protobuf:"varint,5,opt,name=delay_period,json=delayPeriod,proto3" json:"delay_period,omitempty"`
}

func (m *ConnectionEnd) Reset()         { *m = ConnectionEnd{} }
func (m *ConnectionEnd) String() string { return proto.CompactTextString(m) }
func (*ConnectionEnd) ProtoMessage()    {}
func (*ConnectionEnd) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e47af75a3093bd8, []int{0}
}
func (m *ConnectionEnd) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConnectionEnd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConnectionEnd.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConnectionEnd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionEnd.Merge(m, src)
}
func (m *ConnectionEnd) XXX_Size() int {
	return m.Size()
}
func (m *ConnectionEnd) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionEnd.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionEnd proto.InternalMessageInfo

func (m *ConnectionEnd) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ConnectionEnd) GetVersions() []*Version {
	if m != nil {
		return m.Versions
	}
	return nil
}

func (m *ConnectionEnd) GetState() ConnectionEnd_State {
	if m != nil {
		return m.State
	}
	return ConnectionEnd_STATE_UNINITIALIZED_UNSPECIFIED
}

func (m *ConnectionEnd) GetCounterparty() *Counterparty {
	if m != nil {
		return m.Counterparty
	}
	return nil
}

func (m *ConnectionEnd) GetDelayPeriod() uint64 {
	if m != nil {
		return m.DelayPeriod
	}
	return 0
}

// Counterparty defines the counterparty chain associated with a connection end.
type Counterparty struct {
	// identifies the client on the counterparty chain associated with a given
	// connection.
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// identifies the connection end on the counterparty chain associated with a
	// given connection.
	ConnectionId string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	// commitment merkle prefix of the counterparty chain.
	Prefix *MerklePrefix `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (m *Counterparty) Reset()         { *m = Counterparty{} }
func (m *Counterparty) String() string { return proto.CompactTextString(m) }
func (*Counterparty) ProtoMessage()    {}
func (*Counterparty) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e47af75a3093bd8, []int{1}
}
func (m *Counterparty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Counterparty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Counterparty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Counterparty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Counterparty.Merge(m, src)
}
func (m *Counterparty) XXX_Size() int {
	return m.Size()
}
func (m *Counterparty) XXX_DiscardUnknown() {
	xxx_messageInfo_Counterparty.DiscardUnknown(m)
}

var xxx_messageInfo_Counterparty proto.InternalMessageInfo

func (m *Counterparty) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Counterparty) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *Counterparty) GetPrefix() *MerklePrefix {
	if m != nil {
		return m.Prefix
	}
	return nil
}

type MerklePrefix struct {
	KeyPrefix []byte `protobuf:"bytes,1,opt,name=key_prefix,json=keyPrefix,proto3" json:"key_prefix,omitempty"`
}

func (m *MerklePrefix) Reset()         { *m = MerklePrefix{} }
func (m *MerklePrefix) String() string { return proto.CompactTextString(m) }
func (*MerklePrefix) ProtoMessage()    {}
func (*MerklePrefix) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e47af75a3093bd8, []int{2}
}
func (m *MerklePrefix) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MerklePrefix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MerklePrefix.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MerklePrefix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerklePrefix.Merge(m, src)
}
func (m *MerklePrefix) XXX_Size() int {
	return m.Size()
}
func (m *MerklePrefix) XXX_DiscardUnknown() {
	xxx_messageInfo_MerklePrefix.DiscardUnknown(m)
}

var xxx_messageInfo_MerklePrefix proto.InternalMessageInfo

func (m *MerklePrefix) GetKeyPrefix() []byte {
	if m != nil {
		return m.KeyPrefix
	}
	return nil
}

// Version defines the versioning scheme used to negotiate the IBC verison in
// the connection handshake.
type Version struct {
	// unique version identifier
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// list of features compatible with the specified identifier
	Features []string `protobuf:"bytes,2,rep,name=features,proto3" json:"features,omitempty"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e47af75a3093bd8, []int{3}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Version.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return m.Size()
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Version) GetFeatures() []string {
	if m != nil {
		return m.Features
	}
	return nil
}

func init() {
	proto.RegisterEnum("ConnectionEnd_State", ConnectionEnd_State_name, ConnectionEnd_State_value)
	proto.RegisterType((*ConnectionEnd)(nil), "ConnectionEnd")
	proto.RegisterType((*Counterparty)(nil), "Counterparty")
	proto.RegisterType((*MerklePrefix)(nil), "MerklePrefix")
	proto.RegisterType((*Version)(nil), "Version")
}

func init() { proto.RegisterFile("connection/Connection.proto", fileDescriptor_0e47af75a3093bd8) }

var fileDescriptor_0e47af75a3093bd8 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x8e, 0xd2, 0x40,
	0x14, 0x86, 0x19, 0x58, 0x56, 0x38, 0x94, 0x0d, 0x4e, 0xbc, 0x68, 0xdc, 0x58, 0x6b, 0x57, 0x93,
	0xc6, 0x68, 0x89, 0xf8, 0x04, 0x2b, 0x5b, 0x93, 0x49, 0x14, 0xc9, 0xc0, 0x9a, 0xb8, 0x37, 0x0d,
	0xdb, 0x1e, 0xcc, 0x04, 0x32, 0x6d, 0x86, 0x59, 0x95, 0xb7, 0xf0, 0x65, 0x7c, 0x07, 0x2f, 0xf7,
	0xd2, 0x4b, 0x03, 0x2f, 0x62, 0x3a, 0x6d, 0x68, 0xb9, 0xf1, 0xae, 0xe7, 0xfb, 0xbf, 0xe6, 0x9c,
	0x9e, 0x53, 0x38, 0x8f, 0x53, 0x29, 0x31, 0xd6, 0x22, 0x95, 0xc3, 0xf1, 0xe1, 0x31, 0xc8, 0x54,
	0xaa, 0x53, 0xef, 0x57, 0x13, 0xfa, 0x15, 0x0c, 0x65, 0x42, 0xcf, 0xa1, 0x1b, 0xaf, 0x05, 0x4a,
	0x1d, 0x89, 0xc4, 0x26, 0x2e, 0xf1, 0xbb, 0xbc, 0x53, 0x00, 0x96, 0xd0, 0xe7, 0xd0, 0xf9, 0x86,
	0x6a, 0x23, 0x52, 0xb9, 0xb1, 0x9b, 0x6e, 0xcb, 0xef, 0x8d, 0x3a, 0xc1, 0xe7, 0x02, 0xf0, 0x43,
	0x42, 0x5f, 0x42, 0x7b, 0xa3, 0x17, 0x1a, 0xed, 0x96, 0x4b, 0xfc, 0xb3, 0xd1, 0xa3, 0xe0, 0xa8,
	0x43, 0x30, 0xcb, 0x33, 0x5e, 0x28, 0xf4, 0x0d, 0x58, 0x71, 0x7a, 0x27, 0x35, 0xaa, 0x6c, 0xa1,
	0xf4, 0xd6, 0x3e, 0x71, 0x89, 0xdf, 0x1b, 0xf5, 0x83, 0x71, 0x0d, 0xf2, 0x23, 0x85, 0x3e, 0x03,
	0x2b, 0xc1, 0xf5, 0x62, 0x1b, 0x65, 0xa8, 0x44, 0x9a, 0xd8, 0x6d, 0x97, 0xf8, 0x27, 0xbc, 0x67,
	0xd8, 0xd4, 0x20, 0x2f, 0x82, 0xb6, 0xe9, 0x42, 0x2f, 0xe0, 0xe9, 0x6c, 0x7e, 0x39, 0x0f, 0xa3,
	0xeb, 0x09, 0x9b, 0xb0, 0x39, 0xbb, 0xfc, 0xc0, 0x6e, 0xc2, 0xab, 0xe8, 0x7a, 0x32, 0x9b, 0x86,
	0x63, 0xf6, 0x9e, 0x85, 0x57, 0x83, 0x06, 0x3d, 0x03, 0x28, 0xa4, 0x5c, 0x19, 0x10, 0xfa, 0x10,
	0xfa, 0x45, 0x3d, 0xe7, 0x5f, 0x3e, 0x4d, 0xc3, 0xc9, 0xa0, 0x59, 0x29, 0xa6, 0x6e, 0x79, 0xdf,
	0xc1, 0xaa, 0x4f, 0xf8, 0xff, 0xad, 0x5d, 0x40, 0xbf, 0xba, 0x41, 0x2e, 0x34, 0x8d, 0x60, 0x55,
	0x90, 0x25, 0xf4, 0x05, 0x9c, 0x66, 0x0a, 0x97, 0xe2, 0x87, 0xd9, 0x5a, 0xbe, 0x82, 0x8f, 0xa8,
	0x56, 0x6b, 0x9c, 0x1a, 0xc8, 0xcb, 0xd0, 0x7b, 0x0d, 0x56, 0x9d, 0xd3, 0x27, 0x00, 0x2b, 0xdc,
	0x46, 0xe5, 0xab, 0x79, 0x67, 0x8b, 0x77, 0x57, 0xb8, 0x2d, 0x62, 0x2f, 0x84, 0x07, 0xe5, 0x7d,
	0xa8, 0x03, 0x20, 0x12, 0x94, 0x5a, 0x2c, 0x05, 0xaa, 0x72, 0xc6, 0x1a, 0xa1, 0x8f, 0xa1, 0xb3,
	0xc4, 0x85, 0xbe, 0x53, 0x58, 0xdc, 0xb6, 0xcb, 0x0f, 0xf5, 0xbb, 0x57, 0xbf, 0x77, 0x0e, 0xb9,
	0xdf, 0x39, 0xe4, 0xef, 0xce, 0x21, 0x3f, 0xf7, 0x4e, 0xe3, 0x7e, 0xef, 0x34, 0xfe, 0xec, 0x9d,
	0xc6, 0x0d, 0xcd, 0x56, 0x5f, 0x87, 0xe2, 0x36, 0x1e, 0x56, 0x1f, 0x73, 0x7b, 0x6a, 0xfe, 0xad,
	0xb7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xa7, 0x83, 0x81, 0x7a, 0x02, 0x00, 0x00,
}

func (m *ConnectionEnd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectionEnd) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConnectionEnd) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DelayPeriod != 0 {
		i = encodeVarintConnection(dAtA, i, uint64(m.DelayPeriod))
		i--
		dAtA[i] = 0x28
	}
	if m.Counterparty != nil {
		{
			size, err := m.Counterparty.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConnection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.State != 0 {
		i = encodeVarintConnection(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Versions) > 0 {
		for iNdEx := len(m.Versions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Versions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConnection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintConnection(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Counterparty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Counterparty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Counterparty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Prefix != nil {
		{
			size, err := m.Prefix.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConnection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintConnection(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintConnection(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MerklePrefix) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MerklePrefix) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MerklePrefix) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.KeyPrefix) > 0 {
		i -= len(m.KeyPrefix)
		copy(dAtA[i:], m.KeyPrefix)
		i = encodeVarintConnection(dAtA, i, uint64(len(m.KeyPrefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Version) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Version) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Version) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Features) > 0 {
		for iNdEx := len(m.Features) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Features[iNdEx])
			copy(dAtA[i:], m.Features[iNdEx])
			i = encodeVarintConnection(dAtA, i, uint64(len(m.Features[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = encodeVarintConnection(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConnection(dAtA []byte, offset int, v uint64) int {
	offset -= sovConnection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConnectionEnd) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovConnection(uint64(l))
	}
	if len(m.Versions) > 0 {
		for _, e := range m.Versions {
			l = e.Size()
			n += 1 + l + sovConnection(uint64(l))
		}
	}
	if m.State != 0 {
		n += 1 + sovConnection(uint64(m.State))
	}
	if m.Counterparty != nil {
		l = m.Counterparty.Size()
		n += 1 + l + sovConnection(uint64(l))
	}
	if m.DelayPeriod != 0 {
		n += 1 + sovConnection(uint64(m.DelayPeriod))
	}
	return n
}

func (m *Counterparty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovConnection(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovConnection(uint64(l))
	}
	if m.Prefix != nil {
		l = m.Prefix.Size()
		n += 1 + l + sovConnection(uint64(l))
	}
	return n
}

func (m *MerklePrefix) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.KeyPrefix)
	if l > 0 {
		n += 1 + l + sovConnection(uint64(l))
	}
	return n
}

func (m *Version) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovConnection(uint64(l))
	}
	if len(m.Features) > 0 {
		for _, s := range m.Features {
			l = len(s)
			n += 1 + l + sovConnection(uint64(l))
		}
	}
	return n
}

func sovConnection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConnection(x uint64) (n int) {
	return sovConnection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConnectionEnd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnection
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
			return fmt.Errorf("proto: ConnectionEnd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectionEnd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
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
				return ErrInvalidLengthConnection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConnection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Versions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
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
				return ErrInvalidLengthConnection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConnection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Versions = append(m.Versions, &Version{})
			if err := m.Versions[len(m.Versions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= ConnectionEnd_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counterparty", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
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
				return ErrInvalidLengthConnection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConnection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Counterparty == nil {
				m.Counterparty = &Counterparty{}
			}
			if err := m.Counterparty.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelayPeriod", wireType)
			}
			m.DelayPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelayPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConnection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConnection
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
func (m *Counterparty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnection
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
			return fmt.Errorf("proto: Counterparty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Counterparty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
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
				return ErrInvalidLengthConnection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConnection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
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
				return ErrInvalidLengthConnection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConnection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
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
				return ErrInvalidLengthConnection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConnection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Prefix == nil {
				m.Prefix = &MerklePrefix{}
			}
			if err := m.Prefix.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConnection
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
func (m *MerklePrefix) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnection
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
			return fmt.Errorf("proto: MerklePrefix: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MerklePrefix: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyPrefix", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
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
				return ErrInvalidLengthConnection
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConnection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyPrefix = append(m.KeyPrefix[:0], dAtA[iNdEx:postIndex]...)
			if m.KeyPrefix == nil {
				m.KeyPrefix = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConnection
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
func (m *Version) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnection
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
			return fmt.Errorf("proto: Version: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Version: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
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
				return ErrInvalidLengthConnection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConnection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Features", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
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
				return ErrInvalidLengthConnection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConnection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Features = append(m.Features, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConnection
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
func skipConnection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConnection
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
					return 0, ErrIntOverflowConnection
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
					return 0, ErrIntOverflowConnection
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
				return 0, ErrInvalidLengthConnection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConnection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConnection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConnection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConnection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConnection = fmt.Errorf("proto: unexpected end of group")
)
