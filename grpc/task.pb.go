// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

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

// message GameInfoArgs {
//    int32 room_id = 1; //房间场类型
//    int32 match_type = 2;  //赛事类型
//    int32 winner = 3; //胜者id
//    repeated TaskPlayerInfo players = 4;
//    int32 time = 5;  //时间戳
// }
//
// // 玩家信息
// message TaskPlayerInfo{
//    int32 id = 1; // 玩家id
//    int32 hand_level = 2;  // 牌型
//    int32 is_allin = 3; //是否全下0否1是
//    int32 all_bet = 4; //全部下注
// }
type TaskRes struct {
	Ret int32  `protobuf:"varint,1,opt,name=ret" json:"ret,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *TaskRes) Reset()                    { *m = TaskRes{} }
func (m *TaskRes) String() string            { return proto1.CompactTextString(m) }
func (*TaskRes) ProtoMessage()               {}
func (*TaskRes) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *TaskRes) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *TaskRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PlayerActionArgs struct {
	RoomId    int32 `protobuf:"varint,1,opt,name=room_id,json=roomId" json:"room_id,omitempty"`
	MatchType int32 `protobuf:"varint,2,opt,name=match_type,json=matchType" json:"match_type,omitempty"`
	Id        int32 `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	Type      int32 `protobuf:"varint,4,opt,name=type" json:"type,omitempty"`
	Time      int32 `protobuf:"varint,5,opt,name=time" json:"time,omitempty"`
}

func (m *PlayerActionArgs) Reset()                    { *m = PlayerActionArgs{} }
func (m *PlayerActionArgs) String() string            { return proto1.CompactTextString(m) }
func (*PlayerActionArgs) ProtoMessage()               {}
func (*PlayerActionArgs) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *PlayerActionArgs) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *PlayerActionArgs) GetMatchType() int32 {
	if m != nil {
		return m.MatchType
	}
	return 0
}

func (m *PlayerActionArgs) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PlayerActionArgs) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *PlayerActionArgs) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

type UpsetTaskArgs struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *UpsetTaskArgs) Reset()                    { *m = UpsetTaskArgs{} }
func (m *UpsetTaskArgs) String() string            { return proto1.CompactTextString(m) }
func (*UpsetTaskArgs) ProtoMessage()               {}
func (*UpsetTaskArgs) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *UpsetTaskArgs) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto1.RegisterType((*TaskRes)(nil), "proto.TaskRes")
	proto1.RegisterType((*PlayerActionArgs)(nil), "proto.PlayerActionArgs")
	proto1.RegisterType((*UpsetTaskArgs)(nil), "proto.UpsetTaskArgs")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskService service

type TaskServiceClient interface {
	// 游戏结束后信息
	GameOver(ctx context.Context, in *GameTableInfoArgs, opts ...grpc.CallOption) (*TaskRes, error)
	//  玩家动作信息
	PlayerEvent(ctx context.Context, in *PlayerActionArgs, opts ...grpc.CallOption) (*TaskRes, error)
	// 更新任务
	UpsetTask(ctx context.Context, in *UpsetTaskArgs, opts ...grpc.CallOption) (*TaskRes, error)
}

type taskServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskServiceClient(cc *grpc.ClientConn) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) GameOver(ctx context.Context, in *GameTableInfoArgs, opts ...grpc.CallOption) (*TaskRes, error) {
	out := new(TaskRes)
	err := grpc.Invoke(ctx, "/proto.TaskService/GameOver", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) PlayerEvent(ctx context.Context, in *PlayerActionArgs, opts ...grpc.CallOption) (*TaskRes, error) {
	out := new(TaskRes)
	err := grpc.Invoke(ctx, "/proto.TaskService/PlayerEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) UpsetTask(ctx context.Context, in *UpsetTaskArgs, opts ...grpc.CallOption) (*TaskRes, error) {
	out := new(TaskRes)
	err := grpc.Invoke(ctx, "/proto.TaskService/UpsetTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskService service

type TaskServiceServer interface {
	// 游戏结束后信息
	GameOver(context.Context, *GameTableInfoArgs) (*TaskRes, error)
	//  玩家动作信息
	PlayerEvent(context.Context, *PlayerActionArgs) (*TaskRes, error)
	// 更新任务
	UpsetTask(context.Context, *UpsetTaskArgs) (*TaskRes, error)
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_GameOver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameTableInfoArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GameOver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TaskService/GameOver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GameOver(ctx, req.(*GameTableInfoArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_PlayerEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerActionArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).PlayerEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TaskService/PlayerEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).PlayerEvent(ctx, req.(*PlayerActionArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_UpsetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsetTaskArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).UpsetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TaskService/UpsetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).UpsetTask(ctx, req.(*UpsetTaskArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GameOver",
			Handler:    _TaskService_GameOver_Handler,
		},
		{
			MethodName: "PlayerEvent",
			Handler:    _TaskService_PlayerEvent_Handler,
		},
		{
			MethodName: "UpsetTask",
			Handler:    _TaskService_UpsetTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}

func init() { proto1.RegisterFile("task.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xdf, 0x6d, 0x9b, 0xf6, 0xcd, 0x14, 0x4b, 0x19, 0x84, 0x86, 0x82, 0x58, 0x72, 0xea,
	0xc5, 0x1e, 0x2c, 0x88, 0xd7, 0x1e, 0x44, 0x7a, 0x52, 0x62, 0x3c, 0x97, 0x6d, 0x32, 0xd6, 0xa5,
	0xdd, 0x6c, 0xd8, 0x5d, 0x02, 0xb9, 0xfa, 0x99, 0xfc, 0x80, 0xb2, 0x9b, 0x55, 0xf0, 0xcf, 0x69,
	0x67, 0x7f, 0x33, 0xcf, 0xc3, 0x33, 0x03, 0x60, 0xb9, 0x39, 0xae, 0x6a, 0xad, 0xac, 0xc2, 0xc8,
	0x3f, 0xf3, 0xd8, 0x58, 0xd3, 0x91, 0xf4, 0x0a, 0x46, 0x39, 0x37, 0xc7, 0x8c, 0x0c, 0x4e, 0xa1,
	0xaf, 0xc9, 0x26, 0x6c, 0xc1, 0x96, 0x51, 0xe6, 0x4a, 0x47, 0xa4, 0x39, 0x24, 0xbd, 0x05, 0x5b,
	0xc6, 0x99, 0x2b, 0xd3, 0x37, 0x06, 0xd3, 0xc7, 0x13, 0x6f, 0x49, 0x6f, 0x0a, 0x2b, 0x54, 0xb5,
	0xd1, 0x07, 0x83, 0x33, 0x18, 0x69, 0xa5, 0xe4, 0x4e, 0x94, 0x41, 0x3c, 0x74, 0xdf, 0x6d, 0x89,
	0x17, 0x00, 0x92, 0xdb, 0xe2, 0x75, 0x67, 0xdb, 0x9a, 0xbc, 0x4d, 0x94, 0xc5, 0x9e, 0xe4, 0x6d,
	0x4d, 0x38, 0x81, 0x9e, 0x28, 0x93, 0xbe, 0xc7, 0x3d, 0x51, 0x22, 0xc2, 0xc0, 0x0f, 0x0e, 0x3c,
	0xf1, 0xb5, 0x67, 0x42, 0x52, 0x12, 0x05, 0x26, 0x24, 0xa5, 0x97, 0x70, 0xf6, 0x5c, 0x1b, 0xb2,
	0x2e, 0xb8, 0x0f, 0xd0, 0x19, 0xb1, 0x4f, 0xa3, 0xeb, 0x77, 0x06, 0x63, 0xd7, 0x7c, 0x22, 0xdd,
	0x88, 0x82, 0xf0, 0x06, 0xfe, 0xdf, 0x73, 0x49, 0x0f, 0x0d, 0x69, 0x4c, 0xba, 0xc5, 0x57, 0x0e,
	0xe4, 0x7c, 0x7f, 0xa2, 0x6d, 0xf5, 0xa2, 0x9c, 0xcb, 0x7c, 0x12, 0x3a, 0xe1, 0x1e, 0xe9, 0x3f,
	0xbc, 0x85, 0x71, 0xb7, 0xec, 0x5d, 0x43, 0x95, 0xc5, 0x59, 0x18, 0xf8, 0x79, 0x80, 0x3f, 0x94,
	0x6b, 0x88, 0xbf, 0x22, 0xe2, 0x79, 0x68, 0x7f, 0x0b, 0xfd, 0x5b, 0xb4, 0x1f, 0x7a, 0xb0, 0xfe,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xb2, 0x77, 0xf2, 0xb2, 0x01, 0x00, 0x00,
}
