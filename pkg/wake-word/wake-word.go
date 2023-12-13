package wake_word

import (
	"encoding/binary"
	porcupine "github.com/Picovoice/porcupine/binding/go/v2"
	"github.com/gen2brain/malgo"
	"go.uber.org/zap"
	"log"
)

func DetectWakeWord(accessKey string) (porcupine.BuiltInKeyword, error) {
	var err error
	backends := []malgo.Backend{malgo.BackendAlsa}
	context, err := malgo.InitContext(backends, malgo.ContextConfig{}, func(message string) {
		//log.Printf(fmt.Sprintf("%v\n", message))
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

	p := porcupine.Porcupine{
		BuiltInKeywords: []porcupine.BuiltInKeyword{porcupine.HEY_GOOGLE, porcupine.BUMBLEBEE},
		// TODO move this
		//KeywordPaths: []string{"./Hey-Billy-Bass_en_raspberry-pi_v3_0_0.ppn"},
		AccessKey: accessKey,
	}
	err = p.Init()
	if err != nil {
		return "", err
	}
	defer p.Delete()

	var shortBufIndex, shortBufOffset int
	shortBuf := make([]int16, porcupine.FrameLength)

	var keyword porcupine.BuiltInKeyword
	finishedProcessing := false
	onRecvFrames := func(pSample2, pSample []byte, framecount uint32) {
		for i := 0; i < len(pSample); i += 2 {
			shortBuf[shortBufIndex+shortBufOffset] = int16(binary.LittleEndian.Uint16(pSample[i : i+2]))
			shortBufOffset++

			if shortBufIndex+shortBufOffset == porcupine.FrameLength {
				shortBufIndex = 0
				shortBufOffset = 0
				keywordIndex, err := p.Process(shortBuf)
				if err != nil {
					log.Print("Error on processing key word", zap.Error(err))
				} else {
					if keywordIndex >= 0 {
						finishedProcessing = true
						keyword = p.BuiltInKeywords[keywordIndex]
					}
				}
			}
		}
		shortBufIndex += shortBufOffset
		shortBufOffset = 0
	}

	captureCallbacks := malgo.DeviceCallbacks{
		Data: onRecvFrames,
	}
	device, err := malgo.InitDevice(context.Context, deviceConfig, captureCallbacks)
	if err != nil {
		log.Print("Error on init device", zap.Error(err))
		return "", err
	}

	err = device.Start()
	if err != nil {
		log.Print("Error on start device", zap.Error(err))
		return "", err
	}

	for {
		if finishedProcessing {
			break
		}
	}
	device.Uninit()

	return keyword, err
}
