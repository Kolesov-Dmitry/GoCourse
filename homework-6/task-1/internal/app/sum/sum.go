package sum

// Sum calculates the sum of input values
// Input:
//  - values to calculate the sum
// Output:
//  - sum of values
func Sum(values []float64) float64 {
	var sum float64
	for _, val := range values {
		sum += val
	}

	return sum
}
