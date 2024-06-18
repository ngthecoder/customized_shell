package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func HandleExit(command Command) {
	if len(command.Args) == 1 && command.Args[0] == "0" {
		os.Exit(0)
	} else if len(command.Args) != 1 {
		fmt.Fprintln(os.Stderr, "exit: wrong number of arguments")
	} else {
		fmt.Fprintln(os.Stderr, "exit: invalid argument")
	}
}

func HandleEcho(command Command) {
	fmt.Fprintln(os.Stdout, strings.Join(command.Args, " "))
}

func HandlePwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(dir)
}

func HandleCd(command Command, home string) {
	if len(command.Args) == 0 {
		moveToDir(home)
	} else if len(command.Args) == 1 {
		moveToDir(command.Args[0])
	} else {
		fmt.Fprintln(os.Stderr, "cd: wrong number of arguments")
	}
}

func moveToDir(path string) {
	if path == "~" {
		path = os.Getenv("HOME")
	}
	err := os.Chdir(filepath.Clean(path))
	if err != nil {
		fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", path)
	}
}

func HandleType(command Command, paths []string) {
	if len(command.Args) != 1 {
		fmt.Fprintln(os.Stderr, "type: invalid number of arguments")
		return
	}
	commandToSearch := command.Args[0]
	cmds := []string{"exit", "echo", "type", "pwd", "cd"}
	for _, cmd := range cmds {
		if commandToSearch == cmd {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", commandToSearch)
			return
		}
	}
	for _, path := range paths {
		fullPath := filepath.Join(path, commandToSearch)
		if _, err := os.Stat(fullPath); err == nil {
			fmt.Fprintf(os.Stdout, "%s is %s\n", commandToSearch, fullPath)
			return
		}
	}
	fmt.Fprintf(os.Stdout, "%s: not found\n", commandToSearch)
}
