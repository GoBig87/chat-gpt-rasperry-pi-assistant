package main

import (
	"fmt"
	"github.com/GoBig87/chat-gpt-raspberry-pi-assistant/internal/server"
	"github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/api/v1"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func init() {
	// TODO add env vars here
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:           "server",
	Short:         "Service to run a grpc server that transcribes speech via a stream",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runGrpcServer()
		if err != nil {
			return err
		}
		return nil
	},
}

func runGrpcServer() error {
	grpcPort := "50051"
	grpcEndpoint := fmt.Sprintf(":%s", grpcPort)
	log.Printf("gRPC endpoint [%s]", grpcEndpoint)

	grpcServer := grpc.NewServer()

	api.RegisterSpeechToTextServiceServer(grpcServer, server.MakeSpeechToTextServer())
	api.RegisterTextToSpeechServiceServer(grpcServer, server.MakeTextToSpeechServer())
	api.RegisterChatGptServiceServer(grpcServer, server.MakeChatGptServer())

	lis, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	if err = grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
