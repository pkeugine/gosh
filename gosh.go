package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to gosh, the shell written in Go-lang")
	fmt.Println("by: pkeugine (Eugine Park)")
	fmt.Println("---------------------------------------------")
	for {
		fmt.Print(">>> ")
		command, _ := reader.ReadString('\n')
		command = strings.Replace(command, "\n", "", -1)
		if strings.Compare("help", command) == 0 {
			fmt.Println("Commands:\n" +
				"help: Display available commands\n" +
				"exit: Exit out of shell")
		} else if strings.Compare("exit", command) == 0 {
			break
		}

	}
}
