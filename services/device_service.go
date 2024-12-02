package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"onestepgps-backend/models"
	"os"

	"github.com/joho/godotenv"
)

const apiBaseURL = "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key="

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, proceeding with environment variables.")
	}
}

func FetchDevices() ([]models.Device, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("API_KEY environment variable is not set")
	}

	apiURL := apiBaseURL + apiKey
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
