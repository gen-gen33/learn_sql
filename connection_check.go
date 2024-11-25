package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CheckDatabaseConnection() {
	db, err := sql.Open("mysql", AppConfig.DSN())
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Successfully connected to the database!")
}
