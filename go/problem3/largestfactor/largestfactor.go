package largestfactor

func LargestPrimeFactors(numberToCheck int) (largestPrime int) {

	for numberToCheck%2 == 0 {
		largestPrime = 2
		numberToCheck = numberToCheck / 2
	}

	for i := 3; i*i <= numberToCheck; i = i + 2 {
		// while i divides n, append i and divide n
		for numberToCheck%i == 0 {
			largestPrime = i
			numberToCheck = numberToCheck / i
		}
	}

	if numberToCheck > 2 {
		largestPrime = numberToCheck
	}

	return largestPrime
}
