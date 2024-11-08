package main

import "fmt"

func main() {
	// fact5 := factorial(5)
	fact5 := factorialRec(5)
	fmt.Println(fact5)
}

func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result *= i
	}

	return result
}

func factorialRec(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	}
	return number * factorial(number - 1)
}