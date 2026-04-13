package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/harrytripp/north-star/internal/store"
)

func must(v interface{}, err error) interface{} {
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("===== NORTH STAR =====")

	dir, _ := os.Getwd()
	fmt.Printf("\nRunning in: %s\n", dir)

	// Initialises database and assigns it to the var "db"
	db := must(store.InitDatabase("./database/journal.db")).(*store.Store)

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Add manual entry")
		fmt.Println("2. Add test entry")
		fmt.Println("3. View database")
		fmt.Println("0. Exit")
		fmt.Print("Choose: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			// Take use input, save struct output as variable
			entry := must(db.Input()).(*store.Entry)
			// Pass user input and add to database
			must(db.CreateEntry(entry))

			fmt.Println("Entry added: ", entry)

		case "2":
			test := store.Entry{
				Title: "Radical",
				Input: "I jumped the shark.",
			}

			_ = must(db.CreateEntry(&test))
			fmt.Println("Test entry added")

		case "3":
			entries := must(db.AllEntries()).([]store.Entry)
			for _, e := range entries {
				fmt.Println(e)
			}

		case "0":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

/*
	fmt.Println(agents.Response())

	entry, err := db.Input()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.CreateEntry(entry)
	if err != nil {
		log.Fatal(err)
	}

	query, err := db.AllEntries()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", query)

	query, err = db.EntryByModel("gemma-4")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", query)
*/
