// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/dappley/go-dappley/rpc/pb/rpc.proto

package rpcpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pb "github.com/dappley/go-dappley/network/pb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// The request message
type CreateWalletRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateWalletRequest) Reset()         { *m = CreateWalletRequest{} }
func (m *CreateWalletRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWalletRequest) ProtoMessage()    {}
func (*CreateWalletRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_d2dd3042bca6be8a, []int{0}
}
func (m *CreateWalletRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWalletRequest.Unmarshal(m, b)
}
func (m *CreateWalletRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWalletRequest.Marshal(b, m, deterministic)
}
func (dst *CreateWalletRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWalletRequest.Merge(dst, src)
}
func (m *CreateWalletRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWalletRequest.Size(m)
}
func (m *CreateWalletRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWalletRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWalletRequest proto.InternalMessageInfo

func (m *CreateWalletRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetBalanceRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBalanceRequest) Reset()         { *m = GetBalanceRequest{} }
func (m *GetBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*GetBalanceRequest) ProtoMessage()    {}
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_d2dd3042bca6be8a, []int{1}
}
func (m *GetBalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBalanceRequest.Unmarshal(m, b)
}
func (m *GetBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBalanceRequest.Marshal(b, m, deterministic)
}
func (dst *GetBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBalanceRequest.Merge(dst, src)
}
func (m *GetBalanceRequest) XXX_Size() int {
	return xxx_messageInfo_GetBalanceRequest.Size(m)
}
func (m *GetBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBalanceRequest proto.InternalMessageInfo

func (m *GetBalanceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetBalanceRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type SendRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Ammount              int64    `protobuf:"varint,4,opt,name=ammount,proto3" json:"ammount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_d2dd3042bca6be8a, []int{2}
}
func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (dst *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(dst, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SendRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *SendRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SendRequest) GetAmmount() int64 {
	if m != nil {
		return m.Ammount
	}
	return 0
}

type GetPeerInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPeerInfoRequest) Reset()         { *m = GetPeerInfoRequest{} }
func (m *GetPeerInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetPeerInfoRequest) ProtoMessage()    {}
func (*GetPeerInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_d2dd3042bca6be8a, []int{3}
}
func (m *GetPeerInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPeerInfoRequest.Unmarshal(m, b)
}
func (m *GetPeerInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPeerInfoRequest.Marshal(b, m, deterministic)
}
func (dst *GetPeerInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPeerInfoRequest.Merge(dst, src)
}
func (m *GetPeerInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetPeerInfoRequest.Size(m)
}
func (m *GetPeerInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPeerInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPeerInfoRequest proto.InternalMessageInfo

type CreateWalletResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateWalletResponse) Reset()         { *m = CreateWalletResponse{} }
func (m *CreateWalletResponse) String() string { return proto.CompactTextString(m) }
func (*CreateWalletResponse) ProtoMessage()    {}
func (*CreateWalletResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_d2dd3042bca6be8a, []int{4}
}
func (m *CreateWalletResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWalletResponse.Unmarshal(m, b)
}
func (m *CreateWalletResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWalletResponse.Marshal(b, m, deterministic)
}
func (dst *CreateWalletResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWalletResponse.Merge(dst, src)
}
func (m *CreateWalletResponse) XXX_Size() int {
	return xxx_messageInfo_CreateWalletResponse.Size(m)
}
func (m *CreateWalletResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWalletResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWalletResponse proto.InternalMessageInfo

func (m *CreateWalletResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateWalletResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type GetBalanceResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Ammount              int64    `protobuf:"varint,2,opt,name=ammount,proto3" json:"ammount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBalanceResponse) Reset()         { *m = GetBalanceResponse{} }
func (m *GetBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*GetBalanceResponse) ProtoMessage()    {}
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_d2dd3042bca6be8a, []int{5}
}
func (m *GetBalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBalanceResponse.Unmarshal(m, b)
}
func (m *GetBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBalanceResponse.Marshal(b, m, deterministic)
}
func (dst *GetBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBalanceResponse.Merge(dst, src)
}
func (m *GetBalanceResponse) XXX_Size() int {
	return xxx_messageInfo_GetBalanceResponse.Size(m)
}
func (m *GetBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBalanceResponse proto.InternalMessageInfo

func (m *GetBalanceResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetBalanceResponse) GetAmmount() int64 {
	if m != nil {
		return m.Ammount
	}
	return 0
}

type SendResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_d2dd3042bca6be8a, []int{6}
}
func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (dst *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(dst, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func (m *SendResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetPeerInfoResponse struct {
	PeerList             *pb.Peerlist `protobuf:"bytes,1,opt,name=peerList,proto3" json:"peerList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetPeerInfoResponse) Reset()         { *m = GetPeerInfoResponse{} }
func (m *GetPeerInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetPeerInfoResponse) ProtoMessage()    {}
func (*GetPeerInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_d2dd3042bca6be8a, []int{7}
}
func (m *GetPeerInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPeerInfoResponse.Unmarshal(m, b)
}
func (m *GetPeerInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPeerInfoResponse.Marshal(b, m, deterministic)
}
func (dst *GetPeerInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPeerInfoResponse.Merge(dst, src)
}
func (m *GetPeerInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetPeerInfoResponse.Size(m)
}
func (m *GetPeerInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPeerInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPeerInfoResponse proto.InternalMessageInfo

func (m *GetPeerInfoResponse) GetPeerList() *pb.Peerlist {
	if m != nil {
		return m.PeerList
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateWalletRequest)(nil), "rpcpb.CreateWalletRequest")
	proto.RegisterType((*GetBalanceRequest)(nil), "rpcpb.GetBalanceRequest")
	proto.RegisterType((*SendRequest)(nil), "rpcpb.SendRequest")
	proto.RegisterType((*GetPeerInfoRequest)(nil), "rpcpb.GetPeerInfoRequest")
	proto.RegisterType((*CreateWalletResponse)(nil), "rpcpb.CreateWalletResponse")
	proto.RegisterType((*GetBalanceResponse)(nil), "rpcpb.GetBalanceResponse")
	proto.RegisterType((*SendResponse)(nil), "rpcpb.SendResponse")
	proto.RegisterType((*GetPeerInfoResponse)(nil), "rpcpb.GetPeerInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConnectClient is the client API for Connect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectClient interface {
	// Sends a greeting
	RpcCreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error)
	RpcGetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	RpcSend(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	RpcGetPeerInfo(ctx context.Context, in *GetPeerInfoRequest, opts ...grpc.CallOption) (*GetPeerInfoResponse, error)
}

type connectClient struct {
	cc *grpc.ClientConn
}

func NewConnectClient(cc *grpc.ClientConn) ConnectClient {
	return &connectClient{cc}
}

func (c *connectClient) RpcCreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error) {
	out := new(CreateWalletResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.Connect/RpcCreateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) RpcGetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.Connect/RpcGetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) RpcSend(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.Connect/RpcSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) RpcGetPeerInfo(ctx context.Context, in *GetPeerInfoRequest, opts ...grpc.CallOption) (*GetPeerInfoResponse, error) {
	out := new(GetPeerInfoResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.Connect/RpcGetPeerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectServer is the server API for Connect service.
type ConnectServer interface {
	// Sends a greeting
	RpcCreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error)
	RpcGetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	RpcSend(context.Context, *SendRequest) (*SendResponse, error)
	RpcGetPeerInfo(context.Context, *GetPeerInfoRequest) (*GetPeerInfoResponse, error)
}

func RegisterConnectServer(s *grpc.Server, srv ConnectServer) {
	s.RegisterService(&_Connect_serviceDesc, srv)
}

func _Connect_RpcCreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).RpcCreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Connect/RpcCreateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).RpcCreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_RpcGetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).RpcGetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Connect/RpcGetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).RpcGetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_RpcSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).RpcSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Connect/RpcSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).RpcSend(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_RpcGetPeerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).RpcGetPeerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Connect/RpcGetPeerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).RpcGetPeerInfo(ctx, req.(*GetPeerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Connect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.Connect",
	HandlerType: (*ConnectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RpcCreateWallet",
			Handler:    _Connect_RpcCreateWallet_Handler,
		},
		{
			MethodName: "RpcGetBalance",
			Handler:    _Connect_RpcGetBalance_Handler,
		},
		{
			MethodName: "RpcSend",
			Handler:    _Connect_RpcSend_Handler,
		},
		{
			MethodName: "RpcGetPeerInfo",
			Handler:    _Connect_RpcGetPeerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/dappley/go-dappley/rpc/pb/rpc.proto",
}

func init() {
	proto.RegisterFile("github.com/dappley/go-dappley/rpc/pb/rpc.proto", fileDescriptor_rpc_d2dd3042bca6be8a)
}

var fileDescriptor_rpc_d2dd3042bca6be8a = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4b, 0x8f, 0xd3, 0x30,
	0x10, 0xa6, 0x69, 0xd9, 0xc2, 0x2c, 0x2c, 0xc2, 0xd9, 0x43, 0x08, 0x97, 0x55, 0x4e, 0xe5, 0x40,
	0x22, 0x2d, 0x48, 0x9c, 0x69, 0xa5, 0x3e, 0x50, 0x0f, 0x55, 0x38, 0xf4, 0x88, 0x1c, 0x67, 0x5a,
	0x2a, 0x12, 0xdb, 0xd8, 0xae, 0x10, 0xff, 0x06, 0xf1, 0x4b, 0x51, 0x1c, 0x27, 0x34, 0x10, 0x05,
	0x4e, 0xf1, 0x3c, 0xbe, 0x79, 0x7c, 0xf3, 0x05, 0xe2, 0xe3, 0xc9, 0x7c, 0x3e, 0x67, 0x31, 0x13,
	0x65, 0x92, 0x53, 0x29, 0x0b, 0xfc, 0x9e, 0x1c, 0xc5, 0xeb, 0xe6, 0xa9, 0x24, 0x4b, 0x64, 0x56,
	0x7d, 0x62, 0xa9, 0x84, 0x11, 0xe4, 0xa1, 0x92, 0x4c, 0x66, 0xe1, 0xbb, 0x61, 0x18, 0x47, 0xf3,
	0x4d, 0xa8, 0x2f, 0x15, 0x54, 0x22, 0xaa, 0xe2, 0xa4, 0x4d, 0x8d, 0x8f, 0x5e, 0x81, 0xbf, 0x50,
	0x48, 0x0d, 0xee, 0x69, 0x51, 0xa0, 0x49, 0xf1, 0xeb, 0x19, 0xb5, 0x21, 0x04, 0x26, 0x9c, 0x96,
	0x18, 0x8c, 0xee, 0x46, 0xb3, 0xc7, 0xa9, 0x7d, 0x47, 0xef, 0xe1, 0xf9, 0x0a, 0xcd, 0x9c, 0x16,
	0x94, 0x33, 0x1c, 0x48, 0x24, 0x01, 0x4c, 0x69, 0x9e, 0x2b, 0xd4, 0x3a, 0xf0, 0xac, 0xbb, 0x31,
	0xa3, 0x4f, 0x70, 0xfd, 0x11, 0x79, 0x3e, 0x04, 0x26, 0x30, 0x39, 0x28, 0x51, 0x3a, 0xa4, 0x7d,
	0x93, 0x1b, 0xf0, 0x8c, 0x08, 0xc6, 0xd6, 0xe3, 0x19, 0x61, 0x1b, 0x94, 0xa5, 0x38, 0x73, 0x13,
	0x4c, 0xee, 0x46, 0xb3, 0x71, 0xda, 0x98, 0xd1, 0x2d, 0x90, 0x15, 0x9a, 0x1d, 0xa2, 0xda, 0xf0,
	0x83, 0x70, 0x7d, 0xa2, 0x0f, 0x70, 0xdb, 0x5d, 0x52, 0x4b, 0xc1, 0xb5, 0x1d, 0xb4, 0x44, 0xad,
	0xe9, 0xb1, 0x19, 0xa1, 0x31, 0x07, 0x56, 0x58, 0xdb, 0x0e, 0x2d, 0x0b, 0xff, 0x55, 0xc9, 0xcd,
	0xea, 0x75, 0x67, 0x9d, 0xc1, 0x93, 0x9a, 0x8c, 0x7f, 0xd5, 0x88, 0x96, 0xe0, 0x77, 0xb6, 0x72,
	0x80, 0x04, 0x1e, 0x55, 0xd7, 0xdc, 0x9e, 0xb4, 0xb1, 0x88, 0xeb, 0x7b, 0x3f, 0x76, 0x97, 0x96,
	0x59, 0xbc, 0x73, 0x87, 0x4e, 0xdb, 0xa4, 0xfb, 0x1f, 0x1e, 0x4c, 0x17, 0x82, 0x73, 0x64, 0x86,
	0x6c, 0xe1, 0x59, 0x2a, 0xd9, 0x25, 0x2d, 0x24, 0x8c, 0xad, 0x98, 0xe2, 0x1e, 0x41, 0x84, 0x2f,
	0x7b, 0x63, 0xf5, 0x20, 0xd1, 0x03, 0xb2, 0x84, 0xa7, 0xa9, 0x64, 0xbf, 0x89, 0x21, 0x81, 0xcb,
	0xff, 0x4b, 0x31, 0xe1, 0x8b, 0x9e, 0x48, 0x5b, 0xe7, 0x2d, 0x4c, 0x53, 0xc9, 0x2a, 0x5a, 0x08,
	0x71, 0x79, 0x17, 0x82, 0x09, 0xfd, 0x8e, 0xaf, 0x45, 0x6d, 0xe0, 0xa6, 0xee, 0xde, 0x50, 0x44,
	0x2e, 0x9a, 0xfc, 0x21, 0x86, 0x30, 0xec, 0x0b, 0x35, 0xa5, 0xe6, 0x57, 0x3f, 0xbd, 0xf1, 0x7a,
	0xbb, 0xcf, 0xae, 0xec, 0xef, 0xf1, 0xe6, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x0a, 0x51,
	0xaa, 0x90, 0x03, 0x00, 0x00,
}
