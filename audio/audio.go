package audio

import (
	"fmt"
	"ggresillion/disvoice/config"
	"ggresillion/disvoice/plugin"

	"github.com/gordonklaus/portaudio"
)

var streamParameters portaudio.StreamParameters
var stream *portaudio.Stream

func Configure(config config.AudioConfig) {
	portaudio.Terminate()
	fmt.Println("Initializing portaudio...")
	portaudio.Initialize()
	h, err := portaudio.DefaultHostApi()
	chk(err)
	streamParameters = portaudio.LowLatencyParameters(h.DefaultInputDevice, h.DefaultOutputDevice)
	streamParameters.Input.Channels = config.Channels
	streamParameters.Output.Channels = config.Channels
	streamParameters.FramesPerBuffer = config.BufferLength
	fmt.Println("Portaudio initialized")
}

func Start() {
	var err error
	stream, err = portaudio.OpenStream(streamParameters, func(in, out []float32) {
		plugin.ProcessAudio(in, out)
	})
	chk(err)
	chk(stream.Start())
}

func Stop() {
	fmt.Println("Cleaning Up")
	portaudio.Terminate()
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
