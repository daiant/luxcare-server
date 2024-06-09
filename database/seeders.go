package database

import (
	"database/sql"
	"log"
)

func GenerateMigrations() {
	db := Connect()
	exec(db, `CREATE TABLE contacts (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name TEXT,
		email TEXT,
		phone TEXT,
		comments TEXT,
		created_at BIGINT
	)`)
	defer db.Close()
}
func exec(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
