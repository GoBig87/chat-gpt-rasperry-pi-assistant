package text_to_speech

import (
	"bytes"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"context"
	"errors"
	"github.com/gen2brain/malgo"
	"github.com/hajimehoshi/go-mp3"
	"io"
	"log"
)

func TranscribeText(text string) error {
	// Instantiates a client.
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to create text to speech client: \n %v", err)
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
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Fatalf("failed to synthesize speach: \n %v", err)
		return err
	}
	mp3Bytes := resp.AudioContent

	context, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
		//fmt.Printf("LOG <%v>\n", message)
	})
	if err != nil {
		log.Fatalf("failed to init malgo: \n %v", err)
		return err
	}
	defer func() {
		_ = context.Uninit()
		context.Free()
	}()

	reader, err := mp3.NewDecoder(bytes.NewReader(mp3Bytes))
	if err != nil {
		log.Fatalf("failed to create new decoder: \n %v", err)
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
			log.Fatalf("failed to read full: \n %v", err)
		}
	}
	deviceCallbacks := malgo.DeviceCallbacks{
		Data: onSamples,
	}

	device, err := malgo.InitDevice(context.Context, deviceConfig, deviceCallbacks)
	if err != nil {
		log.Fatalf("failed to init device :\n %v", err)
		return err
	}
	defer device.Uninit()

	err = device.Start()
	if err != nil {
		log.Fatalf("failed to start device: \n %v", err)
		return err
	}

	for {
		if finishedProcessing {
			break
		}
	}
	return nil
}