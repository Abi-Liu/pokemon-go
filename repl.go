package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
