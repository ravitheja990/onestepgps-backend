package controllers

import (
	"encoding/json"
	"net/http"
	"onestepgps-backend/models"
	"sync"
)

var (
	preferences = models.Preferences{
		SortOrder:     "name",
		HiddenDevices: []string{},
		CustomIcons:   make(map[string]string),
	}
	mu sync.Mutex
)

// GetPreferencesHandler handles requests to fetch user preferences
func GetPreferencesHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(preferences)
}

// SetPreferencesHandler handles requests to update user preferences
func SetPreferencesHandler(w http.ResponseWriter, r *http.Request) {
	var prefs models.Preferences
	if err := json.NewDecoder(r.Body).Decode(&prefs); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	preferences = prefs
	mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
