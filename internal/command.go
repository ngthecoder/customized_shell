package internal

import "strings"

type Command struct {
	Name string
	Args []string
}

func ParseInput(input string) Command {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return Command{}
	}
	return Command{
		Name: parts[0],
		Args: parts[1:],
	}
}
