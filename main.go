package main

import (
	"fmt"
	"log"
	"net/http"
	"onestepgps-backend/controllers"
)

func main() {
	http.HandleFunc("/api/devices", controllers.GetDevicesHandler)
	http.HandleFunc("/api/preferences", controllers.GetPreferencesHandler)
	http.HandleFunc("/api/preferences/set", controllers.SetPreferencesHandler)

	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
