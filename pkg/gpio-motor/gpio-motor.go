package gpio_motor

import (
	"github.com/GoBig87/chat-gpt-raspberry-pi-assistant/pkg/utils"
	"github.com/stianeikeland/go-rpio"
	"log"
	"time"
)

type GpioMotor struct {
	HeadEnable    int
	HeadIN1       int
	HeadIN2       int
	BodyEnable    int
	BodyIN1       int
	BodyIN2       int
	AudioDetected int
}

func MakeNewGpioMotor(headEnable, headIn1, headIn2, bodyEnable, bodyIn1, bodyIn2, audio int) (*GpioMotor, error) {
	bcmHeadEnable, err := utils.PhysicalPinToBCM(headEnable)
	if err != nil {
		return nil, err
	}
	bcmHeadIn1, err := utils.PhysicalPinToBCM(headIn1)
	if err != nil {
		return nil, err
	}
	bcmHeadIn2, err := utils.PhysicalPinToBCM(headIn2)
	if err != nil {
		return nil, err
	}
	bcmBodyEnable, err := utils.PhysicalPinToBCM(bodyEnable)
	if err != nil {
		return nil, err
	}
	bcmBodyIn1, err := utils.PhysicalPinToBCM(bodyIn1)
	if err != nil {
		return nil, err
	}
	bcmBodyIn2, err := utils.PhysicalPinToBCM(bodyIn2)
	if err != nil {
		return nil, err
	}
	audioDetect, err := utils.PhysicalPinToBCM(audio)
	if err != nil {
		return nil, err
	}
	return &GpioMotor{
		HeadEnable:    bcmHeadEnable,
		HeadIN1:       bcmHeadIn1,
		HeadIN2:       bcmHeadIn2,
		BodyEnable:    bcmBodyEnable,
		BodyIN1:       bcmBodyIn1,
		BodyIN2:       bcmBodyIn2,
		AudioDetected: audioDetect,
	}, nil
}

func (g *GpioMotor) CloseMouth() error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close()
	enableHeadPin := rpio.Pin(g.HeadEnable)
	in3Pin := rpio.Pin(g.HeadIN1)
	in4Pin := rpio.Pin(g.HeadIN2)

	enableHeadPin.Output()
	in3Pin.Output()
	in4Pin.Output()

	enableHeadPin.Low()
	in3Pin.Low()
	in4Pin.Low()

	return nil
}

func (g *GpioMotor) IsAudioDetected() (bool, error) {
	err := rpio.Open()
	if err != nil {
		return false, err
	}
	defer rpio.Close()
	audioPin := rpio.Pin(g.AudioDetected)
	audioPin.Input()
	return audioPin.Read() == rpio.Low, nil
}

func (g *GpioMotor) LowerHead() error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close()

	enablePin := rpio.Pin(g.BodyEnable)
	in1Pin := rpio.Pin(g.BodyIN1)
	in2Pin := rpio.Pin(g.BodyIN2)

	enablePin.Output()
	in1Pin.Output()
	in2Pin.Output()

	// Enable the motor
	in1Pin.Low()
	in2Pin.Low()
	enablePin.Low()

	return nil
}

func (g *GpioMotor) LowerTail() error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close()

	enablePin := rpio.Pin(g.BodyEnable)
	in1Pin := rpio.Pin(g.BodyIN1)
	in2Pin := rpio.Pin(g.BodyIN2)

	enablePin.Output()
	in1Pin.Output()
	in2Pin.Output()

	// Set the direction
	in1Pin.Low()
	in2Pin.Low()
	enablePin.Low()

	return nil
}

func (g *GpioMotor) MoveMouthToSpeech() error {
	// Run for 3 iterations, 30 ms to detect if between words
	silenceCount := 0
	for i := 0; i < 5; i++ {
		detected, err := g.IsAudioDetected()
		if err != nil {
			return err
		}
		if !detected {
			log.Printf("Silence detected incrementing silence count %d", silenceCount)
			silenceCount++
		}
		time.Sleep(10 * time.Millisecond)
	}

	var err error
	if silenceCount >= 4 {
		log.Print("Closing mouth")
		err = g.CloseMouth()
	} else {
		log.Print("Opening mouth")
		err = g.OpenMouth()
	}
	return err
}

func (g *GpioMotor) OpenMouth() error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close()

	enableHeadPin := rpio.Pin(g.HeadEnable)
	in3Pin := rpio.Pin(g.HeadIN1)
	in4Pin := rpio.Pin(g.HeadIN2)

	enableHeadPin.Output()
	in3Pin.Output()
	in4Pin.Output()

	// Enable the motor
	enableHeadPin.High()

	// Set the direction "forward"
	in3Pin.High()
	in4Pin.Low()

	return nil
}

func (g *GpioMotor) RaiseHead() error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close()

	enablePin := rpio.Pin(g.BodyEnable)
	in1Pin := rpio.Pin(g.BodyIN1)
	in2Pin := rpio.Pin(g.BodyIN2)

	enablePin.Output()
	in1Pin.Output()
	in2Pin.Output()

	// Enable the motor
	enablePin.High()

	// Set the direction "forward"
	in1Pin.Low()
	in2Pin.High()

	return nil
}

func (g *GpioMotor) RaiseTail() error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close()

	enablePin := rpio.Pin(g.BodyEnable)
	in1Pin := rpio.Pin(g.BodyIN1)
	in2Pin := rpio.Pin(g.BodyIN2)

	enablePin.Output()
	in1Pin.Output()
	in2Pin.Output()

	// Enable the motor
	enablePin.High()

	// Set the direction "reverse"
	in1Pin.High()
	in2Pin.Low()

	return nil
}

func (g *GpioMotor) ResetAll() error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close()

	enableBodyPin := rpio.Pin(g.BodyEnable)
	in1Pin := rpio.Pin(g.BodyIN1)
	in2Pin := rpio.Pin(g.BodyIN2)

	enableHeadPin := rpio.Pin(g.HeadEnable)
	in3Pin := rpio.Pin(g.HeadIN1)
	in4Pin := rpio.Pin(g.HeadIN2)

	enableBodyPin.Output()
	in1Pin.Output()
	in2Pin.Output()

	enableHeadPin.Output()
	in3Pin.Output()
	in4Pin.Output()

	enableBodyPin.Low()
	in1Pin.Low()
	in2Pin.Low()

	enableHeadPin.Low()
	in3Pin.Low()
	in4Pin.Low()

	return nil
}
