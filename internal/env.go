package internal

import (
	"os"
	"strings"
)

func LoadEnv() ([]string, string) {
	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")
	home := os.Getenv("HOME")
	return paths, home
}
