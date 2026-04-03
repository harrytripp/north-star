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
	db, initErr := store.InitDatabase("./database/journal.db")
	if initErr != nil {
		log.Fatal(initErr) // this prints any returned errors and exits
	}

	entry := store.Entry{
		Title: "my title",
		Model: "Ministral-3-8B-Instruct-2512-Q8_0.gguf",
	}

	_, creaErr := db.CreateEntry(&entry)
	if creaErr != nil {
		log.Fatal(creaErr)
	}

}
