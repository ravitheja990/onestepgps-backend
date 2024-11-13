package models

type Preferences struct {
	SortOrder     string            `json:"sort_order"`
	HiddenDevices []string          `json:"hidden_devices"`
	CustomIcons   map[string]string `json:"custom_icons"`
}
