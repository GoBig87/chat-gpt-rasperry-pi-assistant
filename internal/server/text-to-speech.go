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
	var transcribeError error
	var transcribedText string
	// Send initial processing message
	stream.Send(&api.ProcessTextResponse{Processed: false})
	text := req.Text

	// Run TranscribeText in a goroutine
	go func() {
		transcribeError = t2s.TranscribeText(text)
		if transcribeError != nil {
			fmt.Println("Error in TranscribeSpeech:", transcribeError)
			transcriptionComplete <- true
			return
		}
		log.Println(transcribedText)
		transcriptionComplete <- true
	}()

	// Periodically send processing updates
	for {
		select {
		case <-transcriptionComplete:
			stream.Send(&api.ProcessTextResponse{Processed: true})
			log.Printf("Text to speech complete\n")
			return nil
		case <-time.After(1 * time.Second):
			// Send processing update every second
			stream.Send(&api.ProcessTextResponse{Processed: false})
		}
	}
}
