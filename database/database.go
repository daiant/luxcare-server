package database

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    os.Getenv("MYSQL_NET"),
		Addr:   os.Getenv("MYSQL_ADDR"),
		DBName: os.Getenv("MYSQL_DBNAME"),
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
