// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sts.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	sts.proto

It has these top-level messages:
	GameTableInfoArgs
	CardsInfo
	Player
	StsRes
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

type GameTableInfoArgs struct {
	RoomId  int32        `protobuf:"varint,1,opt,name=room_id,json=roomId" json:"room_id,omitempty"`
	TableId string       `protobuf:"bytes,2,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	Max     int32        `protobuf:"varint,3,opt,name=max" json:"max,omitempty"`
	Start   int32        `protobuf:"varint,4,opt,name=start" json:"start,omitempty"`
	End     int32        `protobuf:"varint,5,opt,name=end" json:"end,omitempty"`
	Cards   []*CardsInfo `protobuf:"bytes,6,rep,name=cards" json:"cards,omitempty"`
	Button  int32        `protobuf:"varint,7,opt,name=button" json:"button,omitempty"`
	Sb      int32        `protobuf:"varint,8,opt,name=sb" json:"sb,omitempty"`
	Bb      int32        `protobuf:"varint,9,opt,name=bb" json:"bb,omitempty"`
	SbPos   int32        `protobuf:"varint,10,opt,name=sb_pos,json=sbPos" json:"sb_pos,omitempty"`
	BbPos   int32        `protobuf:"varint,11,opt,name=bb_pos,json=bbPos" json:"bb_pos,omitempty"`
	Pot     []int32      `protobuf:"varint,12,rep,packed,name=pot" json:"pot,omitempty"`
	Player  []*Player    `protobuf:"bytes,13,rep,name=player" json:"player,omitempty"`
}

func (m *GameTableInfoArgs) Reset()                    { *m = GameTableInfoArgs{} }
func (m *GameTableInfoArgs) String() string            { return proto1.CompactTextString(m) }
func (*GameTableInfoArgs) ProtoMessage()               {}
func (*GameTableInfoArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GameTableInfoArgs) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *GameTableInfoArgs) GetTableId() string {
	if m != nil {
		return m.TableId
	}
	return ""
}

func (m *GameTableInfoArgs) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *GameTableInfoArgs) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *GameTableInfoArgs) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *GameTableInfoArgs) GetCards() []*CardsInfo {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *GameTableInfoArgs) GetButton() int32 {
	if m != nil {
		return m.Button
	}
	return 0
}

func (m *GameTableInfoArgs) GetSb() int32 {
	if m != nil {
		return m.Sb
	}
	return 0
}

func (m *GameTableInfoArgs) GetBb() int32 {
	if m != nil {
		return m.Bb
	}
	return 0
}

func (m *GameTableInfoArgs) GetSbPos() int32 {
	if m != nil {
		return m.SbPos
	}
	return 0
}

func (m *GameTableInfoArgs) GetBbPos() int32 {
	if m != nil {
		return m.BbPos
	}
	return 0
}

func (m *GameTableInfoArgs) GetPot() []int32 {
	if m != nil {
		return m.Pot
	}
	return nil
}

func (m *GameTableInfoArgs) GetPlayer() []*Player {
	if m != nil {
		return m.Player
	}
	return nil
}

type CardsInfo struct {
	Suit  int32 `protobuf:"varint,1,opt,name=suit" json:"suit,omitempty"`
	Value int32 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *CardsInfo) Reset()                    { *m = CardsInfo{} }
func (m *CardsInfo) String() string            { return proto1.CompactTextString(m) }
func (*CardsInfo) ProtoMessage()               {}
func (*CardsInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CardsInfo) GetSuit() int32 {
	if m != nil {
		return m.Suit
	}
	return 0
}

func (m *CardsInfo) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Player struct {
	Id             int32        `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Nickname       string       `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	Avatar         string       `protobuf:"bytes,3,opt,name=avatar" json:"avatar,omitempty"`
	Pos            int32        `protobuf:"varint,4,opt,name=pos" json:"pos,omitempty"`
	Bet            int32        `protobuf:"varint,5,opt,name=bet" json:"bet,omitempty"`
	Win            int32        `protobuf:"varint,6,opt,name=win" json:"win,omitempty"`
	FormerChips    int32        `protobuf:"varint,7,opt,name=former_chips,json=formerChips" json:"former_chips,omitempty"`
	CurrentChips   int32        `protobuf:"varint,8,opt,name=current_chips,json=currentChips" json:"current_chips,omitempty"`
	Action         string       `protobuf:"bytes,9,opt,name=action" json:"action,omitempty"`
	Cards          []*CardsInfo `protobuf:"bytes,10,rep,name=cards" json:"cards,omitempty"`
	HandLevel      int32        `protobuf:"varint,11,opt,name=hand_level,json=handLevel" json:"hand_level,omitempty"`
	HandFinalValue int32        `protobuf:"varint,12,opt,name=hand_final_value,json=handFinalValue" json:"hand_final_value,omitempty"`
}

func (m *Player) Reset()                    { *m = Player{} }
func (m *Player) String() string            { return proto1.CompactTextString(m) }
func (*Player) ProtoMessage()               {}
func (*Player) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Player) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Player) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Player) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Player) GetPos() int32 {
	if m != nil {
		return m.Pos
	}
	return 0
}

func (m *Player) GetBet() int32 {
	if m != nil {
		return m.Bet
	}
	return 0
}

func (m *Player) GetWin() int32 {
	if m != nil {
		return m.Win
	}
	return 0
}

func (m *Player) GetFormerChips() int32 {
	if m != nil {
		return m.FormerChips
	}
	return 0
}

func (m *Player) GetCurrentChips() int32 {
	if m != nil {
		return m.CurrentChips
	}
	return 0
}

func (m *Player) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Player) GetCards() []*CardsInfo {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *Player) GetHandLevel() int32 {
	if m != nil {
		return m.HandLevel
	}
	return 0
}

func (m *Player) GetHandFinalValue() int32 {
	if m != nil {
		return m.HandFinalValue
	}
	return 0
}

type StsRes struct {
	Ret int32  `protobuf:"varint,1,opt,name=ret" json:"ret,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *StsRes) Reset()                    { *m = StsRes{} }
func (m *StsRes) String() string            { return proto1.CompactTextString(m) }
func (*StsRes) ProtoMessage()               {}
func (*StsRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StsRes) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *StsRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto1.RegisterType((*GameTableInfoArgs)(nil), "proto.GameTableInfoArgs")
	proto1.RegisterType((*CardsInfo)(nil), "proto.CardsInfo")
	proto1.RegisterType((*Player)(nil), "proto.Player")
	proto1.RegisterType((*StsRes)(nil), "proto.StsRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StsService service

type StsServiceClient interface {
	// 获取游戏信息
	GameInfo(ctx context.Context, in *GameTableInfoArgs, opts ...grpc.CallOption) (*StsRes, error)
}

type stsServiceClient struct {
	cc *grpc.ClientConn
}

func NewStsServiceClient(cc *grpc.ClientConn) StsServiceClient {
	return &stsServiceClient{cc}
}

func (c *stsServiceClient) GameInfo(ctx context.Context, in *GameTableInfoArgs, opts ...grpc.CallOption) (*StsRes, error) {
	out := new(StsRes)
	err := grpc.Invoke(ctx, "/proto.StsService/GameInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StsService service

type StsServiceServer interface {
	// 获取游戏信息
	GameInfo(context.Context, *GameTableInfoArgs) (*StsRes, error)
}

func RegisterStsServiceServer(s *grpc.Server, srv StsServiceServer) {
	s.RegisterService(&_StsService_serviceDesc, srv)
}

func _StsService_GameInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameTableInfoArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StsServiceServer).GameInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StsService/GameInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StsServiceServer).GameInfo(ctx, req.(*GameTableInfoArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _StsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StsService",
	HandlerType: (*StsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GameInfo",
			Handler:    _StsService_GameInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sts.proto",
}

func init() { proto1.RegisterFile("sts.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4f, 0x8b, 0xdb, 0x3c,
	0x10, 0xc6, 0xdf, 0x38, 0x6b, 0x27, 0x9e, 0x24, 0x4b, 0x5e, 0xd1, 0x3f, 0xea, 0x42, 0x21, 0x4d,
	0x69, 0xc9, 0xa1, 0xec, 0x61, 0xcb, 0x7e, 0x80, 0x12, 0x68, 0x09, 0xf4, 0xb0, 0x38, 0xa5, 0xd7,
	0x20, 0xd9, 0xca, 0xae, 0xa8, 0x2d, 0x19, 0x49, 0xc9, 0xb6, 0xdf, 0xb0, 0xf7, 0x7e, 0xa1, 0x32,
	0x23, 0x25, 0x3d, 0x14, 0x7a, 0xf2, 0x3c, 0xbf, 0x19, 0xec, 0xc7, 0xcf, 0x48, 0x50, 0xfa, 0xe0,
	0xaf, 0x7b, 0x67, 0x83, 0x65, 0x39, 0x3d, 0x96, 0x3f, 0x33, 0xf8, 0xff, 0x93, 0xe8, 0xd4, 0x17,
	0x21, 0x5b, 0xb5, 0x31, 0x7b, 0xfb, 0xc1, 0xdd, 0x7b, 0xf6, 0x1c, 0x46, 0xce, 0xda, 0x6e, 0xa7,
	0x1b, 0x3e, 0x58, 0x0c, 0x56, 0x79, 0x55, 0xa0, 0xdc, 0x34, 0xec, 0x05, 0x8c, 0x03, 0x4e, 0x62,
	0x27, 0x5b, 0x0c, 0x56, 0x65, 0x35, 0x22, 0xbd, 0x69, 0xd8, 0x1c, 0x86, 0x9d, 0xf8, 0xce, 0x87,
	0x34, 0x8f, 0x25, 0x7b, 0x02, 0xb9, 0x0f, 0xc2, 0x05, 0x7e, 0x41, 0x2c, 0x0a, 0x9c, 0x53, 0xa6,
	0xe1, 0x79, 0x9c, 0x53, 0xa6, 0x61, 0x6f, 0x21, 0xaf, 0x85, 0x6b, 0x3c, 0x2f, 0x16, 0xc3, 0xd5,
	0xe4, 0x66, 0x1e, 0x1d, 0x5e, 0xaf, 0x91, 0xa1, 0xa5, 0x2a, 0xb6, 0xd9, 0x33, 0x28, 0xe4, 0x21,
	0x04, 0x6b, 0xf8, 0x28, 0x9a, 0x8a, 0x8a, 0x5d, 0x42, 0xe6, 0x25, 0x1f, 0x13, 0xcb, 0xbc, 0x44,
	0x2d, 0x25, 0x2f, 0xa3, 0x96, 0x92, 0x3d, 0x85, 0xc2, 0xcb, 0x5d, 0x6f, 0x3d, 0x87, 0x64, 0x44,
	0xde, 0x59, 0x8f, 0x58, 0x46, 0x3c, 0x89, 0x58, 0x12, 0x9e, 0xc3, 0xb0, 0xb7, 0x81, 0x4f, 0x17,
	0x43, 0xf4, 0xd7, 0xdb, 0xc0, 0xde, 0x40, 0xd1, 0xb7, 0xe2, 0x87, 0x72, 0x7c, 0x46, 0x06, 0x67,
	0xc9, 0xe0, 0x1d, 0xc1, 0x2a, 0x35, 0x97, 0xb7, 0x50, 0x9e, 0x2d, 0x33, 0x06, 0x17, 0xfe, 0xa0,
	0x43, 0x8a, 0x8f, 0x6a, 0xcc, 0xe3, 0x28, 0xda, 0x83, 0xa2, 0xe4, 0xf2, 0x2a, 0x8a, 0xe5, 0xaf,
	0x0c, 0x8a, 0xf8, 0x26, 0x34, 0x7e, 0x4e, 0x3c, 0xd3, 0x0d, 0xbb, 0x82, 0xb1, 0xd1, 0xf5, 0x37,
	0x23, 0x3a, 0x95, 0xd2, 0x3e, 0x6b, 0x0c, 0x43, 0x1c, 0x45, 0x10, 0x8e, 0x12, 0x2f, 0xab, 0xa4,
	0xa2, 0x7d, 0x9f, 0x22, 0xc7, 0x12, 0x89, 0x54, 0xe1, 0x14, 0xb8, 0x54, 0xb4, 0x82, 0x47, 0x6d,
	0x78, 0x11, 0xc9, 0xa3, 0x36, 0xec, 0x15, 0x4c, 0xf7, 0xd6, 0x75, 0xca, 0xed, 0xea, 0x07, 0xdd,
	0xfb, 0x14, 0xf0, 0x24, 0xb2, 0x35, 0x22, 0xf6, 0x1a, 0x66, 0xf5, 0xc1, 0x39, 0x65, 0x42, 0x9a,
	0x89, 0x81, 0x4f, 0x13, 0x8c, 0x43, 0xe8, 0xaa, 0x0e, 0xda, 0x1a, 0x8a, 0x1f, 0x5d, 0x91, 0xfa,
	0xb3, 0x62, 0xf8, 0xf7, 0x8a, 0x5f, 0x02, 0x3c, 0x08, 0xd3, 0xec, 0x5a, 0x75, 0x54, 0x6d, 0xda,
	0x4b, 0x89, 0xe4, 0x33, 0x02, 0xb6, 0x82, 0x39, 0xb5, 0xf7, 0xda, 0x88, 0x76, 0x17, 0xc3, 0x9c,
	0xd2, 0xd0, 0x25, 0xf2, 0x8f, 0x88, 0xbf, 0x52, 0xaa, 0xef, 0xa0, 0xd8, 0x06, 0x5f, 0x29, 0xfa,
	0x7d, 0xa7, 0x4e, 0x8b, 0xc0, 0x92, 0x4e, 0xaa, 0xbf, 0x4f, 0x89, 0x62, 0x79, 0xb3, 0x06, 0xd8,
	0x06, 0xbf, 0x55, 0xee, 0xa8, 0x6b, 0xc5, 0x6e, 0x61, 0x8c, 0x57, 0x82, 0xf6, 0xc8, 0x93, 0xd3,
	0xbf, 0xee, 0xc8, 0xd5, 0xe9, 0x14, 0xc4, 0xcf, 0x2c, 0xff, 0x93, 0x05, 0xe9, 0xf7, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xe1, 0x5a, 0x52, 0xa3, 0x65, 0x03, 0x00, 0x00,
}