package sum

import "testing"

type testCaseSum struct {
	name     string
	values   []float64
	expected float64
}

func TestSum(t *testing.T) {
	// prepare test cases
	var testCases = [...]testCaseSum{
		{
			name:     "test case 1",
			values:   []float64{1, 2, 3, 4},
			expected: 10,
		},
		{
			name:     "test case 2",
			values:   []float64{1.5, 2.5, 3.5},
			expected: 7.5,
		},
	}

	// do the testing
	for _, testCase := range testCases {
		s := Sum(testCase.values)
		if s != testCase.expected {
			t.Error("For", testCase.name,
				"Expected:", testCase.expected,
				"Got:", s)
		}
	}
}
