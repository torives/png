package main

import (
	"fmt"
	"log"
	"os"

	"github.com/torives/png/repo"

	_ "modernc.org/sqlite"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("wrong syntax")
		os.Exit(1)
	}

	category := args[1]
	db, err := repo.Open("./png.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch args[0] {
	case "add":
		err = db.AddCategory(category)
		if err != nil {
			fmt.Printf("failed to add category. %s\n", err)
			os.Exit(1)
		} else {
			fmt.Printf(`category "%s" was added`, category)
		}
	case "-p":
		id, err := db.NextProjectId(category)
		if err != nil {
			fmt.Printf("failed to get project id. %s\n", err)
			os.Exit(1)
		} else {
			fmt.Printf("project id: %d\n", id)
		}
	default:
		fmt.Println("wrong syntax")
		os.Exit(1)
	}
}
