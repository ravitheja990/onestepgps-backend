package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"onestepgps-backend/models"
)

const apiKey = "Xl-8_ceibpMHqr4YZ72uFy5xQfjbOPXstocE8b_Zkmw"
const apiURL = "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=" + apiKey

// FetchDevices fetches device data from OneStepGPS API
func FetchDevices() ([]models.Device, error) {
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

	devices := make([]models.Device, len(data.ResultList))
	for i, d := range data.ResultList {
		devices[i] = models.Device{
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
