# Customized Shell with Go

`customized_shell` is a simple command-line shell implemented in Go. It supports basic built-in commands like `cd`, `echo`, `pwd`, `type`, and `exit`, as well as executing external commands available in the system's `PATH`.

## Features

- Built-in commands: `cd`, `echo`, `pwd`, `type`, and `exit`
- Executes external commands found in the system's `PATH`
- Handles user input interactively
- Provides error messages for invalid commands and arguments

## Usage

### Building and Running the Shell

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/customized_shell.git
    cd customized_shell
    ```

2. Build and run the shell using the provided script:
    ```sh
    ./customized_shell.sh
    ```

### Commands

- **cd [directory]**: Change the current directory to the specified directory. Use `~` to navigate to the home directory.
- **echo [args...]**: Print the arguments to the standard output.
- **pwd**: Print the current working directory.
- **type [command]**: Display whether the command is a shell builtin or an external command, and show the path for external commands.
- **exit [0]**: Exit the shell. Only `exit 0` is considered a valid exit command.

### Example Usage

```sh
$ cd ~
$ pwd
/home/user
$ echo Hello, World!
Hello, World!
$ type cd
cd is a shell builtin
$ ls -l
total 8
drwxr-xr-x 2 user user 4096 Jun 18 15:07 cmd
drwxr-xr-x 2 user user 4096 Jun 18 15:07 internal
-rwxr-xr-x 1 user user  102 Jun 18 15:07 customized_shell.sh
-rw-r--r-- 1 user user  258 Jun 18 15:07 go.mod
$ exit 0
```

## Development

### Prerequisites

- Go 1.16 or later

### Building from Source

To build the project from source without using the provided script:

```sh
cd cmd
go build -o ../customized_shell
../customized_shell
```

## Testing

To run the automated tests, use the following command:

```sh
go test ./...
```

To see test coverage, use:

```sh
go test -cover ./...
```

By adding automated tests, you can ensure that your code behaves as expected and catch any regressions early in the development process.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with any improvements or bug fixes.