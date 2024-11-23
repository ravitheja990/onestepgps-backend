package main

import (
	"fmt"
	"log"
	"net/http"
	"onestepgps-backend/controllers"
)

// CORS Middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins, or specify the frontend origin like "http://localhost:3000"
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Session-Email")

		// Handle preflight requests (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	// Add your handlers
	mux.HandleFunc("/api/login", controllers.LoginHandler)
	mux.HandleFunc("/api/signup", controllers.SignupHandler)
	mux.HandleFunc("/api/devices", controllers.AuthMiddleware(controllers.GetDevicesHandler))
	mux.HandleFunc("/api/preferences", controllers.AuthMiddleware(controllers.GetPreferencesHandler))
	mux.HandleFunc("/api/preferences/set", controllers.AuthMiddleware(controllers.SetPreferencesHandler))

	// Wrap the handlers with the CORS middleware
	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", corsMiddleware(mux)); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
