package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func HandleExternalCommand(command Command) {
	cmd := exec.Command(command.Name, command.Args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: command not found\n", command.Name)
	}
}
