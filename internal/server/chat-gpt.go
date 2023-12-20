package server

import (
	"context"
	"fmt"
	"github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/api/v1"
	gpt "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/chat-gpt"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func MakeChatGptServer(apiKey, orgId, endpoint, systemPrompt string) *ChatGptServer {
	sp := gpt.RequestMessage{
		Role:    "system",
		Content: systemPrompt,
	}
	return &ChatGptServer{
		client:       gpt.NewChatGptClient(apiKey, orgId, endpoint),
		systemPrompt: sp,
	}
}

type ChatGptServer struct {
	client       *gpt.ChatGptClient
	systemPrompt gpt.RequestMessage
	prevPrompts  []gpt.RequestMessage
	api.UnimplementedChatGptServiceServer
}

func (s *ChatGptServer) ProcessPrompt(ctx context.Context, req *api.ProcessPromptRequest) (*api.ProcessPromptResponse, error) {
	log.Printf("Received prompt: %s\n", req.Prompt)
	prompt := req.Prompt

	var messages []gpt.RequestMessage
	messages = append(messages, s.systemPrompt)
	messages = append(messages, s.prevPrompts...)

	question := gpt.RequestMessage{
		Role:      "user",
		Content:   prompt,
		CreatedAt: time.Now(),
	}

	messages = append(messages, question)

	resp, err := s.client.PromptChatGPT(messages)
	if err != nil {
		log.Print("Error in ProcessPrompt", zap.Error(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error processing prompt: %v", err))
	}
	log.Printf("Response finished: %s\n", resp)
	answer := gpt.RequestMessage{
		Role:      "assistant",
		Content:   resp,
		CreatedAt: time.Now(),
	}
	s.prevPrompts = append(s.prevPrompts, []gpt.RequestMessage{question, answer}...)

	twoHoursAgo := time.Now().Add(-2 * time.Hour)
	var filteredPrompts []gpt.RequestMessage
	for _, message := range s.prevPrompts {
		// Check if the message is within the last two hours
		if message.CreatedAt.After(twoHoursAgo) {
			// Keep the message in the filtered slice
			filteredPrompts = append(filteredPrompts, message)
		}
	}
	s.prevPrompts = filteredPrompts

	return &api.ProcessPromptResponse{Response: resp}, nil
}
