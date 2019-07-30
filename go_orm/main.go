package main

import (
	"fmt"
	"os"
)

func main() {
	conn, err := initDB()
	if err != nil {
		fmt.Printf("Connect...Error!\n")
		os.Exit(1)
	}

	defer conn.Close()
}
