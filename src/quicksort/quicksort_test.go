package quicksort

import (
	"testing"
	"math/rand"
)

func fillWithRandomNumbers(numbers []int) {
	for index, _ := range numbers {
		numbers[index] = rand.Intn(1000)
	}
}

func BenchmarkQuicksortIterative(b *testing.B) {
	if b.N > 10 {
		numbers := make([]int, b.N)
		fillWithRandomNumbers(numbers)
		QuicksortIterative(numbers, 0, len(numbers) - 1)
	}
}

func BenchmarkQuicksortRecursive(b *testing.B) {
	if b.N > 10 {
		numbers := make([]int, b.N)
		fillWithRandomNumbers(numbers)
		QuicksortRecursive(numbers, 0, len(numbers) - 1)
	}
}
