package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/go-sql-driver/mysql" // Import MySQL driver package
	"github.com/google/uuid"         // Import UUID for secure session tokens
	"golang.org/x/crypto/bcrypt"     // Import bcrypt for password hashing
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate email format
func isValidEmail(email string) bool {
	regex := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

// HashPassword hashes a plaintext password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// CheckPassword verifies a hashed password against a plaintext password
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// AuthenticateUser verifies user credentials against the database
func AuthenticateUser(email, password string) error {
	fmt.Println("Authenticating user with email:", email)
	var hashedPassword string

	// Fetch hashed password from the database
	err := DB.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		log.Println("Authentication failed: user not found")
		return errors.New("invalid email or password")
	} else if err != nil {
		log.Println("Database error during authentication:", err)
		return err
	}

	// Verify the hashed password
	err = CheckPassword(hashedPassword, password)
	if err != nil {
		log.Println("Authentication failed: password mismatch")
		return errors.New("invalid email or password")
	}

	log.Println("Authentication successful for user:", email)
	return nil
}

// RegisterUser registers a new user in the database
func RegisterUser(email, password string) error {
	if !isValidEmail(email) {
		return errors.New("invalid email format")
	}
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	// Hash the password before storing it
	hashedPassword, err := HashPassword(password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return errors.New("failed to hash password")
	}

	// Insert user into the database
	_, err = DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", email, hashedPassword)
	if err != nil {
		// Handle duplicate key errors
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return errors.New("user already exists")
		}
		log.Println("Error inserting user into database:", err)
		return err
	}

	log.Println("User registered successfully:", email)
	return nil
}

// CreateSession creates a secure session for the user in the database
func CreateSession(email string) {
	// Generate a secure session token
	sessionToken := uuid.NewString()

	// Insert or update the session in the database
	_, err := DB.Exec("INSERT INTO sessions (email, session_token) VALUES (?, ?) ON DUPLICATE KEY UPDATE session_token = ?", email, sessionToken, sessionToken)
	if err != nil {
		log.Println("Error creating session:", err)
		return
	}

	log.Println("Session created for user:", email)
}

// ValidateSession checks if a session exists in the database
func ValidateSession(email string) bool {
	var exists bool

	// Query to check session existence
	err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM sessions WHERE email = ?)", email).Scan(&exists)
	if err != nil {
		log.Println("Error validating session:", err)
		return false
	}

	log.Println("Session validation for user:", email, "exists:", exists)
	return exists
}

// DeleteSession deletes a session for the user in the database
func DeleteSession(email string) {
	// Delete session from the database
	_, err := DB.Exec("DELETE FROM sessions WHERE email = ?", email)
	if err != nil {
		log.Println("Error deleting session:", err)
		return
	}

	log.Println("Session deleted for user:", email)
}
