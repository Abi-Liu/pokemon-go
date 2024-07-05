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

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		s.Scan()
		text := s.Text()
		command := cleanInput(text)
		fmt.Print(command)
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
	return nil
}

func commandHelp() error {
	for k, v := range getCliCommands() {
		fmt.Printf("%s - %s\n", k, v.description)
	}
	return nil
}
