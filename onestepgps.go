package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Define the API key and the OneStepGPS API URL
const apiKey = "Xl-8_ceibpMHqr4YZ72uFy5xQfjbOPXstocE8b_Zkmw"
const apiURL = "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=" + apiKey

// Device represents the structure for device information (simplified for frontend use)
type Device struct {
	ID          string  `json:"device_id"`
	Name        string  `json:"display_name"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lng"`
	Active      bool    `json:"online"`
	DriveStatus string  `json:"drive_status"`
}

// Preferences represents user preferences
type Preferences struct {
	SortOrder     string            `json:"sort_order"`
	HiddenDevices []string          `json:"hidden_devices"`
	CustomIcons   map[string]string `json:"custom_icons"` // icon paths per device ID
}

// Server holds device data and user preferences
type Server struct {
	mu          sync.Mutex
	preferences Preferences
}

// NewServer initializes a new server with default preferences
func NewServer() *Server {
	return &Server{
		preferences: Preferences{
			SortOrder:     "name",
			HiddenDevices: []string{},
			CustomIcons:   make(map[string]string),
		},
	}
}

// fetchDevices fetches device data from OneStepGPS API
func (s *Server) fetchDevices() ([]Device, error) {
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer response.Body.Close()

	var data struct {
		ResultList []struct {
			DeviceID          string `json:"device_id"`
			DisplayName       string `json:"display_name"`
			Online            bool   `json:"online"`
			LatestDevicePoint struct {
				Lat         float64 `json:"lat"`
				Lng         float64 `json:"lng"`
				DeviceState struct {
					DriveStatus string `json:"drive_status"`
				} `json:"device_state"`
			} `json:"latest_device_point"`
		} `json:"result_list"`
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	// Map relevant fields to Device structure
	devices := make([]Device, len(data.ResultList))
	for i, d := range data.ResultList {
		devices[i] = Device{
			ID:          d.DeviceID,
			Name:        d.DisplayName,
			Latitude:    d.LatestDevicePoint.Lat,
			Longitude:   d.LatestDevicePoint.Lng,
			Active:      d.Online,
			DriveStatus: d.LatestDevicePoint.DeviceState.DriveStatus,
		}
	}
	return devices, nil
}

// getDevicesHandler handles requests to fetch devices
func (s *Server) getDevicesHandler(w http.ResponseWriter, r *http.Request) {
	devices, err := s.fetchDevices()
	if err != nil {
		http.Error(w, "Error fetching device data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(devices)
}

// getPreferencesHandler handles requests to fetch user preferences
func (s *Server) getPreferencesHandler(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.preferences)
}

// setPreferencesHandler handles requests to update user preferences
func (s *Server) setPreferencesHandler(w http.ResponseWriter, r *http.Request) {
	var prefs Preferences
	if err := json.NewDecoder(r.Body).Decode(&prefs); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	s.preferences = prefs
	s.mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	server := NewServer()

	http.HandleFunc("/api/devices", server.getDevicesHandler)
	http.HandleFunc("/api/preferences", server.getPreferencesHandler)
	http.HandleFunc("/api/preferences/set", server.setPreferencesHandler)

	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
