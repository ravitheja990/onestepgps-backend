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
  - [API Endpoints](#api-endpoints)


## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ravitheja990/onestepgps-backend.git
   cd onestepgps-backend

2. Install dependencies: Ensure you have Go 1.23.3 or later installed.

3.  Build the project:

   ```bash
   go build

## Project Structure (MVC)
- main.go: Entry point of the application
- controllers/: Contains the controllers that handle HTTP requests and responses
- models/: Contains the data models and business logic
- services/: Contains the service layer, which manages interaction with the external OneStepGPS API
- utils/: Contains utility functions and shared components

## Configuration
The API requires an API key from OneStepGPS to fetch the GPS data.

Update the apiKey constant in services/device_service.go with your OneStepGPS API key.
Usage
Start the server:

```bash
go run main.go
The server will run on http://localhost:8080 by default.