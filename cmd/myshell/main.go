package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	ExitCmd = "exit 0"
	EchoCmd = "echo"
)
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

		splittedCmd := strings.Split(input, " ")
		cmd := splittedCmd[0]
		if !isValidCmd(cmd) {
			fmt.Printf("%s: command not found\n", cmd)
		}
		
		// skip first item and join the values to get args str
		args := strings.Join(splittedCmd[1:], " ")
		processCmd(cmd, args)

	}
}

func isValidCmd(cmd string) bool {
	validCommands := []string{"echo"}
	return stringFoundInArray(validCommands, cmd)
}

func stringFoundInArray(ary []string, stringToSearch string) bool {
	for _, item := range ary {
		if item == stringToSearch {
			return true
		}
	}
	return false
}

func processCmd(inputCmd string, args string) {
	switch inputCmd {
	case EchoCmd:
		fmt.Println(args)
	default:
		return
	}
}
