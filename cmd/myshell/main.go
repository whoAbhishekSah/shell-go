package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const ExitCmd = "exit 0"

func main() {
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		if err != nil {
			panic("error in reading command")
		}
		
		if input == ExitCmd {
			os.Exit(0)
		}

		if !isValidCmd(input) {
			fmt.Printf("%s: command not found\n", input)
		}
	}
}

func isValidCmd(_ string) bool {
	return false
}
