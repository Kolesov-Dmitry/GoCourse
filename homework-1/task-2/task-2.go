package main

import (
	"fmt"
	"math"
)

// calculateRightTriangle calculates hypotinuse, perimeter and area of right triangle
// Input:
//  - firstLeg is the one of legs of the triangle
//  - secondLeg is the another one of legs of the triangle
// Output:
//  - hypotinuse of the triangle
//  - perimeter of the triangle
//  - area of the triangle
// In case of invalid input returns zero values
func calculateRightTriangle(firstLeg float64, secondLeg float64) (float64, float64, float64) {

	// check the input
	if firstLeg < 0.0 || secondLeg < 0.0 {
		return 0, 0, 0
	}

	var hypotinuse = math.Sqrt(firstLeg * firstLeg + secondLeg * secondLeg)
	var perimeter = firstLeg + secondLeg + hypotinuse
	var area = firstLeg * secondLeg / 2.0

	return hypotinuse, perimeter, area
}

func main() {

	var firstLeg, secondLeg float64
	fmt.Print("Введите первый катет: ")
	fmt.Scan(&firstLeg)
	fmt.Print("Введите второй катет: ")
	fmt.Scan(&secondLeg)

	var hypotinuse, perimeter, area = calculateRightTriangle(firstLeg, secondLeg)

	fmt.Print("\nГипотенуза: ", hypotinuse, "\nПериметр:   ", perimeter, "\nПлощадь:    ", area)
}
