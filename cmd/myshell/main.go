package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	cmd "github.com/codecrafters-io/shell-starter-go/cmd/index"
	"github.com/codecrafters-io/shell-starter-go/cmd/processor"
	"github.com/codecrafters-io/shell-starter-go/utils"
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

		splittedCmd := strings.Split(input, " ")
		command := splittedCmd[0]
		if !isValidCmd(command) {
			fmt.Printf("%s: command not found\n", command)
		}

		// skip first item and join the values to get args str
		args := strings.Join(splittedCmd[1:], " ")
		processCmd(command, args)

	}
}

func isValidCmd(inputCmd string) bool {
	validCommands := cmd.ValidCommands
	return utils.StringFoundInArray(validCommands, inputCmd)
}

func processCmd(inputCmd string, args string) {
	switch inputCmd {
	case cmd.ExitCmd:
		if args == "0"{
			os.Exit(0)
		}
	case cmd.EchoCmd:
		fmt.Println(args)
	case cmd.TypeCmd:
		processor.ProcessTypeCmd(inputCmd, args)
	default:
		return
	}
}
