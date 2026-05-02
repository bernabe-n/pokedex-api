package main

import (
	"github.com/bernabe-n/pokedex-api/internal/pokeapi";
	"fmt";
	"log")

func callbackMap()error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatalf("Error fetching location areas: %v", err)
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	return nil
}