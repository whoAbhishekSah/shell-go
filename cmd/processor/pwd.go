package processor

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/index"
)

func ProcessPwdCmd(cmd string, args string) {
	if cmd != index.PwdCmd {
		panic(fmt.Sprintf("wrong command, %s", cmd))
	}
	
	currDir, _ := os.Getwd() 
	fmt.Println(currDir)
}
