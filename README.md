# my-ls

**my-ls** is a Go implementation of the Unix `ls` command, which lists files and directories within a specified path. If no path is provided, it defaults to the current directory. This project mimics the behavior of the original `ls` command with additional functionality as specified below.

## Features

The behavior of **my-ls** is designed to closely match the original `ls` command with the following flags supported:

- **-l**: Long listing format, displaying additional file details such as permissions, owner, and size. The output must match the system's `ls -l` exactly.
- **-R**: Recursively list directories and their contents.
- **-a**: Include hidden files (those starting with `.`).
- **-r**: Reverse the order of the output.
- **-t**: Sort files by modification time, newest first.

You can combine these flags in the same way they are used in the original `ls` command.

## Installation

Clone the repository and build the project using Go:

```bash
git clone <repository-url>
cd my-ls
go build -o my-ls

./my-ls              # List files in the current directory
./my-ls -l           # Long format
./my-ls -a           # Include hidden files
./my-ls -R           # Recursively list subdirectories
./my-ls -r           # Reverse order
./my-ls -t           # Sort by modification time
./my-ls -laRt        # Combine flags


