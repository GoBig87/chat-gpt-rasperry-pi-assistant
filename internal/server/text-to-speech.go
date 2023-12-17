package server

import (
	"fmt"
	"github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/api/v1"
	t2s "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/text-to-speech"
	"log"
	"time"
)

func MakeTextToSpeechServer() *TextToSpeechServer {
	return &TextToSpeechServer{}
}

type TextToSpeechServer struct {
	api.UnimplementedTextToSpeechServiceServer
}

func (t *TextToSpeechServer) ProcessText(req *api.ProcessTextRequest, stream api.TextToSpeechService_ProcessTextServer) error {
	log.Printf("process text request: %v\n", req)
	// Use a channel to signal when transcription is complete
	transcriptionComplete := make(chan bool)
	defer close(transcriptionComplete)
	transcriptionError := make(chan error)
	defer close(transcriptionError)
	// Send initial processing message
	stream.Send(&api.ProcessTextResponse{Processed: false})
	text := req.Text

	// Run TranscribeText in a goroutine
	go func() {
		err := t2s.TranscribeText(text)
		if err != nil {
			fmt.Println("Error in TranscribeSpeech:", err)
			transcriptionError <- err
			return
		}
		transcriptionComplete <- true
	}()

	// Periodically send processing updates
	for {
		select {
		case <-transcriptionComplete:
			stream.Send(&api.ProcessTextResponse{Processed: true})
			log.Printf("Text to speech complete\n")
			return nil
		case err := <-transcriptionError:
			log.Printf("Error transcribing text: %v\n", err)
			return err
		case <-time.After(1 * time.Second):
			// Send processing update every second
			stream.Send(&api.ProcessTextResponse{Processed: false})
		}
	}
}
