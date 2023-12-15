package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"
	"os"
	"strconv"
	"time"

	gpio_motor "github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/gpio-motor"
)

var gpioCmd = &cobra.Command{
	Use:   "gpio",
	Short: "gpio management commands",
}

var (
	motorMouthEna int
	motorMouthIn1 int
	motorMouthIn2 int
	motorBodyIn3  int
	motorBodyIn4  int
	motorBodyEnb  int
	audioDetect   int
	gpioMotor     *gpio_motor.GpioMotor
)

func init() {
	err := godotenv.Load("/var/lib/gpt/gpio.env")
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

	gpioCmd.AddCommand(closeMouthCmd)
	gpioCmd.AddCommand(lowerHeadCmd)
	gpioCmd.AddCommand(lowerTailCmd)
	gpioCmd.AddCommand(openMouthCmd)
	gpioCmd.AddCommand(raiseHeadCmd)
	gpioCmd.AddCommand(raiseTailCmd)
	gpioCmd.AddCommand(resetAllCmd)
	gpioCmd.AddCommand(testAudioDetectCmd)
}

var closeMouthCmd = &cobra.Command{
	Use:   "close-mouth",
	Short: "closes billy bass's mouth",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := gpioMotor.CloseMouth()
		if err != nil {
			return err
		}
		return nil
	},
}

var lowerHeadCmd = &cobra.Command{
	Use:   "lower-head",
	Short: "lowers billy bass's head",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := gpioMotor.LowerHead()
		if err != nil {
			return err
		}
		return nil
	},
}

var lowerTailCmd = &cobra.Command{
	Use:   "lower-tail",
	Short: "lowers billy bass's tail",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := gpioMotor.LowerTail()
		if err != nil {
			return err
		}
		return nil
	},
}

var openMouthCmd = &cobra.Command{
	Use:   "open-mouth",
	Short: "opens billy bass's mouth",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := gpioMotor.OpenMouth()
		if err != nil {
			return err
		}
		return nil
	},
}

var raiseHeadCmd = &cobra.Command{
	Use:   "raise-head",
	Short: "raises billy bass's head",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := gpioMotor.RaiseHead()
		if err != nil {
			return err
		}
		return nil
	},
}

var raiseTailCmd = &cobra.Command{
	Use:   "raise-tail",
	Short: "raises billy bass's tail",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := gpioMotor.RaiseTail()
		if err != nil {
			return err
		}
		return nil
	},
}

var resetAllCmd = &cobra.Command{
	Use:   "reset-all",
	Short: "resets all billy bass motors",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := gpioMotor.ResetAll()
		if err != nil {
			return err
		}
		return nil
	},
}

var testAudioDetectCmd = &cobra.Command{
	Use:   "test-audio-detect <duration (seconds) optional>",
	Short: "tests audio detection",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		duration := 0
		duration, err := strconv.Atoi(args[0])
		if err != nil {
			log.Print("duration is not set, looping continuously")
		}
		if duration > 0 {
			intitialTime := time.Now().Unix()
			for {
				detected, err := gpioMotor.IsAudioDetected()
				if err != nil {
					return err
				} else {
					log.Printf("audio detection status: %v", detected)
				}
				currentTime := time.Now().Unix()
				diff := int(currentTime) - int(intitialTime)
				if diff >= duration {
					log.Print("duration reached, exiting")
					break
				}
			}

		} else {
			for {
				detected, err := gpioMotor.IsAudioDetected()
				if err != nil {
					return err
				} else {
					log.Printf("audio detection status: %v", detected)
				}
			}
		}
		return nil
	},
}
