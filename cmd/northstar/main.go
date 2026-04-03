package main

import (
	"fmt"
	"log"
	"os"

	"github.com/harrytripp/north-star/internal/store"
)

func main() {
	fmt.Println("===== Calling local agent =====")

	dir, _ := os.Getwd()
	fmt.Printf("\nRunning in: %s\n", dir)

	//fmt.Println(agents.Response())

	// Initialises database and assigns it to the var "db"
	db, err := store.InitDatabase("./database/journal.db")
	if err != nil {
		log.Fatal(err) // this prints any returned errors and exits
	}

	entry := store.Entry{
		Title: "my title",
		Input: "I jumped the shark.",
		Model: "Ministral-3-8B-Instruct-2512-Q8_0.gguf",
	}

	_, err = db.CreateEntry(&entry)
	if err != nil {
		log.Fatal(err)
	}

	query, err := db.AllEntries()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", query)
}
