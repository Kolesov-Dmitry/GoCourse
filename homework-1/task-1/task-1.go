package main

import "fmt"

// The amount of rubles that can be obtained for one dollar
const RubleToDollarRate int = 70

// convert calculates the amount of dollars that can be obtained for input amount of rubles
// Input:
//  - rubles is the amount of rubles to convert to dollars
// Output:
//  - amount of dollars
//  - amount of cents
// If input is invalid zero values are returned
func convert(rubles int) (int, int) {

	// check the input
	if rubles < 0 {
		return 0, 0
	}

	// Calculate the amount of dollars and cents
	var dollars = rubles / RubleToDollarRate
	var cents = float32(rubles - dollars * RubleToDollarRate) * 100.0 / float32(RubleToDollarRate)

	return dollars, int(cents)
}

func main() {

	var rubles int
	fmt.Print("Введите количество рублей: ")
	fmt.Scan(&rubles)

	var dollars, cents = convert(rubles)

	fmt.Print("\nДоллары: ", dollars, "\nЦенты:   ", cents)
}
