package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func AuthenticateUser(email, password string) error {
	fmt.Println("Authenticating user with email:", email)
	var hashedPassword string

	err := DB.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		log.Println("Authentication failed: user not found")
		return errors.New("invalid email or password")
	} else if err != nil {
		log.Println("Database error during authentication:", err)
		return err
	}

	err = CheckPassword(hashedPassword, password)
	if err != nil {
		log.Println("Authentication failed: password mismatch")
		return errors.New("invalid email or password")
	}

	log.Println("Authentication successful for user:", email)
	return nil
}

func RegisterUser(email, password string) error {
	if !isValidEmail(email) {
		return errors.New("invalid email format")
	}
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return errors.New("failed to hash password")
	}

	_, err = DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", email, hashedPassword)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return errors.New("user already exists")
		}
		log.Println("Error inserting user into database:", err)
		return err
	}

	log.Println("User registered successfully:", email)
	return nil
}

func CreateSession(email string) {
	sessionToken := uuid.NewString()

	_, err := DB.Exec("INSERT INTO sessions (email, session_token) VALUES (?, ?) ON DUPLICATE KEY UPDATE session_token = ?", email, sessionToken, sessionToken)
	if err != nil {
		log.Println("Error creating session:", err)
		return
	}

	log.Println("Session created for user:", email)
}

func ValidateSession(email string) bool {
	var exists bool

	err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM sessions WHERE email = ?)", email).Scan(&exists)
	if err != nil {
		log.Println("Error validating session:", err)
		return false
	}

	log.Println("Session validation for user:", email, "exists:", exists)
	return exists
}

func DeleteSession(email string) {
	_, err := DB.Exec("DELETE FROM sessions WHERE email = ?", email)
	if err != nil {
		log.Println("Error deleting session:", err)
		return
	}

	log.Println("Session deleted for user:", email)
}
