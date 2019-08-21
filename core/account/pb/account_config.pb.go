// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/dappley/go-dappley/core/account/pb/account_config.proto

package accountpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AccountConfig struct {
	FilePath             string   `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountConfig) Reset()         { *m = AccountConfig{} }
func (m *AccountConfig) String() string { return proto.CompactTextString(m) }
func (*AccountConfig) ProtoMessage()    {}
func (*AccountConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cd853538fa438eb, []int{0}
}

func (m *AccountConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountConfig.Unmarshal(m, b)
}
func (m *AccountConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountConfig.Marshal(b, m, deterministic)
}
func (m *AccountConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountConfig.Merge(m, src)
}
func (m *AccountConfig) XXX_Size() int {
	return xxx_messageInfo_AccountConfig.Size(m)
}
func (m *AccountConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AccountConfig proto.InternalMessageInfo

func (m *AccountConfig) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func init() {
	proto.RegisterType((*AccountConfig)(nil), "accountpb.AccountConfig")
}

func init() {
	proto.RegisterFile("github.com/dappley/go-dappley/core/account/pb/account_config.proto", fileDescriptor_2cd853538fa438eb)
}

var fileDescriptor_2cd853538fa438eb = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x49, 0x2c, 0x28, 0xc8, 0x49, 0xad, 0xd4, 0x4f,
	0xcf, 0xd7, 0x85, 0x31, 0x93, 0xf3, 0x8b, 0x52, 0xf5, 0x93, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0xf4,
	0x0b, 0x92, 0xf4, 0x13, 0x93, 0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0xe2, 0x93, 0xf3, 0xf3, 0xd2, 0x32,
	0xd3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0xa1, 0xa2, 0x05, 0x49, 0x4a, 0x3a, 0x5c,
	0xbc, 0x8e, 0x10, 0x8e, 0x33, 0x58, 0x85, 0x90, 0x34, 0x17, 0x67, 0x5a, 0x66, 0x4e, 0x6a, 0x7c,
	0x41, 0x62, 0x49, 0x86, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x07, 0x48, 0x20, 0x20, 0xb1,
	0x24, 0x23, 0x89, 0x0d, 0xac, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x8b, 0x14, 0xb7,
	0x84, 0x00, 0x00, 0x00,
}
