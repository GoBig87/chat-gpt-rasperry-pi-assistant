package main

import (
	"context"
	"fmt"
	porcupine "github.com/Picovoice/porcupine/binding/go/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"os"
	"time"

	api_client "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/api/client"
	v1 "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/api/v1"
	ww "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/wake-word"
)

var (
	accessKey string
	client    *api_client.ApiClient
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
	accessKey = os.Getenv("PORCUPINE_ACCESS_KEY")
	if accessKey == "" {
		log.Fatal("PORCUPINE_ACCESS_KEY is not set")
		return
	}
	client, err = api_client.NewApiClient()
	if err != nil {
		log.Fatal("Error creating api client")
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
	Use:           "app",
	Short:         "Main application to handle speech-to-text, chat-gpt and text-to-speech",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		run(ctx)
		return nil
	},
}

func run(ctx context.Context) {
	for {
		// 1. Wake word
		log.Println("Listening for wake word...")
		keyword, err := ww.DetectWakeWord(accessKey)
		if err != nil {
			log.Printf("Error detecting wake word: %v", err)
			continue
		}
		switch keyword {
		case porcupine.ALEXA:
			log.Printf("Unimplemnted wake word: %v", keyword)
			continue
		case porcupine.AMERICANO:
			log.Printf("Unimplemnted wake word: %v", keyword)
			continue
		case porcupine.BLUEBERRY:
			log.Printf("Unimplemnted wake word: %v", keyword)
			continue
		case porcupine.BUMBLEBEE:
			err := process(ctx)
			if err != nil {
				fmt.Printf("Error processing: %v", err)
			}
			continue
		case porcupine.COMPUTER:
			log.Printf("Unimplemnted wake word: %v", keyword)
			continue
		case porcupine.GRAPEFRUIT:
			log.Printf("Unimplemnted wake word: %v", keyword)
			continue
		case porcupine.GRASSHOPPER:
			log.Printf("Unimplemnted wake word: %v", keyword)
			continue
		case porcupine.HEY_GOOGLE:
			err := process(ctx)
			if err != nil {
				fmt.Printf("Error processing: %v", err)
			}
			continue
		case porcupine.HEY_SIRI:
			log.Printf("Unimplemnted wake word: %v", keyword)
			continue
		case porcupine.JARVIS:
			log.Printf("Unimplemnted wake word: %v", keyword)
			continue
		case porcupine.OK_GOOGLE:
			err := process(ctx)
			if err != nil {
				fmt.Printf("Error processing: %v", err)
			}
			continue
		case porcupine.PICOVOICE:
			log.Printf("Unimplemnted wake word: %v", keyword)
			continue
		case porcupine.PORCUPINE:
			log.Printf("Unimplemnted wake word: %v", keyword)
			continue
		case porcupine.TERMINATOR:
			log.Printf("Unimplemnted wake word: %v", keyword)
			continue
		default:
			log.Printf("Unknown wake word: %v", keyword)
			continue
		}
	}
}

func process(ctx context.Context) error {
	// raise head to acknowledge wake word
	_, err := client.MTR.RaiseHead(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}
	defer client.MTR.ResetAll(ctx, &emptypb.Empty{})

	prompt, err := processSpeechToText(ctx)
	if err != nil {
		return err
	}

	done1 := make(chan struct{})
	defer close(done1)
	// Start a goroutine to raise and lower the tail
	go wagTail(ctx, done1)

	answer, err := processChatGptPrompt(ctx, prompt)
	if err != nil {
		return err
	}
	// Stop the goroutine
	done1 <- struct{}{}

	done2 := make(chan struct{})
	defer close(done2)
	// Start a goroutine to move mouth
	go moveMouth(ctx, done2)

	err = processChatGptResponse(ctx, answer)
	if err != nil {
		return err
	}
	// Stop the goroutine
	done2 <- struct{}{}
	return err
}

