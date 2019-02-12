// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pretty_block.proto

package db_inspect_pb

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

type Transaction struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Vin                  []*TXInput  `protobuf:"bytes,2,rep,name=vin,proto3" json:"vin,omitempty"`
	Vout                 []*TXOutput `protobuf:"bytes,3,rep,name=vout,proto3" json:"vout,omitempty"`
	Tip                  string      `protobuf:"bytes,4,opt,name=tip,proto3" json:"tip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d82741f9b246e68, []int{0}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Transaction) GetVin() []*TXInput {
	if m != nil {
		return m.Vin
	}
	return nil
}

func (m *Transaction) GetVout() []*TXOutput {
	if m != nil {
		return m.Vout
	}
	return nil
}

func (m *Transaction) GetTip() string {
	if m != nil {
		return m.Tip
	}
	return ""
}

type TXInput struct {
	Txid                 string   `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	Vout                 int32    `protobuf:"varint,2,opt,name=vout,proto3" json:"vout,omitempty"`
	Signature            string   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	PublicKey            string   `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TXInput) Reset()         { *m = TXInput{} }
func (m *TXInput) String() string { return proto.CompactTextString(m) }
func (*TXInput) ProtoMessage()    {}
func (*TXInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d82741f9b246e68, []int{1}
}

func (m *TXInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TXInput.Unmarshal(m, b)
}
func (m *TXInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TXInput.Marshal(b, m, deterministic)
}
func (m *TXInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TXInput.Merge(m, src)
}
func (m *TXInput) XXX_Size() int {
	return xxx_messageInfo_TXInput.Size(m)
}
func (m *TXInput) XXX_DiscardUnknown() {
	xxx_messageInfo_TXInput.DiscardUnknown(m)
}

var xxx_messageInfo_TXInput proto.InternalMessageInfo

func (m *TXInput) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *TXInput) GetVout() int32 {
	if m != nil {
		return m.Vout
	}
	return 0
}

func (m *TXInput) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *TXInput) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type TXOutput struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	PublicKeyHash        string   `protobuf:"bytes,2,opt,name=public_key_hash,json=publicKeyHash,proto3" json:"public_key_hash,omitempty"`
	Contract             string   `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TXOutput) Reset()         { *m = TXOutput{} }
func (m *TXOutput) String() string { return proto.CompactTextString(m) }
func (*TXOutput) ProtoMessage()    {}
func (*TXOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d82741f9b246e68, []int{2}
}

func (m *TXOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TXOutput.Unmarshal(m, b)
}
func (m *TXOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TXOutput.Marshal(b, m, deterministic)
}
func (m *TXOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TXOutput.Merge(m, src)
}
func (m *TXOutput) XXX_Size() int {
	return xxx_messageInfo_TXOutput.Size(m)
}
func (m *TXOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_TXOutput.DiscardUnknown(m)
}

var xxx_messageInfo_TXOutput proto.InternalMessageInfo

func (m *TXOutput) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *TXOutput) GetPublicKeyHash() string {
	if m != nil {
		return m.PublicKeyHash
	}
	return ""
}

func (m *TXOutput) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

type Block struct {
	Header               *BlockHeader   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Transactions         []*Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d82741f9b246e68, []int{3}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type BlockHeader struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	PreviousHash         string   `protobuf:"bytes,2,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
	Nonce                int64    `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature            string   `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Height               uint64   `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d82741f9b246e68, []int{4}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *BlockHeader) GetPreviousHash() string {
	if m != nil {
		return m.PreviousHash
	}
	return ""
}

func (m *BlockHeader) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *BlockHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*Transaction)(nil), "db_inspect_pb.Transaction")
	proto.RegisterType((*TXInput)(nil), "db_inspect_pb.TXInput")
	proto.RegisterType((*TXOutput)(nil), "db_inspect_pb.TXOutput")
	proto.RegisterType((*Block)(nil), "db_inspect_pb.Block")
	proto.RegisterType((*BlockHeader)(nil), "db_inspect_pb.BlockHeader")
}

func init() { proto.RegisterFile("pretty_block.proto", fileDescriptor_4d82741f9b246e68) }

