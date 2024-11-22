package main

import (
	"fmt"
	"log"
	"net/http"
	"onestepgps-backend/controllers"
)

func main() {
	// Login and Signup routes
	http.HandleFunc("/api/signup", controllers.SignupHandler)
	http.HandleFunc("/api/login", controllers.LoginHandler)

	// Protect other APIs
	http.HandleFunc("/api/devices", controllers.AuthMiddleware(controllers.GetDevicesHandler))
	http.HandleFunc("/api/preferences", controllers.AuthMiddleware(controllers.GetPreferencesHandler))
	http.HandleFunc("/api/preferences/set", controllers.AuthMiddleware(controllers.SetPreferencesHandler))

	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
