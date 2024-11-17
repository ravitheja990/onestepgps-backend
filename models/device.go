package models

type Device struct {
	ID              string                 `json:"device_id"`
	Name            string                 `json:"display_name"`
	Active          bool                   `json:"online"`
	DriveStatus     string                 `json:"drive_status"`
	CurrentPosition map[string]interface{} `json:"current_position"`
}
