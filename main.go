package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/syedazeez337/Goshell/cmd"
)

func main() {
	// read the input from the user
	for {
		fmt.Print("$ ")

		// Read the user input
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			continue
		}

		input = strings.TrimSpace(input)

		// when user enters empty string it continues
		if input == "" {
			continue
		}

		// Exiting the shell environment
		if input == "Exit" {
			break
		}

		args := parseInput(input)
		cmd.ExecuteCommand(args)
		fmt.Println("You entered:", args)
	}

	// goodbye message
	fmt.Println("Goodbye, visit again...")
}

// parsing input
func parseInput(input string) []string {
	return strings.Fields(input)
}
