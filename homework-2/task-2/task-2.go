package main

import "fmt"

// isDevisibleByThree checks if 'number' is divisible by three without remainder
// Input:
//  - 'number' is a signed integer number
// Output:
// - True if number is divisible by three without remainder
// - False otherwise
func isDivisibleByThree(number int) bool {
	return number % 3 == 0
}

func main() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if isDivisibleByThree(number) {
		fmt.Printf("Число %d делится на 3 без остатка", number)
	} else {
		fmt.Printf("Число %d не делится на 3 без остатка", number)
	}
}