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
		command := args[0]
		switch command {
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		default:
			fmt.Println(input + ": command not found")

		}
	}
}
