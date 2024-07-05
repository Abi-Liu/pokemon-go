package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/abi-liu/pokedexcli/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
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
	}
}

func commandExit(config *Config) error {
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	for k, v := range getCliCommands() {
		fmt.Printf("%s - %s\n", k, v.description)
	}
	return nil
}

func commandMap(config *Config) error {
	data, err := api.GetLocationData(config.Client, config.Next)
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

func commandMapb(config *Config) error {
	if config.Previous == nil {
		return errors.New("You are at the first location area. Cannot go back!")
	}

	data, err := api.GetLocationData(config.Client, config.Previous)
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
