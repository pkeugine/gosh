package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to gosh, the shell written in Go-lang")
	fmt.Println("by: pkeugine (Eugine Park)")
	fmt.Println("---------------------------------------------")
	for {
		fmt.Print("gosh> ")
		command, _ := reader.ReadString('\n')
		command = strings.Replace(command, "\n", "", -1)
		if strings.Compare("help", command) == 0 {
			cmd := exec.Command("./src/commands/help")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
			//fmt.Println("Commands:\n" +
			//	"   help: Display available commands\n" +
			//	"   exit: Exit out of shell\n" +
			//	"listall: List all files and directories including hidden ones")
		} else if strings.Compare("exit", command) == 0 {
			break
		} else if strings.Compare("listall", command) == 0 {
			cmd := exec.Command("ls", "-la")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
		}

	}
}
