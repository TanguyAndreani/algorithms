package quicksort

import (
	"math/rand"
	"sort"
	"testing"
)

func fillWithRandomNumbers(numbers []int, max int) {
	for index, _ := range numbers {
		numbers[index] = rand.Intn(max)
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
	var simpleList = []int{5, 3, 2, 9, 10, 8, 7, 4, 1, 6}

	var randomList = make([]int, 100)
	fillWithRandomNumbers(randomList, 1000)

	/* we use this function in g() */
	var f sortf

	var ls []int

	var g testf = func(t *testing.T) {
		quicksorted := make([]int, len(ls))
		copy(quicksorted, ls)

		stdsorted := make([]int, len(ls))
		copy(stdsorted, ls)

		/* the sorting algorithm we're testing */
		f(quicksorted, 0, len(quicksorted)-1)

		/* the algorithm we're testing it against */
		sort.Ints(stdsorted)

		if compareSlice(quicksorted, stdsorted) != 0 {
			t.Error("Expected", stdsorted)
			t.Error("Got", quicksorted)
		}
	}

	ls = simpleList

	f = QuicksortIterative
	t.Run("Quicksort Iterative on simple list", g)

	f = QuicksortRecursive
	t.Run("Quicksort Recursive on simple list", g)

	ls = randomList

	f = QuicksortIterative
	t.Run("Quicksort Iterative on random list", g)

	f = QuicksortRecursive
	t.Run("Quicksort Recursive on random list", g)
}
