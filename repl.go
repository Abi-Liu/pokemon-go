package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Println("Type a command!\n type 'help' to get started")
	commands := getCliCommands()
	for {
		fmt.Print(">")
		s.Scan()
		text := s.Text()
		cleanText := cleanInput(text)

		command, ok := commands[cleanText]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		command.callback()
	}
}

func cleanInput(command string) string {
	lowercaseText := strings.ToLower(command)
	slice := strings.Fields(lowercaseText)
	return strings.Join(slice, " ")
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
