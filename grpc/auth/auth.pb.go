// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/auth.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	auth/auth.proto

It has these top-level messages:
	AuthArgs
	AuthRes
	TokenInfoArgs
	TokenInfoRes
	RefreshTokenArgs
	RefreshTokenRes
	DestroyTokenArgs
	DestroyTokenRes
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthArgs struct {
	UserId     int32  `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	AppFrom    string `protobuf:"bytes,2,opt,name=AppFrom" json:"AppFrom,omitempty"`
	AppChannel string `protobuf:"bytes,3,opt,name=AppChannel" json:"AppChannel,omitempty"`
	AppVer     int32  `protobuf:"varint,4,opt,name=AppVer" json:"AppVer,omitempty"`
	UniqueId   string `protobuf:"bytes,5,opt,name=UniqueId" json:"UniqueId,omitempty"`
	Token      string `protobuf:"bytes,6,opt,name=Token" json:"Token,omitempty"`
}

func (m *AuthArgs) Reset()                    { *m = AuthArgs{} }
func (m *AuthArgs) String() string            { return proto1.CompactTextString(m) }
func (*AuthArgs) ProtoMessage()               {}
func (*AuthArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthArgs) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AuthArgs) GetAppFrom() string {
	if m != nil {
		return m.AppFrom
	}
	return ""
}

func (m *AuthArgs) GetAppChannel() string {
	if m != nil {
		return m.AppChannel
	}
	return ""
}

func (m *AuthArgs) GetAppVer() int32 {
	if m != nil {
		return m.AppVer
	}
	return 0
}

func (m *AuthArgs) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *AuthArgs) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthRes struct {
	Ret int32  `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=Msg" json:"Msg,omitempty"`
}

func (m *AuthRes) Reset()                    { *m = AuthRes{} }
func (m *AuthRes) String() string            { return proto1.CompactTextString(m) }
func (*AuthRes) ProtoMessage()               {}
func (*AuthRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthRes) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *AuthRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type TokenInfoArgs struct {
	UserId     int32  `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	AppFrom    string `protobuf:"bytes,2,opt,name=AppFrom" json:"AppFrom,omitempty"`
	AppChannel string `protobuf:"bytes,3,opt,name=AppChannel" json:"AppChannel,omitempty"`
	AppVer     int32  `protobuf:"varint,4,opt,name=AppVer" json:"AppVer,omitempty"`
	UniqueId   string `protobuf:"bytes,5,opt,name=UniqueId" json:"UniqueId,omitempty"`
	Token      string `protobuf:"bytes,6,opt,name=Token" json:"Token,omitempty"`
}

func (m *TokenInfoArgs) Reset()                    { *m = TokenInfoArgs{} }
func (m *TokenInfoArgs) String() string            { return proto1.CompactTextString(m) }
func (*TokenInfoArgs) ProtoMessage()               {}
func (*TokenInfoArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TokenInfoArgs) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *TokenInfoArgs) GetAppFrom() string {
	if m != nil {
		return m.AppFrom
	}
	return ""
}

func (m *TokenInfoArgs) GetAppChannel() string {
	if m != nil {
		return m.AppChannel
	}
	return ""
}

func (m *TokenInfoArgs) GetAppVer() int32 {
	if m != nil {
		return m.AppVer
	}
	return 0
}

func (m *TokenInfoArgs) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *TokenInfoArgs) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type TokenInfoRes struct {
	Expire int64 `protobuf:"varint,1,opt,name=Expire" json:"Expire,omitempty"`
}

func (m *TokenInfoRes) Reset()                    { *m = TokenInfoRes{} }
func (m *TokenInfoRes) String() string            { return proto1.CompactTextString(m) }
func (*TokenInfoRes) ProtoMessage()               {}
func (*TokenInfoRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TokenInfoRes) GetExpire() int64 {
	if m != nil {
		return m.Expire
	}
	return 0
}

type RefreshTokenArgs struct {
	UserId       int32  `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	AppFrom      string `protobuf:"bytes,2,opt,name=AppFrom" json:"AppFrom,omitempty"`
	AppChannel   string `protobuf:"bytes,3,opt,name=AppChannel" json:"AppChannel,omitempty"`
	AppVer       int32  `protobuf:"varint,4,opt,name=AppVer" json:"AppVer,omitempty"`
	UniqueId     string `protobuf:"bytes,5,opt,name=UniqueId" json:"UniqueId,omitempty"`
	Token        string `protobuf:"bytes,6,opt,name=Token" json:"Token,omitempty"`
	RefreshToken string `protobuf:"bytes,7,opt,name=RefreshToken" json:"RefreshToken,omitempty"`
}

