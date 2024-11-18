package processor

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/cmd/index"
	"github.com/codecrafters-io/shell-starter-go/utils"
)

const TypeCmd = "type"

func ProcessTypeCmd(cmd string, args string) {
	if cmd != TypeCmd {
		panic(fmt.Sprintf("wrong command, %s", cmd))
	}
	if utils.StringFoundInArray(index.ShellBuiltInCommands, args) {
		fmt.Printf("%s is a shell builtin\n", args)
		return
	}
	fmt.Printf("%s: not found\n", args)
}
