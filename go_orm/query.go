package main

import (
	"database/sql"

	_ "github.com/andrewpillar/query"
)

// Query .
type Query struct {
	db *sql.DB
}

func initDB() (*Query, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	return &Query{db: db}, nil
}

func (Q *Query) Close() {
	Q.db.Close()
}

func (Q *Query) Query() (*sql.Rows, error) {
	return Q.db.Query(
		"SELECT data_type, data_status FROM offline_data_sync_status WHERE ID = 1417",
	)
}
