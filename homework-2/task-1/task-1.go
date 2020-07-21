package main

import "fmt"

// isEvenr checks if 'number' is even number
// Input:
//  - 'number' is a signed integer number
// Output:
// - True if number is even 
// - False if number is odd
func isEven(number int) bool {
	return number % 2 == 0
}

func main() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if isEven(number) {
		fmt.Printf("Число %d чётное", number)
	} else {
		fmt.Printf("Число %d нечётное", number)
	}
}