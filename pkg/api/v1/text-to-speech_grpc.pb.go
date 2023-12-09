// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/v1/text-to-speech.proto

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

// TextToSpeechServiceClient is the client API for TextToSpeechService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextToSpeechServiceClient interface {
	// Start Processing Text to Audio
	//
	// This takes in an input of text and produces
	// an audio out put of it.
	ProcessText(ctx context.Context, in *ProcessTextRequest, opts ...grpc.CallOption) (TextToSpeechService_ProcessTextClient, error)
}

type textToSpeechServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTextToSpeechServiceClient(cc grpc.ClientConnInterface) TextToSpeechServiceClient {
	return &textToSpeechServiceClient{cc}
}

func (c *textToSpeechServiceClient) ProcessText(ctx context.Context, in *ProcessTextRequest, opts ...grpc.CallOption) (TextToSpeechService_ProcessTextClient, error) {
	stream, err := c.cc.NewStream(ctx, &TextToSpeechService_ServiceDesc.Streams[0], "/api.v1.text_to_speech.TextToSpeechService/ProcessText", opts...)
	if err != nil {
		return nil, err
	}
	x := &textToSpeechServiceProcessTextClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TextToSpeechService_ProcessTextClient interface {
	Recv() (*ProcessTextResponse, error)
	grpc.ClientStream
}

type textToSpeechServiceProcessTextClient struct {
	grpc.ClientStream
}

func (x *textToSpeechServiceProcessTextClient) Recv() (*ProcessTextResponse, error) {
	m := new(ProcessTextResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TextToSpeechServiceServer is the server API for TextToSpeechService service.
// All implementations must embed UnimplementedTextToSpeechServiceServer
// for forward compatibility
type TextToSpeechServiceServer interface {
	// Start Processing Text to Audio
	//
	// This takes in an input of text and produces
	// an audio out put of it.
	ProcessText(*ProcessTextRequest, TextToSpeechService_ProcessTextServer) error
	mustEmbedUnimplementedTextToSpeechServiceServer()
}

// UnimplementedTextToSpeechServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTextToSpeechServiceServer struct {
}

func (UnimplementedTextToSpeechServiceServer) ProcessText(*ProcessTextRequest, TextToSpeechService_ProcessTextServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessText not implemented")
}
func (UnimplementedTextToSpeechServiceServer) mustEmbedUnimplementedTextToSpeechServiceServer() {}

// UnsafeTextToSpeechServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextToSpeechServiceServer will
// result in compilation errors.
type UnsafeTextToSpeechServiceServer interface {
	mustEmbedUnimplementedTextToSpeechServiceServer()
}

func RegisterTextToSpeechServiceServer(s grpc.ServiceRegistrar, srv TextToSpeechServiceServer) {
	s.RegisterService(&TextToSpeechService_ServiceDesc, srv)
}

func _TextToSpeechService_ProcessText_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProcessTextRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TextToSpeechServiceServer).ProcessText(m, &textToSpeechServiceProcessTextServer{stream})
}

type TextToSpeechService_ProcessTextServer interface {
	Send(*ProcessTextResponse) error
	grpc.ServerStream
}

type textToSpeechServiceProcessTextServer struct {
	grpc.ServerStream
}

func (x *textToSpeechServiceProcessTextServer) Send(m *ProcessTextResponse) error {
	return x.ServerStream.SendMsg(m)
}

// TextToSpeechService_ServiceDesc is the grpc.ServiceDesc for TextToSpeechService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextToSpeechService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.text_to_speech.TextToSpeechService",
	HandlerType: (*TextToSpeechServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProcessText",
			Handler:       _TextToSpeechService_ProcessText_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v1/text-to-speech.proto",
}