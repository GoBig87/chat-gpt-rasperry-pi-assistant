package speech_to_text

import (
	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/gen2brain/malgo"
	"go.uber.org/zap"
	"io"
	"log"
	"time"
)

const (
	sampleRate       = 16000
	bytesPerSample   = 2 // For 16-bit PCM (LINEAR16)
	secondsPerBuffer = 1
)

func TranscribeSpeech() (string, error) {
	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	stream, err := client.StreamingRecognize(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Send the initial configuration message.
	if err := stream.Send(&speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					Encoding:        speechpb.RecognitionConfig_LINEAR16,
					SampleRateHertz: 16000,
					LanguageCode:    "en-US",
				},
			},
		},
	}); err != nil {
		log.Fatal(err)
	}

	backends := []malgo.Backend{malgo.BackendAlsa}
	context, err := malgo.InitContext(backends, malgo.ContextConfig{}, func(message string) {
		//log.Print(fmt.Sprintf("%v\n", message))
	})
	if err != nil {
		log.Print("malgo failed to init", zap.Error(err))
	}
	defer func() {
		_ = context.Uninit()
		context.Free()
	}()

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Duplex)
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = 1
	deviceConfig.SampleRate = 16000

	// Calculate the size of the buffer based on the sample rate and duration
	frameSize := (sampleRate * bytesPerSample * secondsPerBuffer) / bytesPerSample

	// Create a buffer to hold audio data
	audioBuffer := make([]int16, frameSize)

	var streamStopped bool
	// Function to process received frames
	onRecvFrames := func(pSample2, pSample []byte, framecount uint32) {
		if !streamStopped {
			for i := 0; i < len(pSample); i += 2 {
				audioBuffer = append(audioBuffer, int16(binary.LittleEndian.Uint16(pSample[i:i+2])))
				if len(audioBuffer) >= frameSize {
					// Process the audio buffer and send it to the gRPC server
					err = stream.Send(&speechpb.StreamingRecognizeRequest{
						StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
							AudioContent: int16SliceToBytes(audioBuffer),
						},
					})
					if err != nil {
						if err.Error() == "rpc error: code = Internal desc = SendMsg called after CloseSend" {
							fmt.Println("Stream closed, stopping device")
							streamStopped = true
						}
						log.Printf("Could not send audio: %v", err) // I want to break from the loop here
					} else {
						//log.Printf("Sent %d bytes of audio to the API.\n", len(audioBuffer))
					}
					// Clear the audio buffer for the next frame
					audioBuffer = audioBuffer[:0]
				}
			}
		}
	}

	captureCallbacks := malgo.DeviceCallbacks{
		Data: onRecvFrames,
	}
	device, err := malgo.InitDevice(context.Context, deviceConfig, captureCallbacks)
	if err != nil {
		log.Print("failed to init malgo", zap.Error(err))
	}

	go func() {
		err = device.Start()
		if err != nil {
			log.Print("failed to start device", zap.Error(err))
		}
	}()

	responseReceived := make(chan time.Time)
	go func() {
		for {
			select {
			case <-responseReceived:
				continue
			case <-time.After(6 * time.Second):
				fmt.Println("Timeout reached, closing stream.")
				err = stream.CloseSend()
				if err != nil {
					log.Print("Could not close stream", zap.Error(err))
				} else {
					streamStopped = true
				}
				return
			}
		}
	}()

	var ret string
	for {
		responseReceived <- time.Now()
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Print("Cannot stream results", zap.Error(err))
		}
		if err := resp.Error; err != nil {
			// Workaround while the API doesn't give a more informative error.
			if err.Code == 3 || err.Code == 11 {
				log.Print("WARNING: Speech recognition request exceeded limit of 60 seconds.")
			}
		}
		tempRet := ""
		for _, result := range resp.Results {
			fmt.Printf("Result: %+v\n", result)
			tempRet = tempRet + result.Alternatives[0].Transcript
		}
		if tempRet == "" {
			break
		}
		ret = ret + tempRet
		if streamStopped {
			break
		}
	}

	device.Uninit()
	err = stream.CloseSend()
	if err != nil {
		streamStopped = true
	}
	return ret, err
}

// Function to convert int16 array to []byte
func int16SliceToBytes(input []int16) []byte {
	output := make([]byte, len(input)*2)
	for i := range input {
		binary.LittleEndian.PutUint16(output[i*2:i*2+2], uint16(input[i]))
	}
	return output
}
