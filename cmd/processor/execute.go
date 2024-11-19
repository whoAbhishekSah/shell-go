package processor

import (
	"fmt"
	"os/exec"
)

func Execute(userCmd string, args string) {
	cmdFound, pathToCmd := cmdFoundInPath(userCmd)
	if cmdFound {
		cmd := exec.Command(pathToCmd, args)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		fmt.Printf("%s", string(stdout))
		return
	}
	fmt.Printf("%s: command not found\n", userCmd)
}
