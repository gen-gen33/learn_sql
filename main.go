package main

import (
	"log"
)

func main() {
	InitConfig("config.json")
	CheckDatabaseConnection()

	db := ConnectDB()
	defer db.Close()

	CreateTable(db, "users")
	InsertUser(db, "users", "John", 23)
	InsertUser(db, "users", "Gen", 23)
	log.Println("add 2 users")
	GetUsers(db)

	deleteRecord(db, "users", 1) // ID = 1 のレコードを削除
	log.Println("delete record id = 1")
	GetUsers(db)

	updateRecord(db, "users", 2, "Name", "new Gen") // ID = 1 のレコードを更新
	log.Println("update username id = 2 Gen to new value")
	GetUsers(db)

	DropTable(db, "users")
	// GetUsers(db)
}
