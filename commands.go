package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/abi-liu/pokedexcli/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

type locationAreaConfig struct {
	nextUrl string
	prevUrl string
}

func getCliCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 locations inside the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 locations inside the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore the pokemon in a certain area. Ex: explore pastoria-city-area",
			callback:    commandExpolore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon. Ex: catch pikachu",
			callback:    commandCatch,
		},
	}
}

func commandExit(config *Config, opts []string) error {
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, opts []string) error {
	for k, v := range getCliCommands() {
		fmt.Printf("%s - %s\n", k, v.description)
	}
	return nil
}

func commandMap(config *Config, opts []string) error {
	data, err := api.GetLocationData(config.Client, config.Cache, config.Next)
	if err != nil {
		return err
	}
	config.Next = data.Next
	config.Previous = data.Previous

	for i, location := range data.Results {
		fmt.Printf("%v - %v\n", i+1, location.Name)
	}

	return nil
}

func commandMapb(config *Config, opts []string) error {
	if config.Previous == nil {
		return errors.New("You are at the first location area. Cannot go back!")
	}

	data, err := api.GetLocationData(config.Client, config.Cache, config.Previous)
	if err != nil {
		return err
	}

	config.Next = data.Next
	config.Previous = data.Previous

	for i, location := range data.Results {
		fmt.Printf("%v - %v\n", i, location.Name)
	}

	return nil
}

func commandExpolore(config *Config, opts []string) error {
	if len(opts) == 0 {
		return errors.New("Please provide a location to explore")
	}

	fmt.Println("Exploring " + opts[0] + "...")

	data, err := api.ExploreArea(config.Client, config.Cache, opts[0])

	if err != nil {
		return err
	}

	fmt.Println("Found pokemon:")
	for _, p := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *Config, opts []string) error {
	if len(opts) == 0 {
		return errors.New("Please provide a pokemon name to catch")
	}

	fmt.Printf("Catching %s...\n", opts[0])

	data, err := api.GetPokemonInformation(config.Client, config.Cache, opts[0])
	if err != nil {
		return err
	}

	catchThreshold := 50
	probability := rand.Intn(data.BaseExperience)

	if probability <= catchThreshold {
		config.Pokedex[opts[0]] = data
		fmt.Printf("You have sucessfully caught %s!\n", opts[0])
	} else {
		fmt.Printf("%s has fled!\n", opts[0])
	}

	return nil
}
