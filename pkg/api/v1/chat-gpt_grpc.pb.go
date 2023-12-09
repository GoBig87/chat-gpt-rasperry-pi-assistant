// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/v1/chat-gpt.proto

package api

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

// ChatGptServiceClient is the client API for ChatGptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatGptServiceClient interface {
	// Process Prompt
	//
	// Sends a prompt to chatGPT and returns the response.
	ProcessPrompt(ctx context.Context, in *ProcessPromptRequest, opts ...grpc.CallOption) (*ProcessPromptResponse, error)
}

type chatGptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatGptServiceClient(cc grpc.ClientConnInterface) ChatGptServiceClient {
	return &chatGptServiceClient{cc}
}

func (c *chatGptServiceClient) ProcessPrompt(ctx context.Context, in *ProcessPromptRequest, opts ...grpc.CallOption) (*ProcessPromptResponse, error) {
	out := new(ProcessPromptResponse)
	err := c.cc.Invoke(ctx, "/api.v1.chat_gpt.ChatGptService/ProcessPrompt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatGptServiceServer is the server API for ChatGptService service.
// All implementations must embed UnimplementedChatGptServiceServer
// for forward compatibility
type ChatGptServiceServer interface {
	// Process Prompt
	//
	// Sends a prompt to chatGPT and returns the response.
	ProcessPrompt(context.Context, *ProcessPromptRequest) (*ProcessPromptResponse, error)
	mustEmbedUnimplementedChatGptServiceServer()
}

// UnimplementedChatGptServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatGptServiceServer struct {
}

func (UnimplementedChatGptServiceServer) ProcessPrompt(context.Context, *ProcessPromptRequest) (*ProcessPromptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessPrompt not implemented")
}
func (UnimplementedChatGptServiceServer) mustEmbedUnimplementedChatGptServiceServer() {}

// UnsafeChatGptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatGptServiceServer will
// result in compilation errors.
type UnsafeChatGptServiceServer interface {
	mustEmbedUnimplementedChatGptServiceServer()
}

func RegisterChatGptServiceServer(s grpc.ServiceRegistrar, srv ChatGptServiceServer) {
	s.RegisterService(&ChatGptService_ServiceDesc, srv)
}

func _ChatGptService_ProcessPrompt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessPromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGptServiceServer).ProcessPrompt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.chat_gpt.ChatGptService/ProcessPrompt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGptServiceServer).ProcessPrompt(ctx, req.(*ProcessPromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatGptService_ServiceDesc is the grpc.ServiceDesc for ChatGptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatGptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.chat_gpt.ChatGptService",
	HandlerType: (*ChatGptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessPrompt",
			Handler:    _ChatGptService_ProcessPrompt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/chat-gpt.proto",
}