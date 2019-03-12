// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/pb/keypair.proto

package corepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KeyPair struct {
	PrivateKey           []byte   `protobuf:"bytes,1,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyPair) Reset()         { *m = KeyPair{} }
func (m *KeyPair) String() string { return proto.CompactTextString(m) }
func (*KeyPair) ProtoMessage()    {}
func (*KeyPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_keypair_134ba86ecee6b014, []int{0}
}
func (m *KeyPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyPair.Unmarshal(m, b)
}
func (m *KeyPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyPair.Marshal(b, m, deterministic)
}
func (dst *KeyPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyPair.Merge(dst, src)
}
func (m *KeyPair) XXX_Size() int {
	return xxx_messageInfo_KeyPair.Size(m)
}
func (m *KeyPair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyPair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyPair proto.InternalMessageInfo

func (m *KeyPair) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *KeyPair) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyPair)(nil), "corepb.KeyPair")
}

func init() { proto.RegisterFile("core/pb/keypair.proto", fileDescriptor_keypair_134ba86ecee6b014) }

var fileDescriptor_keypair_134ba86ecee6b014 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0x2f, 0x4a,
	0xd5, 0x2f, 0x48, 0xd2, 0xcf, 0x4e, 0xad, 0x2c, 0x48, 0xcc, 0x2c, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x03, 0x09, 0x17, 0x24, 0x29, 0xb9, 0x73, 0xb1, 0x7b, 0xa7, 0x56, 0x06, 0x24,
	0x66, 0x16, 0x09, 0xc9, 0x71, 0x71, 0x15, 0x14, 0x65, 0x96, 0x25, 0x96, 0xa4, 0x7a, 0xa7, 0x56,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x21, 0x89, 0x08, 0xc9, 0x70, 0x71, 0x16, 0x94, 0x26,
	0xe5, 0x64, 0x26, 0x83, 0xa4, 0x99, 0xc0, 0xd2, 0x08, 0x81, 0x24, 0x36, 0xb0, 0xb9, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x21, 0xae, 0x2c, 0x52, 0x70, 0x00, 0x00, 0x00,
}
