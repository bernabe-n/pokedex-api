package main

import (
	"fmt"
)

func callbackHelp(cfg *config) error{
	fmt.Println("Here are your available commands:")

	for _, command := range getCommands() {
		fmt.Printf("- %s: %s\n", command.name, command.description)
	}

	return nil
}