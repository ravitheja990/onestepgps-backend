package main

import (
	"fmt"
	"log"
	"net/http"
	"onestepgps-backend/controllers"
	"onestepgps-backend/models"
)

func main() {
	// Initializing the database
	models.InitDB()

	//api routes
	http.HandleFunc("/login", corsMiddleware(controllers.LoginHandler))
	http.HandleFunc("/signup", corsMiddleware(controllers.SignupHandler))
	http.HandleFunc("/devices", corsMiddleware(controllers.AuthMiddleware(controllers.GetDevicesHandler)))
	http.HandleFunc("/save-preferences", corsMiddleware(controllers.AuthMiddleware(controllers.SavePreferencesHandler)))
	http.HandleFunc("/preferences/get", corsMiddleware(controllers.AuthMiddleware(controllers.GetPreferencesHandler)))

	// starting the server at 8080 port
	fmt.Println("Server starting at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server stopped unexpectedly:", err)
	}
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Session-Email")

		if r.Method == http.MethodOptions {
			fmt.Println("OPTIONS request received")
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Accessing protected route")
	w.Write([]byte("Access granted to protected route"))
}
