package models

type Device struct {
	ID          string  `json:"device_id"`
	Name        string  `json:"display_name"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lng"`
	Active      bool    `json:"online"`
	DriveStatus string  `json:"drive_status"`
}
