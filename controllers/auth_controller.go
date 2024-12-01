package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"onestepgps-backend/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside signuphandler")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var loginData models.User
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Authenticate user
	err = models.AuthenticateUser(loginData.Email, loginData.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Create a session
	models.CreateSession(loginData.Email)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Login successful"}`))
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside signuphandler")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var signupData models.User
	err := json.NewDecoder(r.Body).Decode(&signupData)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = models.RegisterUser(signupData.Email, signupData.Password)
	fmt.Println("signupData.Email is :: ", signupData.Email)
	fmt.Println("signupData.Password is :: ", signupData.Password)
	if err != nil {
		if err.Error() == "user already exists" {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(`{"message": "User already registered"}`))
			return
		}
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Signup successful"}`))
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.Header.Get("X-Session-Email")
		fmt.Println("AuthMiddleware: Received email from header:", email)

		if !models.ValidateSession(email) {
			fmt.Println("AuthMiddleware: Unauthorized, session not found for email:", email)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		fmt.Println("AuthMiddleware: Session validated for email:", email)
		next(w, r)
	}
}
