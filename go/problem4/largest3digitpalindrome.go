package problem4

import "fmt"

type Problem struct {
}

func (p Problem) Question() {
	fmt.Printf("Biggest palindrom is %d\n", LargestPalindrome())
}

func LargestPalindrome() int {
	max := 0
	for first := 999; first > 101; first-- {
		for second := 999; second >= first; second-- {
			product := first * second
			if max > product {
				continue
			}
			if IsPalindrom(product) {
				if max < product {
					max = product
				}
			}
		}
	}
	return max
}

func IsPalindrom(number int) bool {
	reverse := 0
	duplicate := number
	for {
		remainder := duplicate % 10
		reverse = reverse*10 + remainder
		duplicate /= 10

		if duplicate == 0 {
			break
		}
	}
	return reverse == number
}
