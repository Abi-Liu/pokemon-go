package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *Config) {
	s := bufio.NewScanner(os.Stdin)
	fmt.Println("Type a command!\n type 'help' to get started")
	commands := getCliCommands()
	for {
		fmt.Print(">")
		s.Scan()
		text := s.Text()
		cleanText := cleanInput(text)

		command, ok := commands[cleanText[0]]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.callback(config, cleanText[1:])
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(command string) []string {
	lowercaseText := strings.ToLower(command)
	slice := strings.Fields(lowercaseText)
	return slice
}
