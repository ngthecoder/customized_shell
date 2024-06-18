package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ngthecoder/customized_shell/internal"
)

func main() {
	paths, home := internal.LoadEnv()
	for {
		fmt.Print("$ ")

		inputWithDelimiter, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		command := internal.ParseInput(strings.TrimSpace(inputWithDelimiter))

		switch command.Name {
		case "exit":
			internal.HandleExit(command)
		case "echo":
			internal.HandleEcho(command)
		case "type":
			internal.HandleType(command, paths)
		case "pwd":
			internal.HandlePwd()
		case "cd":
			internal.HandleCd(command, home)
		default:
			internal.HandleExternalCommand(command)
		}
	}
}
