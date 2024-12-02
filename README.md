# OneStepGPS Backend (MVC Version)

This project is a backend service for interacting with the OneStepGPS API, retrieving and managing GPS device data and user preferences. It follows the Model-View-Controller (MVC) architecture, making the code more modular, maintainable, and scalable.

## Features

- Fetch the latest GPS data from the OneStepGPS API
- Retrieve and update user preferences for device sorting and customizations
- Expose RESTful endpoints to interact with devices and preferences

## Table of Contents

- [OneStepGPS Backend (MVC Version)](#onestepgps-backend-mvc-version)
  - [Features](#features)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Project Structure (MVC)](#project-structure-mvc)
  - [Configuration](#configuration)

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/ravitheja990/onestepgps-backend.git
   cd onestepgps-backend

2. Install dependencies: Ensure you have Go 1.23.3 or higher version installed.

3. DB Setup: The sql files are stored in db folder and the database name is onestepgps. 
    ```
    mysql -u <username> -p <database_name> < /path/to/file.sql

4. OneStepGPS Api Key Setup: 
   Store the api key in below format:
   ```
   API_KEY=<onestepgps-api-key>

## Project Structure:

- main.go The app's entry point, sets up the server and routes.

- controllers/
  - `auth_controller.go`: User login and signup.
  - `device_controller.go`: Device-related APIs.
  - `preference_controller.go`: User preferences APIs.

- db/
  - `onestepgps_preferences.sql`: User preferences schema.
  - `onestepgps_sessions.sql`: Session management schema.
  - `onestepgps_users.sql`: User accounts schema.

- models/
  - `db.go`: Database connection setup.
  - `device.go`, `preferences.go`, `user.go`: Data models for devices, preferences, and users.

- services/
  - `device_service.go`: Fetches device data from the OneStepGPS API.
  - `preferences_service.go`: Manages user preferences.

## Configuration
The API requires an API key from OneStepGPS to fetch the GPS data.

Update the apiKey constant in services/device_service.go with your OneStepGPS API key.
Usage
Start the server:

```
go run main.go
The server will run on http://localhost:8080.