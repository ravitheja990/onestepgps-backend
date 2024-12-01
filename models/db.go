package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/onestepgps")
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		return err
	}

	if err = DB.Ping(); err != nil {
		log.Println("Database ping failed. Please check your connection settings.")
		return err
	}

	fmt.Println("Database connection established.")
	return nil
}
