package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "mysql",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "luxcare_demo",
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
