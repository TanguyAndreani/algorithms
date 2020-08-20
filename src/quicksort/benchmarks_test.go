package quicksort

import (
	"testing"
)

func BenchmarkQuicksortIterative(b *testing.B) {
	if b.N > 10 {
		numbers := make([]int, b.N)
		fillWithRandomNumbers(numbers, 1000)
		QuicksortIterative(numbers, 0, len(numbers)-1)
	}
}

func BenchmarkQuicksortRecursive(b *testing.B) {
	if b.N > 10 {
		numbers := make([]int, b.N)
		fillWithRandomNumbers(numbers, 1000)
		QuicksortRecursive(numbers, 0, len(numbers)-1)
	}
}
