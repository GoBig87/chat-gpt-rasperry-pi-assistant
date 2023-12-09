package server

import (
	"context"
	"fmt"
	"github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/api/v1"
	gpt "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/chat-gpt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func MakeChatGptServer() *ChatGptServer {
	return &ChatGptServer{}
}

type ChatGptServer struct {
	api.UnimplementedChatGptServiceServer
}

func (s *ChatGptServer) ProcessPrompt(ctx context.Context, req *api.ProcessPromptRequest) (*api.ProcessPromptResponse, error) {
	log.Printf("Received prompt: %s\n", req.Prompt)
	prompt := req.Prompt
	resp, err := gpt.PromptChatGPT(prompt)
	if err != nil {
		log.Fatalf("Error in ProcessPrompt: %v", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error processing prompt: %v", err))
	}
	log.Printf("Response finished: %s\n", resp)
	return &api.ProcessPromptResponse{Response: resp}, nil
}
