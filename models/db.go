package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Update the connection string with your MySQL credentials
	DB, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/onestepgps")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Verify the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connected successfully.")
}
