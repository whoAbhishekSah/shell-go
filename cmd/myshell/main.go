package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err:= bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		if err!=nil {
			panic("error in reading command")
		}
		if !isValidCmd(input) {
			fmt.Printf("%s: command not found\n", input)
		}
	
	}
}

func isValidCmd(_ string) bool {
	return false
}
