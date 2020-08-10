package quadratic

import "math"

// Roots of a quadratic equation
type Roots struct {
	X1 float64
	X2 float64
}

// Solve finds the roots of the quadratic equation
// Input:
//  - a, b, c are the coefficients of the equation
// Output:
//  - roots, True if the equation has real roots
//  - zeroes, False if the equation doesn't have real roots
func Solve(a, b, c float64) (Roots, bool) {

	d := b*b - 4*a*c
	if d < 0 {
		return Roots{}, false
	}

	x1 := (-b + math.Sqrt(d)) / (2 * a)
	x2 := (-b - math.Sqrt(d)) / (2 * a)

	return Roots{x1, x2}, true
}
