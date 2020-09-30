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
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		if strings.Compare("help", input) == 0 {
			fmt.Println("Commands:\nhelp: Display available commands")
		} else {
			fmt.Println(input)
		}

	}
}
