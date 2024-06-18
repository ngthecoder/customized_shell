package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")

		inputWithDelimiter, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		input := strings.Trim(inputWithDelimiter, "\n")

		fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
	}
}
