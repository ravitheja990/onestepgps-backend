package main

import (
	"log"
	"net/http"
	"onestepgps-backend/controllers"
	"onestepgps-backend/models"
)

func main() {
	// Initialize the database
	models.InitDB()

	// Wrap all routes with the CORS middleware
	http.HandleFunc("/login", corsMiddleware(controllers.LoginHandler))
	http.HandleFunc("/signup", corsMiddleware(controllers.SignupHandler))
	http.HandleFunc("/devices", corsMiddleware(controllers.AuthMiddleware(controllers.GetDevicesHandler)))
	http.HandleFunc("/preferences", corsMiddleware(controllers.AuthMiddleware(controllers.SavePreferencesHandler)))
	http.HandleFunc("/preferences/get", corsMiddleware(controllers.AuthMiddleware(controllers.GetPreferencesHandler)))

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081") // Allow requests from frontend origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Session-Email")

		// Handle preflight (OPTIONS) request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next(w, r)
	}
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Access granted to protected route"))
}
