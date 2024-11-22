package models

import (
	"errors"
	"fmt"
	"sync"
)

// User represents the structure for user credentials
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	users = []User{
		{Email: "test@example.com", Password: "password123"}, // Hardcoded user for testing
	}
	sessions = make(map[string]string) // In-memory session storage
	mu       sync.Mutex
)

// AuthenticateUser verifies user credentials
func AuthenticateUser(email, password string) error {
	fmt.Println("Authenticating user with email:", email)
	for _, user := range users {
		if user.Email == email && user.Password == password {
			fmt.Println("Authentication successful")
			return nil
		}
	}
	fmt.Println("Authentication failed")
	return errors.New("invalid email or password")
}

// RegisterUser registers a new user
func RegisterUser(email, password string) error {
	fmt.Println("inside RegisterUser")
	fmt.Println("email :: ", email)
	fmt.Println("password :: ", password)
	mu.Lock()
	defer mu.Unlock()

	// Check if the user already exists
	for _, user := range users {
		if user.Email == email {
			return errors.New("user already exists")
		}
	}

	// Add the new user
	users = append(users, User{Email: email, Password: password})
	fmt.Println("User registered successfully:", email)
	return nil
}

// CreateSession creates a session for the user
func CreateSession(email string) {
	mu.Lock()
	defer mu.Unlock()
	sessions[email] = email
	fmt.Println("Session created for user:", email)
}

// ValidateSession checks if a session exists
func ValidateSession(email string) bool {
	mu.Lock()
	defer mu.Unlock()
	_, exists := sessions[email]
	fmt.Println("Session validation for user:", email, "exists:", exists)
	return exists
}

// DeleteSession deletes a session for the user
func DeleteSession(email string) {
	mu.Lock()
	defer mu.Unlock()
	delete(sessions, email)
	fmt.Println("Session deleted for user:", email)
}
