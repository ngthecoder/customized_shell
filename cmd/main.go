package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func main() {
	paths, home := loadEnv()
	for {
		fmt.Print("$ ")

		inputWithDelimiter, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		command := parseInput(strings.TrimSpace(inputWithDelimiter))

		switch command.Name {
		case "exit":
			handleExit(command)
		case "echo":
			handleEcho(command)
		case "type":
			handleType(command, paths)
		case "pwd":
			handlePwd()
		case "cd":
			handleCd(command, home)
		default:
			handleExternalCommand(command)
		}
	}
}

func loadEnv() ([]string, string) {
	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")
	home := os.Getenv("HOME")
	return paths, home
}

func parseInput(input string) Command {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return Command{}
	}
	return Command{
		Name: parts[0],
		Args: parts[1:],
	}
}

func handleExit(command Command) {
	if len(command.Args) == 1 && command.Args[0] == "0" {
		os.Exit(0)
	} else if len(command.Args) != 1 {
		fmt.Fprintln(os.Stderr, "exit: wrong number of arguments")
	} else {
		fmt.Fprintln(os.Stderr, "exit: invalid argument")
	}
}

func handleEcho(command Command) {
	fmt.Fprintln(os.Stdout, strings.Join(command.Args, " "))
}

func handlePwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(dir)
}

func handleCd(command Command, home string) {
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

func handleType(command Command, paths []string) {
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

func handleExternalCommand(command Command) {
	cmd := exec.Command(command.Name, command.Args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: command not found\n", command.Name)
	}
}
