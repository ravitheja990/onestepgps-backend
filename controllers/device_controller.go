package controllers

import (
	"encoding/json"
	"net/http"
	"onestepgps-backend/services"
)

func GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	devices, err := services.FetchDevices()
	if err != nil {
		http.Error(w, "Failed to fetch device data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(devices); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
