package main

import (
	"fmt"
	"github.com/GoBig87/chat-gpt-raspberry-pi-assistant/internal/server"
	"github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/api/v1"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var (
	chatGptApiEndpoint string
	chatGptApiKey      string
	chatGptOrgID       string
)

func init() {
	err := godotenv.Load("/var/lib/gpt/config.env")
	if err != nil {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
			return
		}
	}
	chatGptApiEndpoint = os.Getenv("CHAT_GPT_API_ENDPOINT")
	if chatGptApiEndpoint == "" {
		log.Fatal("CHAT_GPT_API_ENDPOINT is not set")
		return
	}
	chatGptApiKey = os.Getenv("CHAT_GPT_API_KEY")
	if chatGptApiKey == "" {
		log.Fatal("CHAT_GPT_API_KEY is not set")
		return
	}
	chatGptOrgID = os.Getenv("CHAT_GPT_ORG_ID")
	if chatGptOrgID == "" {
		log.Fatal("CHAT_GPT_ORG_ID is not set")
		return
	}
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
	api.RegisterChatGptServiceServer(grpcServer, server.MakeChatGptServer(chatGptApiKey, chatGptOrgID, chatGptApiEndpoint))

	lis, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	if err = grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
