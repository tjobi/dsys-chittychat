// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ChittyChatClient is the client API for ChittyChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittyChatClient interface {
	JoinRoom(ctx context.Context, in *ClientJoin, opts ...grpc.CallOption) (*ServerWelcome, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChittyChat_SendMessageClient, error)
}

type chittyChatClient struct {
	cc grpc.ClientConnInterface
}

func NewChittyChatClient(cc grpc.ClientConnInterface) ChittyChatClient {
	return &chittyChatClient{cc}
}

func (c *chittyChatClient) JoinRoom(ctx context.Context, in *ClientJoin, opts ...grpc.CallOption) (*ServerWelcome, error) {
	out := new(ServerWelcome)
	err := c.cc.Invoke(ctx, "/grpcBrr.ChittyChat/JoinRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChittyChat_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChittyChat_ServiceDesc.Streams[0], "/grpcBrr.ChittyChat/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chittyChatSendMessageClient{stream}
	return x, nil
}

type ChittyChat_SendMessageClient interface {
	Send(*ChatMessage) error
	Recv() (*ServerReponse, error)
	grpc.ClientStream
}

type chittyChatSendMessageClient struct {
	grpc.ClientStream
}

func (x *chittyChatSendMessageClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chittyChatSendMessageClient) Recv() (*ServerReponse, error) {
	m := new(ServerReponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChittyChatServer is the server API for ChittyChat service.
// All implementations must embed UnimplementedChittyChatServer
// for forward compatibility
type ChittyChatServer interface {
	JoinRoom(context.Context, *ClientJoin) (*ServerWelcome, error)
	SendMessage(ChittyChat_SendMessageServer) error
	mustEmbedUnimplementedChittyChatServer()
}

// UnimplementedChittyChatServer must be embedded to have forward compatible implementations.
type UnimplementedChittyChatServer struct {
}

func (UnimplementedChittyChatServer) JoinRoom(context.Context, *ClientJoin) (*ServerWelcome, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (UnimplementedChittyChatServer) SendMessage(ChittyChat_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChittyChatServer) mustEmbedUnimplementedChittyChatServer() {}

// UnsafeChittyChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittyChatServer will
// result in compilation errors.
type UnsafeChittyChatServer interface {
	mustEmbedUnimplementedChittyChatServer()
}

func RegisterChittyChatServer(s grpc.ServiceRegistrar, srv ChittyChatServer) {
	s.RegisterService(&ChittyChat_ServiceDesc, srv)
}

func _ChittyChat_JoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientJoin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServer).JoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcBrr.ChittyChat/JoinRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServer).JoinRoom(ctx, req.(*ClientJoin))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChat_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChittyChatServer).SendMessage(&chittyChatSendMessageServer{stream})
}

type ChittyChat_SendMessageServer interface {
	Send(*ServerReponse) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type chittyChatSendMessageServer struct {
	grpc.ServerStream
}

func (x *chittyChatSendMessageServer) Send(m *ServerReponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chittyChatSendMessageServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChittyChat_ServiceDesc is the grpc.ServiceDesc for ChittyChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChittyChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcBrr.ChittyChat",
	HandlerType: (*ChittyChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinRoom",
			Handler:    _ChittyChat_JoinRoom_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _ChittyChat_SendMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/brr.proto",
}
