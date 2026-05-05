package main

import (
	"fmt"
	"errors"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Please provide exactly one argument: the name of the location area.")
	}

	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return fmt.Errorf("Error fetching location areas: %v", err)
	}
	fmt.Printf("Pokemon in %s:\n", locationAreaName)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}