package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"onestepgps-backend/models"
)

const apiKey = "Xl-8_ceibpMHqr4YZ72uFy5xQfjbOPXstocE8b_Zkmw"
const apiURL = "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=" + apiKey

// FetchDevices fetches and returns device data from the OneStepGPS API
func FetchDevices() ([]models.Device, error) {
	// Perform GET request
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer response.Body.Close()

	// Parse API response
	var data struct {
		ResultList []struct {
			DeviceID          string `json:"device_id"`
			DisplayName       string `json:"display_name"`
			Online            bool   `json:"online"`
			LatestDevicePoint struct {
				Lat         float64 `json:"lat"`
				Lng         float64 `json:"lng"`
				Altitude    float64 `json:"altitude"`
				Speed       float64 `json:"speed"`
				Angle       float64 `json:"angle"`
				DeviceState struct {
					DriveStatus string `json:"drive_status"`
				} `json:"device_state"`
			} `json:"latest_device_point"`
		} `json:"result_list"`
	}

	// Decode the response body
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	// Map data to the models.Device structure
	devices := make([]models.Device, len(data.ResultList))
	for i, d := range data.ResultList {
		devices[i] = models.Device{
			ID:          d.DeviceID,
			Name:        d.DisplayName,
			Active:      d.Online,
			DriveStatus: d.LatestDevicePoint.DeviceState.DriveStatus,
			CurrentPosition: map[string]interface{}{
				"Lat":   d.LatestDevicePoint.Lat,
				"Lng":   d.LatestDevicePoint.Lng,
				"Alt":   d.LatestDevicePoint.Altitude,
				"Speed": d.LatestDevicePoint.Speed,
				"Angle": d.LatestDevicePoint.Angle,
			},
		}
	}
	return devices, nil
}

// PrintDevices formats and prints the fetched device information
func PrintDevices(devices []models.Device) {
	if len(devices) == 0 {
		fmt.Println("No devices found.")
		return
	}

	for _, device := range devices {
		// Extract current_position data for printing
		currentPosition := device.CurrentPosition
		fmt.Printf("Name: %s\n", device.Name)
		fmt.Printf("Current Position: {Lat=%.6f, Lng=%.6f, Alt=%.2f m, Angle=%.2f degrees}\n",
			currentPosition["Lat"], currentPosition["Lng"], currentPosition["Alt"], currentPosition["Angle"])
		fmt.Printf("Active: %t\n", device.Active)
		fmt.Printf("Drive Status: %s\n", device.DriveStatus)
		fmt.Println("--------------------------------")
	}
}
