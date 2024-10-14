package main

import (
	"database/sql"

	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	fmt.Println("Hello world")
	db, err := sql.Open("sqlite", "./png.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok")
	defer db.Close()
}
