package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{-1, 2, 1},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d, want %d", test.a, test.b, result, test.expected)
		} else {
			t.Logf("Passed")
		}
	}
}
