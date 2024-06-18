package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
			fmt.Println(err)
			return
		}

		command := parseInput(strings.Trim(inputWithDelimiter, "\n"))

		switch command.Name {
		case "exit":
			if len(command.Args) == 1 && command.Args[0] == "0" {
				return
			} else if len(command.Args) != 1 {
				fmt.Fprintln(os.Stdout, "exit: wrong number of arguments")
			} else {
				fmt.Fprintln(os.Stdout, "exit: invalid argument")
			}
		case "echo":
			fmt.Fprintln(os.Stdout, strings.Join(command.Args, " "))
		case "type":
			showCommandType(command.Args[0], paths)
		case "pwd":
			showCurrentDir()
		case "cd":
			if len(command.Args) == 0 {
				moveToDir(home, home)
			} else if len(command.Args) == 1 {
				moveToDir(command.Args[0], home)
			} else {
				fmt.Fprintln(os.Stdout, "cd: wrong number of arguments")
			}
		default:
			cmd := exec.Command(command.Name, command.Args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", command.Name)
			}
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
	parts := strings.Split(input, " ")

	return Command{
		Name: parts[0],
		Args: parts[1:],
	}
}

func showCurrentDir() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dir)
}

func moveToDir(path string, home string) {
	if path == "~" {
		path = home
	}
	err := os.Chdir(path)
	if err != nil {
		fmt.Println(err)
	}
}

func showCommandType(commandToSearch string, paths []string) {
	cmds := []string{"exit", "echo", "type", "pwd", "cd"}
	for _, cmd := range cmds {
		if commandToSearch == cmd {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", commandToSearch)
			return
		}
	}
	for _, path := range paths {
		if _, err := os.Stat(path + "/" + commandToSearch); err == nil {
			fmt.Fprintf(os.Stdout, "%s is %s/%s\n", commandToSearch, path, commandToSearch)
			return
		}
	}
	fmt.Fprintf(os.Stdout, "%s: not found\n", commandToSearch)
}
