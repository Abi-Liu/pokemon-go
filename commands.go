package main

import (
	"fmt"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
			description: "Displays the names of 20 locations inside the Pokemon world. Each additional call will display the next 20 locations",
			callback:    commandMap,
		},
	}
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandHelp() error {
	for k, v := range getCliCommands() {
		fmt.Printf("%s - %s\n", k, v.description)
	}
	return nil
}

func commandMap() error {
	res, err := http.Get(BASE_URL + "/location-area/")
}
