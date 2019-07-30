package main

import (
	"fmt"
	"log"
	"testing"
)

func TestQuery(T *testing.T) {
	conn, _ := initDB()
	rows, _ := conn.Query()
	defer rows.Close()

	for rows.Next() {
		var data_type int
		var date_status int
		if err := rows.Scan(&data_type, &date_status); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("data_type:%d ,date_status:is %d\n", data_type, date_status)
	}
}
