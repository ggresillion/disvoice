package audio

import (
	"fmt"
	"ggresillion/disvoice/config"
	"ggresillion/disvoice/plugin"

	"github.com/gordonklaus/portaudio"
)

var streamParameters portaudio.StreamParameters
var stream *portaudio.Stream
var currentPlugin *plugin.Plugin

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

func StartPlayback() {
	fmt.Println("Starting playback")
	var err error
	stream, err = portaudio.OpenStream(streamParameters, process)
	chk(err)
	chk(stream.Start())
}

func SetProcessPlugin(p plugin.Plugin) {
	currentPlugin = &p
	chk(stream.Stop())
	var err error
	stream, err = portaudio.OpenStream(streamParameters, func(in, out []float32) {
		process(in, out)
	})
	chk(err)
	chk(stream.Start())
}

func process(in, out []float32) {
	if currentPlugin == nil {
		for i := range out {
			out[i] = in[i]
		}
	} else {
		fmt.Println("processing")
		currentPlugin.ProcessAudio(in, out)
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
