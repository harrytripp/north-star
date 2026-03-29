package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world")

	dir, _ := os.Getwd()
	fmt.Printf("Running in: %s\n", dir)
}
