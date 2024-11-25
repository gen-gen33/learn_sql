package main

import (
	"database/sql"
	"fmt"
)

func CreateTable(db *sql.DB, tableName string) {
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id INT AUTO_INCREMENT,
			name VARCHAR(50) NOT NULL,
			age INT NOT NULL,
			PRIMARY KEY (id)
		)
	`, tableName)
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

// tableを削除する関数
func DropTable(db *sql.DB, tableName string) {
	query := fmt.Sprintf("DROP TABLE %s", tableName)
	_, err := db.Query(query)
	if err != nil {
		panic(err)
	}
}

func InsertUser(db *sql.DB, tableName string, name string, age int) {
	query := fmt.Sprintf("INSERT INTO %s (name, age) VALUES (?, ?)", tableName)
	_, err := db.Exec(query, name, age)
	if err != nil {
		panic(err)
	}
}

func GetUsers(db *sql.DB) error {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return fmt.Errorf("et table: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		rows.Scan(&id, &name, &age)
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
	return nil
}

// Delete a record from the database
func deleteRecord(db *sql.DB, tableName string, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", tableName)
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("deleteRecord: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Deleted %d record(s).\n", rowsAffected)
	return nil
}

// Update a record in the database
func updateRecord(db *sql.DB, tableName string, id int, columnName string, newValue string) error {
	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE id = ?", tableName, columnName)
	result, err := db.Exec(query, newValue, id)
	if err != nil {
		return fmt.Errorf("updateRecord: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Updated %d record(s).\n", rowsAffected)
	return nil
}
