package processor

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/index"
)

func ProcessCdCmd(cmd string, args string) {
	if cmd != index.CdCmd {
		panic(fmt.Sprintf("wrong command, %s", cmd))
	}
	
	err := os.Chdir(args)
	if err!=nil {
		fmt.Printf("cd: %s: No such file or directory\n", args)
	}
}