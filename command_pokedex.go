package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	pokemons := cfg.caughtPokemon
	if pokemons == nil {
		return fmt.Errorf("no pokemons in pokedex yet")
	}
	fmt.Println("Pokemon in Pokedex:")
	for _, pokemon := range pokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
