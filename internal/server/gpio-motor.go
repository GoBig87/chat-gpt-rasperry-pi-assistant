package server

import (
	"context"
	"github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/api/v1"
	gpio_motor "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/gpio-motor"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func MakeGpioMotorServer(gpioMotor *gpio_motor.GpioMotor) *GpioMotorServer {
	return &GpioMotorServer{
		gpioMotor: gpioMotor,
	}
}

type GpioMotorServer struct {
	gpioMotor *gpio_motor.GpioMotor
	api.UnimplementedGpioMotorServiceServer
}

// CloseMouth implements the CloseMouth RPC.
func (s *GpioMotorServer) CloseMouth(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.gpioMotor.CloseMouth()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to close mouth: %v", zap.Error(err))
	}
	return &emptypb.Empty{}, nil
}

// LowerHead implements the LowerHead RPC.
func (s *GpioMotorServer) LowerHead(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.gpioMotor.LowerHead()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to lower head: %v", zap.Error(err))
	}
	return &emptypb.Empty{}, nil
}

// LowerTail implements the LowerTail RPC.
func (s *GpioMotorServer) LowerTail(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.gpioMotor.LowerTail()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to lower tail: %v", zap.Error(err))
	}
	return &emptypb.Empty{}, nil
}

// OpenMouth implements the OpenMouth RPC.
func (s *GpioMotorServer) OpenMouth(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.gpioMotor.OpenMouth()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to open mouth: %v", zap.Error(err))
	}
	return &emptypb.Empty{}, nil
}

// RaiseHead implements the RaiseHead RPC.
func (s *GpioMotorServer) RaiseHead(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.gpioMotor.RaiseHead()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to raise head: %v", zap.Error(err))
	}
	return &emptypb.Empty{}, nil
}

// RaiseTail implements the RaiseTail RPC.
func (s *GpioMotorServer) RaiseTail(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.gpioMotor.RaiseTail()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to raise tail: %v", zap.Error(err))
	}
	return &emptypb.Empty{}, nil
}

// ResetAll implements the ResetAll RPC.
func (s *GpioMotorServer) ResetAll(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.gpioMotor.ResetAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to reset all: %v", zap.Error(err))
	}
	return &emptypb.Empty{}, nil
}
