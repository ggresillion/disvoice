package main

import (
	"github.com/ggresillion/disvoice/backend"
	"github.com/ggresillion/disvoice/plugin"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func start() {
	plugin.Start()
}

func getConfig() backend.Config {
	return *backend.GetConfig()
}

func toggleEffect(id string) {
	effect := backend.GetConfig().GetEffect(id)
	backend.GetConfig().GetPlugin(effect.PluginID).Load()
	effect.ToggleEffect()
}

func showSettings(id string) {
	effect := backend.GetConfig().GetEffect(id)
	plugin := backend.GetConfig().GetPlugin(effect.PluginID)
	plugin.Load()
	plugin.Open()
}

func main() {

	js := mewn.String("./frontend/dist/my-app/main.js")
	css := mewn.String("./frontend/dist/my-app/styles.scss")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "disvoice",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(start)
	app.Bind(toggleEffect)
	app.Bind(getConfig)
	app.Bind(showSettings)
	plugin.Start()
	app.Run()
}
