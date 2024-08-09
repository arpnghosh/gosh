package main

import (
	"bufio"
	"fmt"
	"os"
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

		case "ls":
			files, _ := os.ReadDir(".")
			for _, file := range files {
				fmt.Println(file.Name())
			}

		case "cd":
			if len(args) == 1 {
				os.Chdir("/home")
			} else if len(args) == 2 && args[1] == ".." {
				os.Chdir("..")
			} else if len(args) > 2 {
				fmt.Println("cd: too many arguments for the cd command")
			} else {
				dir := strings.Join(args[1:], " ")
				dirInfo, _ := os.Stat(dir)
				if dirInfo.IsDir() {
					os.Chdir(dir)
				} else {
					fmt.Printf("cd: %s is not a directory\n", dir)
				}
			}

		default:
			fmt.Println("gosh: command not found: " + input)
		}
	}
}
