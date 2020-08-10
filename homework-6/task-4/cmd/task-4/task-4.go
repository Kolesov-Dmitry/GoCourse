package main

import (
	"fmt"

	"homework.go/v6/task-4/internal/app/quadratic"
)

func main() {

	if roots, valid := quadratic.Solve(3, -14, -5); valid {
		fmt.Print(roots)
	}
}
