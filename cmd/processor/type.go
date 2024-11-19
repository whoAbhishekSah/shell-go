package processor

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/index"
	"github.com/codecrafters-io/shell-starter-go/utils"
)

const PathEnvVarKeyName = "PATH"

func ProcessTypeCmd(cmd string, args string) {
	if cmd != index.TypeCmd {
		panic(fmt.Sprintf("wrong command, %s", cmd))
	}
	if utils.StringFoundInArray(index.ShellBuiltInCommands, args) {
		fmt.Printf("%s is a shell builtin\n", args)
		return
	}
	cmdFound, pathToCmd := cmdFoundInPath(args)
	if cmdFound {
		fmt.Printf("%s is %s\n", args, pathToCmd)
		return
	}

	fmt.Printf("%s: not found\n", args)
}

func cmdFoundInPath(cmd string) (bool, string) {
	path := os.Getenv(PathEnvVarKeyName)
	splittedPath := strings.Split(path, ":")
	for _, currPath := range splittedPath {
		entries, _ := os.ReadDir(currPath)
		for _, e := range entries {
			if e.Name() == cmd {
				return true, fmt.Sprintf("%s/%s", currPath, e.Name())
			}
		}
	}
	return false, ""
}
