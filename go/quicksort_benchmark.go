package main

import (
	"benchmark"
	"fmt"
	"math/rand"
	"quicksort"
)

func fillWithRandomNumbers(numbers []int) {
	for index, _ := range numbers {
		numbers[index] = rand.Intn(1000)
	}
}

func main() {
	size := 300000

	numbers1 := make([]int, size)
	numbers2 := make([]int, len(numbers1))

	fillWithRandomNumbers(numbers1)
	copy(numbers2, numbers1)

	fmt.Println("Benchmark for slices of", size, "items sorted randomly")

	fmt.Println("Iterative Quicksort:", benchmark.Measure(func() {
		quicksort.QuicksortIterative(numbers1, 0, size-1)
	}))

	fmt.Println("Recursive Quicksort:", benchmark.Measure(func() {
		quicksort.QuicksortRecursive(numbers2, 0, size-1)
	}))
}
