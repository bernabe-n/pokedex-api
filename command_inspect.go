package main

import (
	"fmt"
	"errors"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided. Please provide the name of the pokemon you want to inspect.")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		fmt.Printf("%s not found in your collection.\n", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("Type: %s\n", typeInfo.Type.Name)
	}

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-  %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	for _, moveInfo := range pokemon.Moves {
		fmt.Printf("Move: %s\n", moveInfo.Move.Name)
	}

	return nil
}