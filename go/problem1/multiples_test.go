package main_test

import "testing"

func test_sumOfMulitiples(t *testing.T) {
	sum := SumOfMulitiples(10, 3, 5)
	if sum != 23 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 23)
	}
}
