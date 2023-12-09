// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/v1/speech-to-text.proto

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

// SpeechToTextServiceClient is the client API for SpeechToTextService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpeechToTextServiceClient interface {
	// Start Processing Audio
	//
	// This is kicked off when the wake word is registered.
	// When speech is no longer detected, the response will return
	// false and the wake word will begin listening again.
	ProcessSpeech(ctx context.Context, in *ProcessSpeechRequest, opts ...grpc.CallOption) (SpeechToTextService_ProcessSpeechClient, error)
}

type speechToTextServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpeechToTextServiceClient(cc grpc.ClientConnInterface) SpeechToTextServiceClient {
	return &speechToTextServiceClient{cc}
}

func (c *speechToTextServiceClient) ProcessSpeech(ctx context.Context, in *ProcessSpeechRequest, opts ...grpc.CallOption) (SpeechToTextService_ProcessSpeechClient, error) {
	stream, err := c.cc.NewStream(ctx, &SpeechToTextService_ServiceDesc.Streams[0], "/api.v1.speech_to_text.SpeechToTextService/ProcessSpeech", opts...)
	if err != nil {
		return nil, err
	}
	x := &speechToTextServiceProcessSpeechClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SpeechToTextService_ProcessSpeechClient interface {
	Recv() (*ProcessSpeechResponse, error)
	grpc.ClientStream
}

type speechToTextServiceProcessSpeechClient struct {
	grpc.ClientStream
}

func (x *speechToTextServiceProcessSpeechClient) Recv() (*ProcessSpeechResponse, error) {
	m := new(ProcessSpeechResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpeechToTextServiceServer is the server API for SpeechToTextService service.
// All implementations must embed UnimplementedSpeechToTextServiceServer
// for forward compatibility
type SpeechToTextServiceServer interface {
	// Start Processing Audio
	//
	// This is kicked off when the wake word is registered.
	// When speech is no longer detected, the response will return
	// false and the wake word will begin listening again.
	ProcessSpeech(*ProcessSpeechRequest, SpeechToTextService_ProcessSpeechServer) error
	mustEmbedUnimplementedSpeechToTextServiceServer()
}

// UnimplementedSpeechToTextServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSpeechToTextServiceServer struct {
}

func (UnimplementedSpeechToTextServiceServer) ProcessSpeech(*ProcessSpeechRequest, SpeechToTextService_ProcessSpeechServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessSpeech not implemented")
}
func (UnimplementedSpeechToTextServiceServer) mustEmbedUnimplementedSpeechToTextServiceServer() {}

// UnsafeSpeechToTextServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpeechToTextServiceServer will
// result in compilation errors.
type UnsafeSpeechToTextServiceServer interface {
	mustEmbedUnimplementedSpeechToTextServiceServer()
}

func RegisterSpeechToTextServiceServer(s grpc.ServiceRegistrar, srv SpeechToTextServiceServer) {
	s.RegisterService(&SpeechToTextService_ServiceDesc, srv)
}

func _SpeechToTextService_ProcessSpeech_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProcessSpeechRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SpeechToTextServiceServer).ProcessSpeech(m, &speechToTextServiceProcessSpeechServer{stream})
}

type SpeechToTextService_ProcessSpeechServer interface {
	Send(*ProcessSpeechResponse) error
	grpc.ServerStream
}

type speechToTextServiceProcessSpeechServer struct {
	grpc.ServerStream
}

func (x *speechToTextServiceProcessSpeechServer) Send(m *ProcessSpeechResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SpeechToTextService_ServiceDesc is the grpc.ServiceDesc for SpeechToTextService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpeechToTextService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.speech_to_text.SpeechToTextService",
	HandlerType: (*SpeechToTextServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProcessSpeech",
			Handler:       _SpeechToTextService_ProcessSpeech_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v1/speech-to-text.proto",
}