package main

import (
	"fmt"
	"math/rand"
)

func shuffle(numbers []int) {
	// Fisher–Yates shuffle
	n := len(numbers)
	for i := range numbers {
		j := i + rand.Intn(n-i)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

func main() {
	cards := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	fmt.Printf("before: %+v \n", cards)

	shuffle(cards)

	fmt.Printf("after: %+v \n", cards)
}
