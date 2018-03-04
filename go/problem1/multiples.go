package problem1

import "fmt"

type Problem struct {
}

func (p Problem) Question() {
	fmt.Printf("%d\n", sumOfMulitiples(10, 3, 5))
}

func sumOfMulitiples(max, firstNumber, secondNumber int) int {
	sum := 0
	for i := 0; i < max; i++ {
		if i%firstNumber == 0 || i%secondNumber == 0 {
			sum += i
		}
	}
	return sum
}
