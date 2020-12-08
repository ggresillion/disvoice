package audio

import (
	"fmt"
	"ggresillion/disvoice/config"
	"ggresillion/disvoice/plugin"

	"github.com/gordonklaus/portaudio"
)

var streamParameters portaudio.StreamParameters
var stream *portaudio.Stream
var toggle bool

func Configure(config config.AudioConfig) {
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

func Start(cb func(in, out []float32)) {
	var err error
	stream, err = portaudio.OpenStream(streamParameters, cb)
	chk(err)
	chk(stream.Start())
}

func Toggle() {
	toggle = true
}

func process(in, out []float32) {
	fmt.Println(toggle)
	if toggle == true {
		plugin.ProcessAudio(in, out)
	} else {
		for i := range out {
			out[i] = in[i]
		}
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
