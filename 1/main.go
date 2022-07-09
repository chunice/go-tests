package main

import "fmt"

func maxSubtractNumber(numbers []int) int {
	min, max := 1<<31, 0
	for _, n := range numbers {
		if n > max {
			max = n
		} else if n < min {
			min = n
		}
	}
	return max - min
}

func main() {
	numbers := []int{5, 8, 10, 1, 3}

	sub := maxSubtractNumber(numbers)

	fmt.Printf("max subtract number: %d", sub)
}
