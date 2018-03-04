package main

import (
	"fmt"

	"github.com/sks/euler/go/problem3/largestfactor"
)

const number = 600851475143

func main() {
	primeFactors := largestfactor.LargestPrimeFactors(number)
	fmt.Println(primeFactors)
}
