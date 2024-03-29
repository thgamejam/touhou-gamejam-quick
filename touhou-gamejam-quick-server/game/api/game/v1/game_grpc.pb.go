// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameClient interface {
	CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*CreateGameReply, error)
	UpdateGame(ctx context.Context, in *UpdateGameRequest, opts ...grpc.CallOption) (*UpdateGameReply, error)
	DeleteGame(ctx context.Context, in *DeleteGameRequest, opts ...grpc.CallOption) (*DeleteGameReply, error)
	GetGame(ctx context.Context, in *GetGameRequest, opts ...grpc.CallOption) (*GetGameReply, error)
	GetGameDownload(ctx context.Context, in *GetGameDownloadRequest, opts ...grpc.CallOption) (*GetGameDownloadReply, error)
	ListGame(ctx context.Context, in *ListGameRequest, opts ...grpc.CallOption) (*ListGameReply, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*CreateGameReply, error) {
	out := new(CreateGameReply)
	err := c.cc.Invoke(ctx, "/touhou.api.game.v1.Game/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) UpdateGame(ctx context.Context, in *UpdateGameRequest, opts ...grpc.CallOption) (*UpdateGameReply, error) {
	out := new(UpdateGameReply)
	err := c.cc.Invoke(ctx, "/touhou.api.game.v1.Game/UpdateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) DeleteGame(ctx context.Context, in *DeleteGameRequest, opts ...grpc.CallOption) (*DeleteGameReply, error) {
	out := new(DeleteGameReply)
	err := c.cc.Invoke(ctx, "/touhou.api.game.v1.Game/DeleteGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GetGame(ctx context.Context, in *GetGameRequest, opts ...grpc.CallOption) (*GetGameReply, error) {
	out := new(GetGameReply)
	err := c.cc.Invoke(ctx, "/touhou.api.game.v1.Game/GetGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GetGameDownload(ctx context.Context, in *GetGameDownloadRequest, opts ...grpc.CallOption) (*GetGameDownloadReply, error) {
	out := new(GetGameDownloadReply)
	err := c.cc.Invoke(ctx, "/touhou.api.game.v1.Game/GetGameDownload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) ListGame(ctx context.Context, in *ListGameRequest, opts ...grpc.CallOption) (*ListGameReply, error) {
	out := new(ListGameReply)
	err := c.cc.Invoke(ctx, "/touhou.api.game.v1.Game/ListGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServer is the server API for Game service.
// All implementations must embed UnimplementedGameServer
// for forward compatibility
type GameServer interface {
	CreateGame(context.Context, *CreateGameRequest) (*CreateGameReply, error)
	UpdateGame(context.Context, *UpdateGameRequest) (*UpdateGameReply, error)
	DeleteGame(context.Context, *DeleteGameRequest) (*DeleteGameReply, error)
	GetGame(context.Context, *GetGameRequest) (*GetGameReply, error)
	GetGameDownload(context.Context, *GetGameDownloadRequest) (*GetGameDownloadReply, error)
	ListGame(context.Context, *ListGameRequest) (*ListGameReply, error)
	mustEmbedUnimplementedGameServer()
}

// UnimplementedGameServer must be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (UnimplementedGameServer) CreateGame(context.Context, *CreateGameRequest) (*CreateGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (UnimplementedGameServer) UpdateGame(context.Context, *UpdateGameRequest) (*UpdateGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGame not implemented")
}
func (UnimplementedGameServer) DeleteGame(context.Context, *DeleteGameRequest) (*DeleteGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGame not implemented")
}
func (UnimplementedGameServer) GetGame(context.Context, *GetGameRequest) (*GetGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGame not implemented")
}
func (UnimplementedGameServer) GetGameDownload(context.Context, *GetGameDownloadRequest) (*GetGameDownloadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameDownload not implemented")
}
func (UnimplementedGameServer) ListGame(context.Context, *ListGameRequest) (*ListGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGame not implemented")
}
func (UnimplementedGameServer) mustEmbedUnimplementedGameServer() {}

// UnsafeGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServer will
// result in compilation errors.
type UnsafeGameServer interface {
	mustEmbedUnimplementedGameServer()
}

func RegisterGameServer(s grpc.ServiceRegistrar, srv GameServer) {
	s.RegisterService(&Game_ServiceDesc, srv)
}

func _Game_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/touhou.api.game.v1.Game/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).CreateGame(ctx, req.(*CreateGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_UpdateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).UpdateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/touhou.api.game.v1.Game/UpdateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).UpdateGame(ctx, req.(*UpdateGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_DeleteGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).DeleteGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/touhou.api.game.v1.Game/DeleteGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).DeleteGame(ctx, req.(*DeleteGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GetGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/touhou.api.game.v1.Game/GetGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetGame(ctx, req.(*GetGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GetGameDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetGameDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/touhou.api.game.v1.Game/GetGameDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetGameDownload(ctx, req.(*GetGameDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_ListGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).ListGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/touhou.api.game.v1.Game/ListGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).ListGame(ctx, req.(*ListGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Game_ServiceDesc is the grpc.ServiceDesc for Game service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Game_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "touhou.api.game.v1.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGame",
			Handler:    _Game_CreateGame_Handler,
		},
		{
			MethodName: "UpdateGame",
			Handler:    _Game_UpdateGame_Handler,
		},
		{
			MethodName: "DeleteGame",
			Handler:    _Game_DeleteGame_Handler,
		},
		{
			MethodName: "GetGame",
			Handler:    _Game_GetGame_Handler,
		},
		{
			MethodName: "GetGameDownload",
			Handler:    _Game_GetGameDownload_Handler,
		},
		{
			MethodName: "ListGame",
			Handler:    _Game_ListGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/game/v1/game.proto",
}
