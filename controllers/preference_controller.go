package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"onestepgps-backend/models"
	"onestepgps-backend/services"
)

func SavePreferencesHandler(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("X-Session-Email")
	if email == "" {
		http.Error(w, "Missing session email", http.StatusUnauthorized)
		return
	}

	bodyBytes, _ := io.ReadAll(r.Body)
	log.Printf("Raw body received: %s\n", string(bodyBytes))

	var preferences models.Preferences
	if err := json.Unmarshal(bodyBytes, &preferences); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	preferences.Email = email
	log.Printf("Parsed preferences: %+v\n", preferences)

	if err := services.SavePreferences(preferences); err != nil {
		http.Error(w, "Failed to save preferences", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Preferences saved successfully"))
}

func GetPreferencesHandler(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("X-Session-Email")
	if email == "" {
		http.Error(w, "Missing session email", http.StatusUnauthorized)
		return
	}

	preferences, err := services.GetPreferences(email)
	if err != nil {
		http.Error(w, "Preferences not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(preferences)
}
