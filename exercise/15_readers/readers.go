package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CmdHello   = "hello"
	CmdGoodbye = "bye"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numLines := 0
	numCommands := 0
	for scanner.Scan() {
		if strings.ToUpper(scanner.Text()) == "Q" {
			break
		} else {
			text := strings.TrimSpace(scanner.Text())
			switch text {
			case CmdHello:
				numCommands += 1
				fmt.Println("Command response: hi")
			case CmdGoodbye:
				numCommands += 1
				fmt.Println("Command response: bye")
			}
			if text != "" {
				numLines += 1
			}
		}
	}

	fmt.Printf("You entered %v lines\n", numLines)
	fmt.Printf("You entered %v commands\n", numCommands)
}