var fileDescriptor_4d82741f9b246e68 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x4d, 0x6b, 0xdc, 0x30,
	0x10, 0xc5, 0x5f, 0xdb, 0x78, 0x9c, 0x6d, 0x8b, 0x08, 0xa9, 0x09, 0x2d, 0x18, 0x17, 0x8a, 0xa1,
	0xb0, 0x87, 0xed, 0xbd, 0x87, 0x9e, 0x52, 0x7a, 0x28, 0x88, 0x1c, 0x7a, 0x33, 0xb2, 0x2d, 0x62,
	0x91, 0x8d, 0x24, 0xac, 0xb1, 0xe9, 0xd2, 0x7b, 0x7f, 0x4f, 0x7f, 0x62, 0x91, 0xac, 0x5d, 0xef,
	0x6e, 0x6e, 0x33, 0xc3, 0xd3, 0x7b, 0x33, 0xef, 0x09, 0x88, 0x1e, 0x38, 0xe2, 0xbe, 0x6e, 0x76,
	0xaa, 0x7d, 0xda, 0xe8, 0x41, 0xa1, 0x22, 0xeb, 0xae, 0xa9, 0x85, 0x34, 0x9a, 0xb7, 0x58, 0xeb,
	0xa6, 0xfc, 0x1b, 0x40, 0xf6, 0x30, 0x30, 0x69, 0x58, 0x8b, 0x42, 0x49, 0xf2, 0x1a, 0x42, 0xd1,
	0xe5, 0x41, 0x11, 0x54, 0x29, 0x0d, 0x45, 0x47, 0x2a, 0x88, 0x26, 0x21, 0xf3, 0xb0, 0x88, 0xaa,
	0x6c, 0x7b, 0xbb, 0x39, 0x7b, 0xbc, 0x79, 0xf8, 0xf5, 0x5d, 0xea, 0x11, 0xa9, 0x85, 0x90, 0xcf,
	0x10, 0x4f, 0x6a, 0xc4, 0x3c, 0x72, 0xd0, 0x77, 0x2f, 0xa0, 0x3f, 0x47, 0xb4, 0x58, 0x07, 0x22,
	0x6f, 0x21, 0x42, 0xa1, 0xf3, 0xd8, 0xe9, 0xd8, 0xb2, 0x94, 0xf0, 0xca, 0xd3, 0x11, 0x02, 0x31,
	0xfe, 0x3e, 0x6e, 0xe1, 0x6a, 0x3b, 0x73, 0xec, 0x61, 0x11, 0x54, 0x89, 0x27, 0x79, 0x0f, 0xa9,
	0x11, 0x8f, 0x92, 0xe1, 0x38, 0xf0, 0x3c, 0x72, 0xe0, 0x65, 0x40, 0x3e, 0x00, 0xe8, 0xb1, 0xd9,
	0x89, 0xb6, 0x7e, 0xe2, 0x7b, 0xaf, 0x94, 0xce, 0x93, 0x1f, 0x7c, 0x5f, 0x76, 0x70, 0x75, 0xd8,
	0x89, 0xdc, 0x40, 0x32, 0xb1, 0xdd, 0xc8, 0xbd, 0xe2, 0xdc, 0x90, 0x4f, 0xf0, 0x66, 0x21, 0xa8,
	0x7b, 0x66, 0x7a, 0xa7, 0x9e, 0xd2, 0xf5, 0x91, 0xe5, 0x9e, 0x99, 0x9e, 0xdc, 0xc1, 0x55, 0xab,
	0x24, 0x0e, 0xac, 0x45, 0xbf, 0xc5, 0xb1, 0x2f, 0xff, 0x40, 0xf2, 0xcd, 0x9a, 0x4f, 0xb6, 0xb0,
	0xea, 0x39, 0xeb, 0xf8, 0xe0, 0x34, 0xb2, 0xed, 0xdd, 0x85, 0x3f, 0x0e, 0x75, 0xef, 0x10, 0xd4,
	0x23, 0xc9, 0x57, 0xb8, 0xc6, 0x25, 0x1a, 0xe3, 0x43, 0xb8, 0x7c, 0x79, 0x92, 0x1e, 0x3d, 0xc3,
	0x97, 0xff, 0x02, 0xc8, 0x4e, 0x78, 0xad, 0x87, 0xee, 0x0a, 0xef, 0xab, 0xad, 0xc9, 0x47, 0x58,
	0xeb, 0x81, 0x4f, 0x42, 0x8d, 0xe6, 0xf4, 0xc4, 0xeb, 0xc3, 0xd0, 0x5d, 0x78, 0x03, 0x89, 0x54,
	0xb2, 0x9d, 0x4d, 0x8e, 0xe8, 0xdc, 0x58, 0xfb, 0x51, 0x3c, 0x73, 0x83, 0xec, 0x79, 0x4e, 0x32,
	0xa2, 0xcb, 0xe0, 0x3c, 0x9c, 0xe4, 0x32, 0x9c, 0x5b, 0x6b, 0x87, 0x78, 0xec, 0x31, 0x5f, 0x15,
	0x41, 0x15, 0x53, 0xdf, 0x35, 0x2b, 0xf7, 0x49, 0xbf, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x96,
	0x2a, 0x20, 0xbf, 0xba, 0x02, 0x00, 0x00,
}