func processSpeechToText(ctx context.Context) (string, error) {
	log.Println("Processing speech to text")
	// 2. Speech-to-text
	req := &v1.ProcessSpeechRequest{
		// TODO UNHARDCODE
		WakeWord: v1.WakeWord_WAKE_WORD_HEY_GOOGLE,
	}
	clnt, err := client.S2T.ProcessSpeech(ctx, req)
	if err != nil {
		log.Printf("error processing speech: %v", err)
		return "", err
	}
	processed := false
	text := ""
	for {
		if processed {
			break
		}
		resp, err := clnt.Recv()
		if err != nil {
			log.Printf("error receiving speech: %v", err)
			return "", err
		}
		if !resp.Processing {
			text = resp.TranscribedText
			processed = true
		}
	}
	if text == "" {
		log.Printf("text is empty")
		return "", fmt.Errorf("text is empty")
	}
	log.Println("finished processing wake word")
	return text, nil
}

func processChatGptPrompt(ctx context.Context, prompt string) (string, error) {
	log.Printf("Processing Prompt: %v\n", prompt)
	// 3. Chat GPT Response
	req := &v1.ProcessPromptRequest{
		Prompt: prompt,
	}
	resp, err := client.GPT.ProcessPrompt(ctx, req)
	if err != nil {
		log.Printf("error processing prompt: %v", err)
		return "", err
	}
	answer := resp.Response
	if answer == "" {
		log.Printf("response is empty")
		return "", fmt.Errorf("response is empty")
	}
	log.Println("finished processing prompt")
	return answer, nil
}

func processChatGptResponse(ctx context.Context, response string) error {
	log.Printf("Processing Response: %v\n", response)
	// 4. Text-to-speech
	req := &v1.ProcessTextRequest{
		Text: response,
	}
	clnt, err := client.T2S.ProcessText(ctx, req)
	if err != nil {
		log.Printf("error processing response: %v", err)
		return err
	}
	processed := false
	for {
		if processed {
			break
		}
		resp, err := clnt.Recv()
		if err != nil {
			log.Printf("error receiving response: %v", err)
			return err
		}
		if resp.Processed {
			processed = true
		}
	}
	log.Println("finished processing response")
	return nil
}

func wagTail(ctx context.Context, done chan struct{}) {
	// TODO move this to a GRPC stream
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	_, err := client.MTR.ResetAll(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("Error lowering head: %v", err)
	}
	// Sleep for half a second
	time.Sleep(500 * time.Millisecond)
	raised := false
	for {
		select {
		case <-done:
			if _, err := client.MTR.ResetAll(ctx, &emptypb.Empty{}); err != nil {
				log.Printf("Error reseting: %v", err)
			}
			return
		case <-ticker.C:
			if !raised {
				// Raise the tail
				if _, err := client.MTR.RaiseTail(ctx, &emptypb.Empty{}); err != nil {
					log.Printf("Error raising tail: %v", err)
				} else {
					raised = true
				}
			} else {
				// Lower the tail
				if _, err := client.MTR.LowerTail(ctx, &emptypb.Empty{}); err != nil {
					log.Printf("Error lowering tail: %v", err)
				} else {
					raised = false
				}
			}
		}
	}
}

func moveMouth(ctx context.Context, done chan struct{}) {
	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()
	if _, err := client.MTR.LowerTail(ctx, &emptypb.Empty{}); err != nil {
		log.Printf("Error lowering tail: %v", err)
	}
	time.Sleep(500 * time.Millisecond)
	_, err := client.MTR.RaiseHead(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("Error lowering head: %v", err)
	}
	stream, err := client.MTR.MoveMouthToSpeech(ctx)
	if err != nil {
		log.Printf("Error creating moving mouth to speech stream: %v", err)
		return
	}

	for {
		select {
		case <-done:
			req := &v1.MoveMouthToSpeechRequest{
				Stop: true,
			}
			err = stream.Send(req)
			if err != nil {
				log.Printf("Error sending stop to stream: %v", err)
				return
			}
			_, err := stream.CloseAndRecv()
			if err != nil {
				log.Printf("Error closing stream: %v", err)
				return
			}
			if _, err := client.MTR.ResetAll(ctx, &emptypb.Empty{}); err != nil {
				log.Printf("Error reseting: %v", err)
			}
			return
		case <-ticker.C:
			req := &v1.MoveMouthToSpeechRequest{
				Stop: false,
			}
			err = stream.Send(req)
			if err != nil {
				log.Printf("Error sending stop to stream: %v", err)
				return
			}
		}
	}
}
