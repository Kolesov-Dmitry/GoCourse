package main

import (
	"fmt"
	"math/big"
)

// In this task we have to use math/big package 
// due the fact that already 95th Fibbonaci number doesn't fit into 64 bits

// calculateFibbonaci calculates the fisrt 'amount' primes Fibbonaci numbers
// Input:
//  - amount is the amount of Fibbonaci numbers to calculate
// Output:
//  - slice of Fibbonaci numbers 
func calculateFibbonaci(amount uint) []*big.Int {

	if amount < 2 {
		return []*big.Int { big.NewInt(0) }
	}

	var output = make([]*big.Int, 0, amount)
	output = append(output, big.NewInt(0), big.NewInt(1)) // the first two Fibbonaci numbers are 0 and 1
	
	// calculates 'amount' numbers
	for length := uint(len(output)); length < amount; length++ {
		var nextNumber = big.NewInt(0)
		nextNumber = nextNumber.Add(output[length - 2], output[length - 1])

		output = append(output, nextNumber)
	}

	return output
}

func main() {
	
	for idx, number := range calculateFibbonaci(100) {
		fmt.Printf("%03d - %v\n", idx + 1, number.String())
	}

}