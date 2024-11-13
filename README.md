# OneStepGPS Backend

This project is a backend service for interacting with the OneStepGPS API, retrieving and managing GPS device data and user preferences. It includes endpoints to fetch device details, get user preferences, and update preferences.

## Features

- Fetch the latest GPS data from the OneStepGPS API
- Retrieve and update user preferences for device sorting and customizations
- Expose RESTful endpoints to interact with devices and preferences

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Development](#development)

## Installation

Clone the repository:

```bash
git clone https://github.com/ravitheja990/onestepgps-backend.git
cd onestepgps-backend

Install dependencies: Ensure you have Go 1.23.3 or later installed.

go build


Configuration
The API requires an API key from OneStepGPS to fetch the GPS data.

Update the apiKey constant in onestepgps.go with your OneStepGPS API key.

Usage
Start the server:

bash
Copy code

go run onestepgps.go

The server will run on http://localhost:8080 by default.

API Endpoints
GET /api/devices
Fetches the latest device data from the OneStepGPS API.

Response:

json
Copy code


[
  {
    "device_id": "string",
    "display_name": "string",
    "lat": float,
    "lng": float,
    "online": boolean,
    "drive_status": "string"
  },
  ...
]

GET /api/preferences
Retrieves user preferences such as sort order, hidden devices, and custom icons.

Response:

json
Copy code

{
  "sort_order": "string",
  "hidden_devices": ["device_id1", "device_id2"],
  "custom_icons": {
    "device_id": "icon_path"
  }
}

POST /api/preferences/set
Updates user preferences.

Request Body:

{
  "sort_order": "string",
  "hidden_devices": ["device_id1", "device_id2"],
  "custom_icons": {
    "device_id": "icon_path"
  }
}