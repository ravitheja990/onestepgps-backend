package controllers

import (
	"encoding/json"
	"net/http"
	"onestepgps-backend/services"
)

// GetDevicesHandler handles requests to fetch devices
func GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	devices, err := services.FetchDevices()
	if err != nil {
		http.Error(w, "Error fetching device data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(devices)
}
