package main

import (
	"fmt"
	"os"

	"github.com/harrytripp/north-star/internal/agents"
)

func main() {
	fmt.Println("Hello, world")

	dir, _ := os.Getwd()
	fmt.Printf("Running in: %s\n", dir)

	agents.Response()
}
