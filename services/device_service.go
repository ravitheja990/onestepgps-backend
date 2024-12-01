package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"onestepgps-backend/models"
	"strings"
)

const (
	apiKey = "Xl-8_ceibpMHqr4YZ72uFy5xQfjbOPXstocE8b_Zkmw"
	apiURL = "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=" + apiKey
)

func FetchDevices() ([]models.Device, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
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

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %w", err)
	}

	devices := make([]models.Device, len(result.ResultList))
	for i, item := range result.ResultList {
		devices[i] = models.Device{
			ID:          item.DeviceID,
			Name:        item.DisplayName,
			Active:      item.Online,
			DriveStatus: item.LatestDevicePoint.DeviceState.DriveStatus,
			CurrentPosition: map[string]interface{}{
				"Lat":   item.LatestDevicePoint.Lat,
				"Lng":   item.LatestDevicePoint.Lng,
				"Alt":   item.LatestDevicePoint.Altitude,
				"Speed": item.LatestDevicePoint.Speed,
				"Angle": item.LatestDevicePoint.Angle,
			},
		}
	}
	return devices, nil
}

// PrintDevices outputs the fetched device information in a formatted manner.
func PrintDevices(devices []models.Device) {
	if len(devices) == 0 {
		fmt.Println("No devices available.")
		return
	}

	for _, device := range devices {
		pos := device.CurrentPosition
		fmt.Printf("Name: %s\n", device.Name)
		fmt.Printf("Position: {Lat: %.6f, Lng: %.6f, Alt: %.2f m, Angle: %.2f°}\n",
			pos["Lat"], pos["Lng"], pos["Alt"], pos["Angle"])
		fmt.Printf("Active: %t\n", device.Active)
		fmt.Printf("Drive Status: %s\n", device.DriveStatus)
		fmt.Println(strings.Repeat("-", 40))
	}
}
