package quadratic

import "testing"

type testCaseSolve struct {
	name         string
	coefficients [3]float64
	expected     Roots
	valid        bool
}

func TestSolve(t *testing.T) {
	// prepare test cases
	var testCases = [...]testCaseSolve{
		{
			name:         "test case 1",
			coefficients: [3]float64{1, 2, 3},
			expected:     Roots{},
			valid:        false,
		},
		{
			name:         "test case 2",
			coefficients: [3]float64{1, 2, 1},
			expected:     Roots{-1, -1},
			valid:        true,
		},
		{
			name:         "test case 3",
			coefficients: [3]float64{3, -14, -5},
			expected:     Roots{5, -0.3333333333333333},
			valid:        true,
		},
	}
	// do the testing
	for _, testCase := range testCases {
		roots, valid := Solve(testCase.coefficients[0], testCase.coefficients[1], testCase.coefficients[2])
		if roots != testCase.expected || valid != testCase.valid {
			t.Errorf("\nFor '%s'\nExpected: %v, %v\nGot: %v, %v",
				testCase.name, testCase.expected, testCase.valid, roots, valid)
		}
	}
}
