package main

import (
	"fmt"
	"homework.com/v6/task-1/internal/app/sum"
)

func main() {
	result := sum.Sum([]float64{1, 2, 3})
	fmt.Println(result)
}
