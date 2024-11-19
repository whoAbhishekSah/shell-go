package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	cmd "github.com/codecrafters-io/shell-starter-go/cmd/index"
	"github.com/codecrafters-io/shell-starter-go/cmd/processor"
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

		// skip first item and join the values to get args str
		args := strings.Join(splittedCmd[1:], " ")
		processCmd(command, args)

	}
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
	case cmd.PwdCmd:
		processor.ProcessPwdCmd(inputCmd, args)
	case cmd.CdCmd:
		processor.ProcessCdCmd(inputCmd, args)
	default:
		processor.Execute(inputCmd, args)
	}
}
