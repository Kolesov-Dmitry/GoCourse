package main

// First things first!
// To have an apportunity to implement 'help function' for a calculator we need to implement the calculator itself

// Step One - implement the calculator
// Step Two - help function

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

// printHelp prints usage datails
func printHelp() {
	fmt.Println("List of supported operations:")
	fmt.Println("   '+' - plus operator")
	fmt.Println("   '-' - minus operator")
	fmt.Println("   '*' - multiply operator")
	fmt.Println("   '/' - division operator")
	fmt.Println("Type 'exit' to quit.")
}

func main() {
	var solver expression.Solver
	for {
		userInput := getUserInput()
		if userInput == "exit" {
			break
		} else if userInput == "help" {
			printHelp()
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
