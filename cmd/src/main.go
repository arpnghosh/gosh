package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprint(os.Stderr, "Error reading input:", err)
		}
		input = strings.TrimSpace(input)

		if input == "exit" {
			os.Exit(0)
		}
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}
		command := args[0]
		switch command {

		case "echo":
			fmt.Println(strings.Join(args[1:], " "))

		case "pwd":
			fmt.Println(os.Getwd())

		case "cd":
			if len(args) == 1 {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error getting user's home directory", err)
				}
				os.Chdir(homeDir)
			} else if len(args) == 2 && args[1] == ".." {
				os.Chdir("..")
			} else if len(args) > 2 {
				fmt.Println("cd: too many arguments for the cd command")
			} else {
				dir := strings.Join(args[1:], " ")
				dirInfo, err := os.Stat(dir)
				if err != nil {
					fmt.Printf("cd: directory %s does not exist\n", dir)
				} else if dirInfo.IsDir() {
					os.Chdir(dir)
				} else {
					fmt.Printf("cd: %s is not a directory\n", dir)
				}
			}

		default:
			cmd := exec.Command(command, args[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("gosh: command not found: " + input)
			}
		}
	}
}
