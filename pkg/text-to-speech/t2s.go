package text_to_speech

import (
	"bytes"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"context"
	"errors"
	"github.com/gen2brain/malgo"
	"github.com/hajimehoshi/go-mp3"
	"go.uber.org/zap"
	"io"
	"log"
)

func TranscribeText(text string) error {
	// Instantiates a client.
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Print("failed to create text to speech client", zap.Error(err))
		return err
	}
	defer client.Close()

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			Name:         "en-US-Studio-Q",
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_MALE,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Print("failed to synthesize speach", zap.Error(err))
		return err
	}
	mp3Bytes := resp.AudioContent

	context, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
		//fmt.Printf("LOG <%v>\n", message)
	})
	if err != nil {
		log.Print("failed to init malgo", zap.Error(err))
		return err
	}
	defer func() {
		_ = context.Uninit()
		context.Free()
	}()

	reader, err := mp3.NewDecoder(bytes.NewReader(mp3Bytes))
	if err != nil {
		log.Print("failed to create new decoder", zap.Error(err))
		return err
	}

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Playback)
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 2
	deviceConfig.SampleRate = uint32(reader.SampleRate())
	deviceConfig.Alsa.NoMMap = 1

	// This is the function that's used for sending more data to the device for playback.
	finishedProcessing := false
	onSamples := func(pOutputSample, pInputSamples []byte, framecount uint32) {
		_, err = io.ReadFull(reader, pOutputSample)
		if errors.Is(io.ErrUnexpectedEOF, err) {
			finishedProcessing = true
		} else if err != nil {
			log.Print("failed to read full", zap.Error(err))
		}
	}
	deviceCallbacks := malgo.DeviceCallbacks{
		Data: onSamples,
	}

	device, err := malgo.InitDevice(context.Context, deviceConfig, deviceCallbacks)
	if err != nil {
		log.Print("failed to init device", zap.Error(err))
		return err
	}
	defer device.Uninit()

	err = device.Start()
	if err != nil {
		log.Print("failed to start device", zap.Error(err))
		return err
	}

	for {
		if finishedProcessing {
			break
		}
	}
	return nil
}
