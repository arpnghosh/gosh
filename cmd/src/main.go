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
	amp := make(map[string]string)

	for {
		fmt.Println(CustomPrompt())

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

		case "alias": // currently works for only external commands

			if len(args) == 1 {
				fmt.Println("Usage: alias name = command")
			} else if len(args) == 2 && args[1] == "list" {
				if len(amp) == 0 {
					fmt.Println("No aliases defined")
				} else {
					for key, value := range amp {
						fmt.Printf("Alias: %s Command: %s \n", key, value)
					}
				}
			} else if len(args) >= 4 && args[2] == "=" {
				alias_name := strings.Join(args[1:2], " ")
				alias_command := strings.Join(args[3:], " ")
				amp[alias_name] = alias_command
				fmt.Printf("Command: %v set to Alias: %v \n", alias_command, alias_name)
			} else {
				fmt.Println("Usage: alias name = command")
			}

		case "unalias": // unaliasing support
			nameOfAlias := strings.Join(args[1:2], " ")
			if len(args) == 1 {
				fmt.Println("Usage: unalias name")
			} else if len(args) >= 3 {
				fmt.Println("Delete one alias name at a time")
			} else if _, exists := amp[nameOfAlias]; exists {
				delete(amp, nameOfAlias)
				fmt.Printf("Deleted alias: %s \n", nameOfAlias)
			} else {
				fmt.Println("Alias does not exist")
			}

		default:
			if aliasCmd, exists := amp[command]; exists {
				args = strings.Fields(aliasCmd)
				command = args[0]
			}
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
