package main

import (
	"fmt"
	"log"

	"github.com/torives/png/repo"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := repo.Open("./png.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok")
	defer db.Close()
}
