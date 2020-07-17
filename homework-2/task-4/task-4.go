package main

import "fmt"

// The alghoritm is based on the folowing statements
//  - the first prime number is 2
//  - the rest primes are odd numbers
//  - to find out if a number is prime it's enought to check 
//    the rest of division only with already found primes

// findPrimeNumbers finds the fisrt amount primes
// Input:
//  - amount is the amount of primes to find
// Output:
//  - slice of primes 
func findPrimeNumbers(amount uint) []uint {
	var output = make([]uint, 0, amount)
	output = append(output, 2) // we know that the first prime number is 2
	
	var number uint = 3
	for len(output) < int(amount) {	

		// 'number' is prime if the rest of division it 
		// with any of the found primes isn't equal to zero
		var is_prime = true
		for _, value := range(output) {

			// it's enought to check only those primes whose square is less than 'number'
			if value * value > number {
				break
			} 
			
			if number % value == 0 {
				is_prime = false
				break
			}
		}

		if is_prime {
			output = append(output, number)
		}

		// select next odd number
		number += 2
	}

	return output
}

func main() {
	var primes = findPrimeNumbers(100)
	for idx, number := range primes {
		fmt.Printf("%03d - %d\n", idx + 1, number)
	}
}