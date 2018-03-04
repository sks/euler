package largestfactor_test

import (
	"testing"

	"github.com/sks/euler/go/problem3/largestfactor"
)

func TestLargestPrimeFactors(t *testing.T) {
	largestPrime := largestfactor.LargestPrimeFactors(600851475143)
	if largestPrime != 6857 {
		t.Errorf("Expected %d to be %d\n", largestPrime, 6857)
	}
}
