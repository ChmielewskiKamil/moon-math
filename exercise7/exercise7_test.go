package main

import "testing"

func TestLongDivision(t *testing.T) {
	// Table test for long division
	tests := []struct {
		dividend  int
		divisor   int
		quotient  int
		remainder uint
	}{
		{10, 3, 3, 1},
		{0, 5, 0, 0},
		{-7, 3, -3, 2},
	}

	for _, test := range tests {
		q, r := longDivision(test.dividend, test.divisor)
		if q != test.quotient || r != test.remainder {
			t.Errorf("longDivision(%d, %d) = %d, %d, want %d, %d", test.dividend, test.divisor, q, r, test.quotient, test.remainder)
		}
	}
}
