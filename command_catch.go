package main

import (
	"fmt"
	"errors"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided. Please provide the name of the pokemon you want to catch.")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("Error fetching pokemon: %v", err)
	}

	const threshold = 1000

	randomNumber := rand.Intn(pokemon.BaseExperience) // Random number between 0 and BaseExperience
	//fmt.Println(pokemon.Name, pokemon.BaseExperience, randomNumber, threshold)
	if randomNumber > threshold {
		fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	cfg.caughtPokemon[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)

	return nil
}