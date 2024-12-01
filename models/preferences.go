package models

type Preferences struct {
	Email        string `json:"email"`
	SortOrder    string `json:"sortOrder"`
	HideInactive bool   `json:"hideInactive"`
	MapZoomLevel int    `json:"mapZoomLevel"`
}
