package problem2

import "fmt"

const max = 4000000

type Problem struct {
}

func (p Problem) Question() {
	i := 1
	nextNumber := 2
	sum := 0

	for {
		if nextNumber >= max {
			break
		}
		if nextNumber%2 == 0 {
			sum += nextNumber
		}
		nextNumber, i = nextNumber+i, nextNumber
	}
	fmt.Printf("Sum is : %d\n", sum)
}
