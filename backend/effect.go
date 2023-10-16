package backend

type Effect struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	PluginID string `json:"pluginId"`
}

func (e *Effect) ToggleEffect() {
	plugin := GetConfig().GetPlugin(e.PluginID)
	plugin.Load()
}
