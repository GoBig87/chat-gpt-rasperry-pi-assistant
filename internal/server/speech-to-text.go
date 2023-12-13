package server

import (
	"github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/api/v1"
	s2t "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/speech-to-text"
	"log"
	"time"
)

func MakeSpeechToTextServer() *SpeechToTextServer {
	return &SpeechToTextServer{}
}

type SpeechToTextServer struct {
	api.UnimplementedSpeechToTextServiceServer
}

func (s *SpeechToTextServer) ProcessSpeech(req *api.ProcessSpeechRequest, stream api.SpeechToTextService_ProcessSpeechServer) error {
	log.Printf("process speech request: %v\n", req)
	// Use a channel to signal when transcription is complete
	transcriptionComplete := make(chan bool)
	var transcribeError error
	var transcribedText string
	// Send initial processing message
	stream.Send(&api.ProcessSpeechResponse{Processing: true, TranscribedText: ""})

	// Run TranscribeSpeech in a goroutine
	go func() {
		transcribedText, transcribeError = s2t.TranscribeSpeech()
		if transcribeError != nil {
			// Handle error, you may want to send an error response to the client
			log.Println("Error in TranscribeSpeech:", transcribeError)
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
			// Transcription completed, send the final processing message
			if transcribeError != nil {
				return transcribeError
			}
			stream.Send(&api.ProcessSpeechResponse{Processing: false, TranscribedText: transcribedText})
			log.Printf("transcribed text finshed: %s\n", transcribedText)
			return nil
		case <-time.After(1 * time.Second):
			// Send processing update every second
			stream.Send(&api.ProcessSpeechResponse{Processing: true, TranscribedText: ""})
		}
	}
}