func (m *RefreshTokenArgs) Reset()                    { *m = RefreshTokenArgs{} }
func (m *RefreshTokenArgs) String() string            { return proto1.CompactTextString(m) }
func (*RefreshTokenArgs) ProtoMessage()               {}
func (*RefreshTokenArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RefreshTokenArgs) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RefreshTokenArgs) GetAppFrom() string {
	if m != nil {
		return m.AppFrom
	}
	return ""
}

func (m *RefreshTokenArgs) GetAppChannel() string {
	if m != nil {
		return m.AppChannel
	}
	return ""
}

func (m *RefreshTokenArgs) GetAppVer() int32 {
	if m != nil {
		return m.AppVer
	}
	return 0
}

func (m *RefreshTokenArgs) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *RefreshTokenArgs) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RefreshTokenArgs) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type RefreshTokenRes struct {
	Ret          int32  `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
	Msg          string `protobuf:"bytes,2,opt,name=Msg" json:"Msg,omitempty"`
	UserId       int32  `protobuf:"varint,3,opt,name=UserId" json:"UserId,omitempty"`
	Token        string `protobuf:"bytes,4,opt,name=Token" json:"Token,omitempty"`
	RefreshToken string `protobuf:"bytes,5,opt,name=RefreshToken" json:"RefreshToken,omitempty"`
	Expire       int64  `protobuf:"varint,6,opt,name=Expire" json:"Expire,omitempty"`
}

func (m *RefreshTokenRes) Reset()                    { *m = RefreshTokenRes{} }
func (m *RefreshTokenRes) String() string            { return proto1.CompactTextString(m) }
func (*RefreshTokenRes) ProtoMessage()               {}
func (*RefreshTokenRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RefreshTokenRes) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *RefreshTokenRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RefreshTokenRes) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RefreshTokenRes) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RefreshTokenRes) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *RefreshTokenRes) GetExpire() int64 {
	if m != nil {
		return m.Expire
	}
	return 0
}

type DestroyTokenArgs struct {
	UserId     int32  `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	AppFrom    string `protobuf:"bytes,2,opt,name=AppFrom" json:"AppFrom,omitempty"`
	AppChannel string `protobuf:"bytes,3,opt,name=AppChannel" json:"AppChannel,omitempty"`
	AppVer     int32  `protobuf:"varint,4,opt,name=AppVer" json:"AppVer,omitempty"`
	UniqueId   string `protobuf:"bytes,5,opt,name=UniqueId" json:"UniqueId,omitempty"`
	Token      string `protobuf:"bytes,6,opt,name=Token" json:"Token,omitempty"`
}

func (m *DestroyTokenArgs) Reset()                    { *m = DestroyTokenArgs{} }
func (m *DestroyTokenArgs) String() string            { return proto1.CompactTextString(m) }
func (*DestroyTokenArgs) ProtoMessage()               {}
func (*DestroyTokenArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DestroyTokenArgs) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *DestroyTokenArgs) GetAppFrom() string {
	if m != nil {
		return m.AppFrom
	}
	return ""
}

func (m *DestroyTokenArgs) GetAppChannel() string {
	if m != nil {
		return m.AppChannel
	}
	return ""
}

func (m *DestroyTokenArgs) GetAppVer() int32 {
	if m != nil {
		return m.AppVer
	}
	return 0
}

