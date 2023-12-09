package main

import (
	"context"
	"fmt"
	porcupine "github.com/Picovoice/porcupine/binding/go/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"os"

	api_client "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/api/client"
	v1 "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/api/v1"
	ww "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/wake-word"
)

var (
	accessKey string
	client    *api_client.ApiClient
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	accessKey = os.Getenv("PORCUPINE_ACCESS_KEY")
	if accessKey == "" {
		log.Fatal("PORCUPINE_ACCESS_KEY is not set")
		return
	}
	modelPath := os.Getenv("PORCUPINE_MODEL_PATH")
	if modelPath == "" {
		log.Fatal("PORCUPINE_MODEL_PATH is not set")
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
	Short:         "Main application to handle speech-to-text and chat-gpt",
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
	prompt, err := processSpeechToText(ctx)
	if err != nil {
		return err
	}
	answer, err := processChatGptPrompt(ctx, prompt)
	if err != nil {
		return err
	}
	return processChatGptResponse(ctx, answer)
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
			log.Fatalf("error receiving speech: %v", err)
			return "", err
		}
		if !resp.Processing {
			text = resp.TranscribedText
			processed = true
		}
	}
	if text == "" {
		log.Fatalf("text is empty")
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
		log.Fatalf("response is empty")
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
			log.Fatalf("error receiving response: %v", err)
			return err
		}
		if resp.Processed {
			processed = true
		}
	}
	log.Println("finished processing response")
	return nil
}
