package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println(`
Welcome to the Pokedex!
Usage:

help:Display a help message
exit:Exit the Pokedex

	`)
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func main() {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(`Pokedex >   `)
		scanner.Scan()
		text := scanner.Text()
		val, ok := commands[text]
		if !ok {
			fmt.Println("Not a valid command! Try again!")
			continue
		}
		val.callback()
	}

}
