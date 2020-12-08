package main

import (
	"fmt"
	"ggresillion/disvoice/audio"
	"ggresillion/disvoice/config"
	"ggresillion/disvoice/plugin"
	"ggresillion/disvoice/window"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	config := config.AudioConfig{SampleRate: 44100, BufferLength: 256, Channels: 1}

	audio.Configure(config)

	plugin.InitPlugin("plugins/TAL-Reverb-4-64.dll")

	plugin.Configure(config)
	plugin.Start()

	audio.Start()
	go func() {
		fmt.Println("WTF")
		fmt.Println(plugin.GetEditorRect())
		fmt.Println("DONE")
	}()

	fmt.Println("Creating new window...")
	hwnd := window.Create()
	fmt.Println("Open plugin GUI...")
	plugin.OpenGui(hwnd)
	// fmt.Println("Setting window size...")
	// width, height := p.GetEditorRect()
	// window.SetSize(hwnd, width, height)
	// fmt.Println("Starting window...")

	// p.LoadSettings()

	fmt.Println("Showing window")
	window.Show(hwnd)
	// p.SaveSettings()
	for true {

	}
}

func cleanup() {
	audio.Stop()
	plugin.Stop()
}
