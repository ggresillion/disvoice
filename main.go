package main

import (
	"fmt"
	"ggresillion/disvoice/audio"
	"ggresillion/disvoice/config"
	"ggresillion/disvoice/plugin"
	"ggresillion/disvoice/window"
)

var p plugin.Plugin

func main() {

	fmt.Println("Starting disvoice...")
	p = plugin.InitPlugin("plugins/TAL-Reverb-4-64.dll")
	config := config.AudioConfig{SampleRate: 44100, BufferLength: 256, Channels: 1}
	p.Configure(config)

	fmt.Println("Creating new window...")
	window.Create()
	// fmt.Println("Open plugin GUI...")
	// p.OpenGui(hwnd)
	// fmt.Println("Setting window size...")
	// width, height := p.GetEditorRect()
	// window.SetSize(hwnd, width, height)
	// fmt.Println("Starting window...")

	// p.LoadSettings()
	audio.Configure(config)
	audio.StartPlayback()
	p.Start()
	audio.SetProcessPlugin(p)
	// window.Show(hwnd)
	// p.SaveSettings()
	for true {

	}
}

func runGUI() {
	p.SaveSettings()

}
