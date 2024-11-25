package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	// グローバル変数 AppConfig を利用
	db, err := sql.Open("mysql", AppConfig.DSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}