func (m *DestroyTokenArgs) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *DestroyTokenArgs) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DestroyTokenRes struct {
	Ret int32 `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
}

func (m *DestroyTokenRes) Reset()                    { *m = DestroyTokenRes{} }
func (m *DestroyTokenRes) String() string            { return proto1.CompactTextString(m) }
func (*DestroyTokenRes) ProtoMessage()               {}
func (*DestroyTokenRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DestroyTokenRes) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func init() {
	proto1.RegisterType((*AuthArgs)(nil), "proto.AuthArgs")
	proto1.RegisterType((*AuthRes)(nil), "proto.AuthRes")
	proto1.RegisterType((*TokenInfoArgs)(nil), "proto.TokenInfoArgs")
	proto1.RegisterType((*TokenInfoRes)(nil), "proto.TokenInfoRes")
	proto1.RegisterType((*RefreshTokenArgs)(nil), "proto.RefreshTokenArgs")
	proto1.RegisterType((*RefreshTokenRes)(nil), "proto.RefreshTokenRes")
	proto1.RegisterType((*DestroyTokenArgs)(nil), "proto.DestroyTokenArgs")
	proto1.RegisterType((*DestroyTokenRes)(nil), "proto.DestroyTokenRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthService service

type AuthServiceClient interface {
	Auth(ctx context.Context, in *AuthArgs, opts ...grpc.CallOption) (*AuthRes, error)
	TokenInfo(ctx context.Context, in *TokenInfoArgs, opts ...grpc.CallOption) (*TokenInfoRes, error)
	RefreshToken(ctx context.Context, in *RefreshTokenArgs, opts ...grpc.CallOption) (*RefreshTokenRes, error)
	DestroyToken(ctx context.Context, in *DestroyTokenArgs, opts ...grpc.CallOption) (*DestroyTokenRes, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Auth(ctx context.Context, in *AuthArgs, opts ...grpc.CallOption) (*AuthRes, error) {
	out := new(AuthRes)
	err := grpc.Invoke(ctx, "/proto.AuthService/Auth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) TokenInfo(ctx context.Context, in *TokenInfoArgs, opts ...grpc.CallOption) (*TokenInfoRes, error) {
	out := new(TokenInfoRes)
	err := grpc.Invoke(ctx, "/proto.AuthService/TokenInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenArgs, opts ...grpc.CallOption) (*RefreshTokenRes, error) {
	out := new(RefreshTokenRes)
	err := grpc.Invoke(ctx, "/proto.AuthService/RefreshToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DestroyToken(ctx context.Context, in *DestroyTokenArgs, opts ...grpc.CallOption) (*DestroyTokenRes, error) {
	out := new(DestroyTokenRes)
	err := grpc.Invoke(ctx, "/proto.AuthService/DestroyToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceServer interface {
	Auth(context.Context, *AuthArgs) (*AuthRes, error)
	TokenInfo(context.Context, *TokenInfoArgs) (*TokenInfoRes, error)
	RefreshToken(context.Context, *RefreshTokenArgs) (*RefreshTokenRes, error)
	DestroyToken(context.Context, *DestroyTokenArgs) (*DestroyTokenRes, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Auth(ctx, req.(*AuthArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_TokenInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenInfoArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).TokenInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/TokenInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).TokenInfo(ctx, req.(*TokenInfoArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshToken(ctx, req.(*RefreshTokenArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DestroyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyTokenArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DestroyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/DestroyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DestroyToken(ctx, req.(*DestroyTokenArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _AuthService_Auth_Handler,
		},
		{
			MethodName: "TokenInfo",
			Handler:    _AuthService_TokenInfo_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AuthService_RefreshToken_Handler,
		},
		{
			MethodName: "DestroyToken",
			Handler:    _AuthService_DestroyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

func init() { proto1.RegisterFile("auth/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0xad, 0xbf, 0x34, 0x69, 0x7b, 0xbf, 0x42, 0x2a, 0x53, 0x15, 0x2b, 0x03, 0xaa, 0x8c, 0x84,
	0xca, 0x40, 0x91, 0x60, 0x61, 0x8d, 0xf8, 0x91, 0x3a, 0xb0, 0x04, 0xca, 0x5e, 0xa8, 0xdb, 0x54,
	0x40, 0x62, 0x9c, 0x14, 0xc1, 0xd3, 0xb0, 0xb1, 0x20, 0xf1, 0x30, 0x3c, 0x0d, 0x23, 0xb2, 0x93,
	0x5a, 0x4e, 0xa8, 0x10, 0x63, 0x59, 0x12, 0x9f, 0x73, 0xec, 0xab, 0x73, 0xae, 0x7f, 0xc0, 0x1d,
	0xcd, 0xd3, 0x70, 0x5f, 0x7e, 0xfa, 0x5c, 0xc4, 0x69, 0x8c, 0x6d, 0xf5, 0xa3, 0xaf, 0x08, 0xea,
	0xfe, 0x3c, 0x0d, 0x7d, 0x31, 0x4d, 0x70, 0x07, 0x9c, 0x61, 0xc2, 0xc4, 0x60, 0x4c, 0x50, 0x17,
	0xf5, 0xec, 0x20, 0x47, 0x98, 0x40, 0xcd, 0xe7, 0xfc, 0x4c, 0xc4, 0xf7, 0xe4, 0x5f, 0x17, 0xf5,
	0x1a, 0xc1, 0x02, 0xe2, 0x2d, 0x00, 0x9f, 0xf3, 0xe3, 0x70, 0x14, 0x45, 0xec, 0x8e, 0x58, 0x4a,
	0x34, 0x18, 0x59, 0xd1, 0xe7, 0xfc, 0x8a, 0x09, 0x52, 0xcd, 0x2a, 0x66, 0x08, 0x7b, 0x50, 0x1f,
	0x46, 0xb3, 0x87, 0x39, 0x1b, 0x8c, 0x89, 0xad, 0x56, 0x69, 0x8c, 0xdb, 0x60, 0x5f, 0xc6, 0xb7,
	0x2c, 0x22, 0x8e, 0x12, 0x32, 0x40, 0xf7, 0xa0, 0x26, 0x7d, 0x06, 0x2c, 0xc1, 0x2d, 0xb0, 0x02,
	0x96, 0xe6, 0x1e, 0xe5, 0x50, 0x32, 0xe7, 0xc9, 0x34, 0x37, 0x27, 0x87, 0xf4, 0x0d, 0xc1, 0x9a,
	0x5a, 0x38, 0x88, 0x26, 0xf1, 0xca, 0x87, 0xdb, 0x81, 0xa6, 0x36, 0x2b, 0x13, 0x76, 0xc0, 0x39,
	0x7d, 0xe2, 0x33, 0xc1, 0x94, 0x57, 0x2b, 0xc8, 0x11, 0xfd, 0x40, 0xd0, 0x0a, 0xd8, 0x44, 0xb0,
	0x24, 0x54, 0xf3, 0x57, 0x3d, 0x18, 0xa6, 0xd0, 0x34, 0xfd, 0x92, 0x9a, 0x12, 0x0b, 0x1c, 0x7d,
	0x41, 0xe0, 0x9a, 0xc4, 0x2f, 0xb7, 0xd8, 0xc8, 0x6d, 0x15, 0x72, 0x6b, 0x27, 0xd5, 0x9f, 0x9c,
	0xd8, 0xdf, 0x9d, 0x18, 0x6d, 0x77, 0x0a, 0x6d, 0x7f, 0x47, 0xd0, 0x3a, 0x61, 0x49, 0x2a, 0xe2,
	0xe7, 0x3f, 0xd1, 0x76, 0xba, 0x0d, 0xae, 0xe9, 0x77, 0x69, 0x47, 0x0f, 0x3e, 0x11, 0xfc, 0x97,
	0x57, 0xea, 0x82, 0x89, 0xc7, 0xd9, 0x0d, 0xc3, 0xbb, 0x50, 0x95, 0x10, 0xbb, 0xd9, 0x0b, 0xd1,
	0x5f, 0x3c, 0x0b, 0xde, 0xba, 0x41, 0x04, 0x2c, 0xa1, 0x15, 0x7c, 0x04, 0x0d, 0x7d, 0x5e, 0x71,
	0x3b, 0x97, 0x0b, 0xd7, 0xcd, 0xdb, 0x28, 0xb3, 0xd9, 0x4a, 0xbf, 0xb8, 0x0d, 0x78, 0x33, 0x9f,
	0x56, 0x3e, 0xd5, 0x5e, 0x67, 0x89, 0xa0, 0x4b, 0x98, 0xe1, 0x74, 0x89, 0xf2, 0x0e, 0xe9, 0x12,
	0xa5, 0x56, 0xd0, 0xca, 0xb5, 0xa3, 0x84, 0xc3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x5b,
	0x98, 0xd2, 0x16, 0x05, 0x00, 0x00,
}
