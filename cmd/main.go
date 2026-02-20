package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("AgentMigrate starting...")
	if err := initialize(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("AgentMigrate ready")
}

func initialize() error {
	return nil
}
