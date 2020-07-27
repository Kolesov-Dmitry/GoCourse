package main

// First things first!
// To have an apportunity to implement 'help function' for a calculator we need to implement the calculator itself

// Step One - implement the calculator

import (
	"bufio"
	"fmt"
	"os"

	"homework.com/v4/task-3/internal/app/expression"
)

// getUserInput gets user input from stdin
func getUserInput() string {
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func main() {
	solver := expression.NewSolver()
	for {
		userInput := getUserInput()
		if userInput == "exit" {
			break
		} else {
			result, err := solver.Solve(userInput)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		}
	}
}
