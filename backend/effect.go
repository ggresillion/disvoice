package backend

import (
	log "github.com/sirupsen/logrus"
)

type Effect struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	PluginID string `json:"pluginId"`
	plugin   *Plugin
}

var currentEffect *Effect

func GetEffect(id string) *Effect {
	for _, effect := range GetConfig().Effects {
		if effect.ID == id {
			return &effect
		}
	}
	return nil
}

func (e *Effect) ToggleEffect() {
	if currentEffect != nil {
		currentEffect.plugin.Toggle()
		if currentEffect.ID == e.ID {
			return
		}
	}
	plugin := GetPlugin(e.PluginID)
	e.plugin = plugin
	plugin.Load()
	plugin.Toggle()
	currentEffect = e
	log.Info("Toggled effect: ", e.Name)
}
