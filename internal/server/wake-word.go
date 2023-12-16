package server

import (
	"github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/api/v1"
	ww "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/wake-word"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"sync"
	"time"
)

func MakeWakeWordServer(accessKey string) *WakeWordServer {
	return &WakeWordServer{
		accessKey: accessKey,
	}
}

type WakeWordServer struct {
	accessKey string
	api.UnimplementedWakeWordServiceServer
}

func (s *WakeWordServer) DetectWakeWord(req *emptypb.Empty, stream api.WakeWordService_DetectWakeWordServer) error {
	stopCh := make(chan struct{})
	wakeWordDetectedCh := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)

	// Run DetectWakeWordRoutine in a Goroutine
	go func() {
		defer wg.Done()

		wakeWord, err := ww.DetectWakeWordRoutine(s.accessKey, stopCh)
		if err != nil {
			log.Printf("Error in DetectWakeWordRoutine: %v", err)
		}
		log.Printf("Wake word %s detected!", string(wakeWord))
		wakeWordDetectedCh <- string(wakeWord)
	}()

	for {
		// check to see if a wake word was detected
		// Check if a wake word was detected
		select {
		case detectedKeyword := <-wakeWordDetectedCh:
			log.Printf("Wake word %s detected!", detectedKeyword)
			// Handle the wake word detection as needed
			resp := &api.WakeWordResponse{
				BuiltInKeyword: detectedKeyword,
				CustomKeyword:  "",
				Detected:       true,
			}
			if err := stream.Send(resp); err != nil {
				log.Printf("Error sending built in keyword stream info: %v", err)
				return status.Errorf(codes.Internal, "Error sending stream info: %v", err)
			}
			close(stopCh)
			wg.Wait()
			return nil
		default:
			// No wake word detected yet, continue processing other data
			resp := &api.WakeWordResponse{
				BuiltInKeyword: "",
				CustomKeyword:  "",
				Detected:       false,
			}
			if err := stream.Send(resp); err != nil {
				log.Printf("Error sending default stream info: %v", err)
				close(stopCh)
				wg.Wait()
				return status.Errorf(codes.Internal, "Error sending stream info: %v", err)
			}
		}
		time.Sleep(250 * time.Millisecond)
	}
}