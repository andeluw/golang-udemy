package main

import "fmt"

func main() {
	// sum := sumup(1, 2, 3, 4)
	num := []int{1, 5, 10}
	sum := sumup(num...)
	fmt.Println(sum)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}