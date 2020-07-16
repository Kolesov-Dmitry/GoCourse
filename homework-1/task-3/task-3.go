package main

import (
	"fmt"
	"math"
)

func main() {
	
	var deposit, interestRate float64
	fmt.Print("Введите сумму вклада: ")
	fmt.Scan(&deposit)
	fmt.Print("Введите процентную ставку: ")
	fmt.Scan(&interestRate)

	// To calculate the amount of money that wil be on the account after 5 years
	// we need to multiply initial value of the deposit with value (1 + interestRate / 100.0) five times
	// The value (1 + interestRate / 100.0) is the deposit increment for one year
	var total = deposit * math.Pow((1 + interestRate / 100.0), 5)
	
	fmt.Print("\nСумма вклада через 5 лет: ", total)
}