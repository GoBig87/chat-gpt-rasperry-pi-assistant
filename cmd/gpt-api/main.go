package main

import (
	"fmt"
	"github.com/GoBig87/chat-gpt-raspberry-pi-assistant/internal/server"
	"github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/api/v1"
	gpio_motor "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/gpio-motor"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
)

var (
	chatGptApiEndpoint string
	chatGptApiKey      string
	chatGptOrgID       string
	motorMouthEna      int
	motorMouthIn1      int
	motorMouthIn2      int
	motorBodyIn3       int
	motorBodyIn4       int
	motorBodyEnb       int
	audioDetect        int
	gpioMotor          *gpio_motor.GpioMotor
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
	chatGptApiEndpoint = os.Getenv("CHAT_GPT_API_ENDPOINT")
	if chatGptApiEndpoint == "" {
		log.Fatal("CHAT_GPT_API_ENDPOINT is not set")
		return
	}
	chatGptApiKey = os.Getenv("CHAT_GPT_API_KEY")
	if chatGptApiKey == "" {
		log.Fatal("CHAT_GPT_API_KEY is not set")
		return
	}
	chatGptOrgID = os.Getenv("CHAT_GPT_ORG_ID")
	if chatGptOrgID == "" {
		log.Fatal("CHAT_GPT_ORG_ID is not set")
		return
	}

	err = godotenv.Load("/var/lib/gpt/gpio.env")
	if err != nil {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
			return
		}
	}
	mouthEnaStr := os.Getenv("MOTOR_MOUTH_ENA")
	motorMouthEna, err = strconv.Atoi(mouthEnaStr)
	if err != nil {
		log.Fatal("MOTOR_MOUTH_ENA is not set")
		return
	}
	mouthIn1Str := os.Getenv("MOTOR_MOUTH_IN1")
	motorMouthIn1, err = strconv.Atoi(mouthIn1Str)
	if err != nil {
		log.Fatal("MOTOR_MOUTH_IN1 is not set")
		return
	}
	mouthIn2Str := os.Getenv("MOTOR_MOUTH_IN2")
	motorMouthIn2, err = strconv.Atoi(mouthIn2Str)
	if err != nil {
		log.Fatal("MOTOR_MOUTH_IN2 is not set")
		return
	}
	bodyIn3Str := os.Getenv("MOTOR_BODY_IN3")
	motorBodyIn3, err = strconv.Atoi(bodyIn3Str)
	if err != nil {
		log.Fatal("MOTOR_BODY_IN3 is not set")
		return
	}
	bodyIn4Str := os.Getenv("MOTOR_BODY_IN4")
	motorBodyIn4, err = strconv.Atoi(bodyIn4Str)
	if err != nil {
		log.Fatal("MOTOR_BODY_IN4 is not set")
		return
	}
	bodyEnbStr := os.Getenv("MOTOR_BODY_ENB")
	motorBodyEnb, err = strconv.Atoi(bodyEnbStr)
	if err != nil {
		log.Fatal("MOTOR_BODY_ENB is not set")
		return
	}
	audioDetectStr := os.Getenv("AUDIO_DETECTOR")
	audioDetect, err = strconv.Atoi(audioDetectStr)
	if err != nil {
		log.Fatal("AUDIO_DETECT is not set")
		return
	}
	gpioMotor, err = gpio_motor.MakeNewGpioMotor(
		motorMouthEna,
		motorMouthIn1,
		motorMouthIn2,
		motorBodyEnb,
		motorBodyIn3,
		motorBodyIn4,
		audioDetect)
	if err != nil {
		log.Fatal("failed to create gpio motor", zap.Error(err))
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
	Use:           "server",
	Short:         "Service to run a grpc server that transcribes speech via a stream",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runGrpcServer()
		if err != nil {
			return err
		}
		return nil
	},
}

func runGrpcServer() error {
	grpcPort := "50051"
	grpcEndpoint := fmt.Sprintf(":%s", grpcPort)
	log.Printf("gRPC endpoint [%s]", grpcEndpoint)

	grpcServer := grpc.NewServer()

	api.RegisterChatGptServiceServer(grpcServer, server.MakeChatGptServer(chatGptApiKey, chatGptOrgID, chatGptApiEndpoint))
	api.RegisterGpioMotorServiceServer(grpcServer, server.MakeGpioMotorServer(gpioMotor))
	api.RegisterSpeechToTextServiceServer(grpcServer, server.MakeSpeechToTextServer())
	api.RegisterTextToSpeechServiceServer(grpcServer, server.MakeTextToSpeechServer())

	lis, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	if err = grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
