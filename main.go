package main

import (
	"fmt"
	"os"

	"github.com/rammanokar/pgtest/cmd"
	"github.com/rammanokar/pgtest/config"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		fmt.Printf("Warning: Error loading config: %v\n", err)
		fmt.Println("Proceeding with default or command-line provided values.")
	}

	if err := cmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
