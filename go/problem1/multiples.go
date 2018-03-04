package main

func SumOfMulitiples(max, firstNumber, secondNumber int) int {
	sum := 0
	for i := 0; i < max; i++ {
		if i%firstNumber == 0 || i%secondNumber == 0 {
			sum += i
		}
	}
	return sum
}
