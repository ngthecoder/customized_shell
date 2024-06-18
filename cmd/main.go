package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func main() {
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
				fmt.Println("exit: wrong number of arguments")
			} else {
				fmt.Println("exit: invalid argument")
			}
		case "echo":
			fmt.Println(strings.Join(command.Args, " "))
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command.Name)
		}
	}
}

func parseInput(input string) Command {
	parts := strings.Split(input, " ")

	return Command{
		Name: parts[0],
		Args: parts[1:],
	}
}
