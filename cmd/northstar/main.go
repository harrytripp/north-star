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

	_, err := store.InitDatabase("./database/journal.db")
	if err != nil {
		log.Fatal(err) // this prints and exits
	}
}
