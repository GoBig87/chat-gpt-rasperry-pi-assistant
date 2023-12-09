package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"os"

	gpt "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/chat-gpt"
	s2t "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/speech-to-text"
	t2s "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/text-to-speech"
	ww "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/wake-word"
)

var accessKey string

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

	rootCmd.InitDefaultHelpCmd()
	walk(rootCmd, func(c *cobra.Command) {
		if c.Name() == "help" {
			c.Short = "help about any command"
		}
	})

	rootCmd.AddCommand(s2tCmd)
	rootCmd.AddCommand(t2sCmd)
	rootCmd.AddCommand(chatGptCmd)
	rootCmd.AddCommand(wakeCmd)
}

func main() {
	errOutput := rootCmd.OutOrStderr()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(errOutput, "Error: %v\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:           "gptctl",
	Short:         "control commands",
	SilenceErrors: true,
	SilenceUsage:  true,
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var s2tCmd = &cobra.Command{
	Use:           "s2t",
	Short:         "Process speech from microphone and into to text",
	SilenceErrors: false,
	SilenceUsage:  false,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := s2t.TranscribeSpeech()
		if err != nil {
			fmt.Printf("Error: %v\n\n", err)
			return err
		}
		fmt.Printf("resp: %v\n", resp)
		return nil
	},
}

var t2sCmd = &cobra.Command{
	Use:           "t2s <text to transcribe to audio>",
	Short:         "Process text and turns into audio",
	SilenceErrors: false,
	SilenceUsage:  false,
	Args:          cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		text := args[0]
		err := t2s.TranscribeText(text)
		if err != nil {
			fmt.Printf("Error: %v\n\n", err)
			return err
		}
		return nil
	},
}

var chatGptCmd = &cobra.Command{
	Use:           "prompt <prompt for chat gpt>",
	Short:         "Sends a question to chat gpt",
	SilenceErrors: false,
	SilenceUsage:  false,
	Args:          cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		text := args[0]
		resp, err := gpt.PromptChatGPT(text)
		if err != nil {
			fmt.Printf("Error: %v\n\n", err)
			return err
		}
		fmt.Println(resp)
		return nil
	},
}

var wakeCmd = &cobra.Command{
	Use:           "wake",
	Short:         "Starts the wake word detection",
	SilenceErrors: false,
	SilenceUsage:  false,
	Args:          cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		keyword, err := ww.DetectWakeWord(accessKey)
		if err != nil {
			fmt.Printf("Error: %v\n\n", err)
			return err
		}
		fmt.Println(keyword)
		return nil
	},
}

func walk(c *cobra.Command, f func(*cobra.Command)) {
	f(c)
	for _, c := range c.Commands() {
		walk(c, f)
	}
}
