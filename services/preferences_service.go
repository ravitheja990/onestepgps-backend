package services

import (
	"database/sql"
	"errors"
	"log"
	"onestepgps-backend/models"
)

func SavePreferences(preferences models.Preferences) error {
	log.Println("Saving these preferences:", preferences)

	_, err := models.DB.Exec(`
        INSERT INTO preferences (email, sort_order, hide_inactive, map_zoom_level)
        VALUES (?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE
        sort_order = VALUES(sort_order),
        hide_inactive = VALUES(hide_inactive),
        map_zoom_level = VALUES(map_zoom_level)
    `, preferences.Email, preferences.SortOrder, preferences.HideInactive, preferences.MapZoomLevel)

	if err != nil {
		log.Println("Error saving preferences:", err)
		return errors.New("failed to save preferences")
	}

	log.Println("Preferences saved for user:", preferences.Email)
	return nil
}

func GetPreferences(email string) (models.Preferences, error) {
	var preferences models.Preferences

	err := models.DB.QueryRow(`
        SELECT email, sort_order, hide_inactive, map_zoom_level
        FROM preferences
        WHERE email = ?
    `, email).Scan(&preferences.Email, &preferences.SortOrder, &preferences.HideInactive, &preferences.MapZoomLevel)

	if err == sql.ErrNoRows {
		log.Println("Preferences not found for user:", email)
		return preferences, errors.New("preferences not found")
	}

	if err != nil {
		log.Println("Error retrieving preferences:", err)
		return preferences, errors.New("failed to fetch preferences")
	}

	log.Println("Preferences retrieved for user:", email)
	return preferences, nil
}
