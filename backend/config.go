package backend

type Config struct {
	Plugins []Plugin `json:"plugins"`
	Effects []Effect `json:"effects"`
}

func GetConfig() *Config {
	return &Config{
		Effects: []Effect{{ID: "1", Name: "Echo", PluginID: "1"}},
		Plugins: []Plugin{{ID: "1", Name: "TAL-Reverb", Filename: "TAL-Reverb-4-64.dll"}},
	}
}
