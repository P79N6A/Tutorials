package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "elemeuser:xBe6pXzy7r@alops1-mysql-20.db.elenet.me:3306/push_knight")
	defer db.Close()
	return db, err
}
