package main

import (
	"fmt"
	"os"

	"github.com/harrytripp/north-star/internal/agents"
)

func main() {
	fmt.Println("===== Calling local agent =====")

	dir, _ := os.Getwd()
	fmt.Printf("\nRunning in: %s\n", dir)

	// agents.Response()
	fmt.Println(agents.Response())
}
