package internal

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHandlePwd(t *testing.T) {
	originalDir, _ := os.Getwd()

	output := captureOutput(func() {
		HandlePwd()
	})

	output = strings.TrimSpace(output)

	if output != originalDir {
		t.Errorf("Expected %s, got %s", originalDir, output)
	}
}

func TestHandleCd(t *testing.T) {
	home := os.Getenv("HOME")
	originalDir, _ := os.Getwd()

	tests := []struct {
		path     string
		expected string
	}{
		{path: home, expected: home},
		{path: ".", expected: originalDir},
		{path: "..", expected: filepath.Dir(originalDir)},
	}

	for _, test := range tests {
		if err := os.Chdir(originalDir); err != nil {
			t.Fatalf("could not change to original directory: %v", err)
		}

		moveToDir(test.path)
		dir, _ := os.Getwd()
		if dir != filepath.Clean(test.expected) {
			t.Errorf("moveToDir(%s) = %s; want %s", test.path, dir, test.expected)
		}
	}

	os.Chdir(originalDir)
}

func TestHandleEcho(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{args: []string{"Hello", "World"}, expected: "Hello World\n"},
		{args: []string{"Go", "Testing"}, expected: "Go Testing\n"},
	}

	for _, test := range tests {
		command := Command{Name: "echo", Args: test.args}
		output := captureOutput(func() {
			HandleEcho(command)
		})
		if output != test.expected {
			t.Errorf("HandleEcho(%v) = %s; want %s", test.args, output, test.expected)
		}
	}
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f()
	w.Close()
	os.Stdout = old
	return <-outC
}
