package client

import (
	"fmt"
	"github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/api/v1"
	"google.golang.org/grpc"
)

type ApiClient struct {
	Conn *grpc.ClientConn

	S2T api.SpeechToTextServiceClient
	T2S api.TextToSpeechServiceClient
	GPT api.ChatGptServiceClient
}

func NewApiClient() (*ApiClient, error) {
	conn, err := ApiConn("50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := &ApiClient{
		Conn: conn,
		S2T:  api.NewSpeechToTextServiceClient(conn),
		T2S:  api.NewTextToSpeechServiceClient(conn),
		GPT:  api.NewChatGptServiceClient(conn),
	}
	return client, nil

}

func ApiConn(port string, dailOpts ...grpc.DialOption) (*grpc.ClientConn, error) {
	// TODO Should this be a constant somewhere?
	host := fmt.Sprintf("localhost:%s", port)
	return grpc.Dial(host, dailOpts...)
}
