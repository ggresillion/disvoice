package main

import (
	"fmt"
	"ggresillion/disvoice/audio"
	"ggresillion/disvoice/config"
	"ggresillion/disvoice/plugin"
	"ggresillion/disvoice/window"
	"time"
)

func main() {

	config := config.AudioConfig{SampleRate: 44100, BufferLength: 256, Channels: 1}

	audio.Configure(config)

	plugin.InitPlugin("plugins/TAL-Reverb-4-64.dll")
	plugin.Configure(config)
	plugin.Start()
	width, height := plugin.GetEditorRect()

	time.Sleep(time.Second)

	// audio.Start(plugin.ProcessAudio)
	// fmt.Println("Setting window size...")
	// p.GetEditorRect()

	fmt.Println("Setting process plugin")

	fmt.Println("Creating new window...")
	hwnd := window.Create()
	fmt.Println("Open plugin GUI...")
	plugin.OpenGui(hwnd)
	// fmt.Println("Setting window size...")
	// width, height := p.GetEditorRect()
	window.SetSize(hwnd, width, height)
	// fmt.Println("Starting window...")

	// p.LoadSettings()

	fmt.Println("Showing window")
	window.Show(hwnd)
	// p.SaveSettings()
}
