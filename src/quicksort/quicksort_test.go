package quicksort

import (
	"math/rand"
	"sort"
	"testing"
)

var simpleList = []int{5, 3, 2, 9, 10, 8, 7, 4, 1, 6}

func fillWithRandomNumbers(numbers []int) {
	for index, _ := range numbers {
		numbers[index] = rand.Intn(1000)
	}
}

func BenchmarkQuicksortIterative(b *testing.B) {
	if b.N > 10 {
		numbers := make([]int, b.N)
		fillWithRandomNumbers(numbers)
		QuicksortIterative(numbers, 0, len(numbers)-1)
	}
}

func BenchmarkQuicksortRecursive(b *testing.B) {
	if b.N > 10 {
		numbers := make([]int, b.N)
		fillWithRandomNumbers(numbers)
		QuicksortRecursive(numbers, 0, len(numbers)-1)
	}
}

func compareSlice(a []int, b []int) int {
	if len(a) != len(b) {
		return len(a) - len(b)
	}

	i := 0
	for i < len(a) {
		if a[i] != b[i] {
			return a[i] - b[i]
		}
		i++
	}
	return 0
}

type sortf func([]int, int, int)
type testf func(*testing.T)

func TestQuicksort(t *testing.T) {
	/* we use this function in g() */
	var f sortf

	var g testf = func(t *testing.T) {
		quicksorted := make([]int, len(simpleList))
		copy(quicksorted, simpleList)

		stdsorted := make([]int, len(simpleList))
		copy(stdsorted, simpleList)

		/* the sorting algorithm we're testing */
		f(quicksorted, 0, len(quicksorted)-1)

		/* the algorithm we're testing it against */
		sort.Ints(stdsorted)

		if compareSlice(quicksorted, stdsorted) != 0 {
			t.Error("Expected", stdsorted)
			t.Error("Got", quicksorted)
		}
	}

	f = QuicksortIterative
	t.Run("Quicksort Iterative", g)

	f = QuicksortRecursive
	t.Run("Quicksort Recursive", g)
}