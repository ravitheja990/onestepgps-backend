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
	log.Printf("Raw body: %s\n", string(bodyBytes))

	var preferences models.Preferences
	err := json.Unmarshal(bodyBytes, &preferences)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	preferences.Email = email
	log.Printf("Decoded preferences object: %+v\n", preferences)

	err = services.SavePreferences(preferences)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Preferences saved successfully"))
}

// GetPreferencesHandler handles fetching user preferences
func GetPreferencesHandler(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("X-Session-Email")
	if email == "" {
		http.Error(w, "Missing session email", http.StatusUnauthorized)
		return
	}

	preferences, err := services.GetPreferences(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(preferences)
}
