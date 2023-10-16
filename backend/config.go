package backend

import "github.com/ggresillion/disvoice/plugin"

type Config struct {
	Plugins []plugin.Plugin `json:"plugins"`
	Effects []Effect        `json:"effects"`
}

func GetConfig() *Config {
	return &Config{
		Effects: []Effect{{ID: "1", Name: "Echo", PluginID: "1"}},
		Plugins: []plugin.Plugin{{ID: "1", Name: "TAL-Reverb", Filename: "TAL-Reverb-4-64.dll"}},
	}
}

func (c *Config) GetEffect(id string) *Effect {
	for _, effect := range c.Effects {
		if effect.ID == id {
			return &effect
		}
	}
	return nil
}

func (c *Config) GetPlugin(id string) *plugin.Plugin {
	for _, plugin := range c.Plugins {
		if plugin.ID == id {
			return &plugin
		}
	}
	return nil
}
