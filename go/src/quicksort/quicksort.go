package quicksort

func swap(numbers *[]int, i int, j int) {
	tmp := (*numbers)[i]
	(*numbers)[i] = (*numbers)[j]
	(*numbers)[j] = tmp
}

func partition(numbers *[]int, l int, r int) int {
	pivot := (*numbers)[r]
	i := l

	for j := l; j < r; j++ {
		if (*numbers)[j] < pivot {
			swap(numbers, i, j)
			i++
		}
	}
	swap(numbers, r, i)
	return i
}

func QuicksortRecursive(numbers *[]int, l int, r int) {
	if l < r {
		pivot := partition(numbers, l, r)

		QuicksortRecursive(numbers, l, pivot-1)
		QuicksortRecursive(numbers, pivot+1, r)
	}
}

func QuicksortIterative(numbers *[]int, l int, r int) {
	stack := make([]int, r-l+1)
	var stack_idx int

	stack[stack_idx] = l

	stack_idx++
	stack[stack_idx] = r

	for stack_idx >= 0 {
		r = stack[stack_idx]
		stack_idx--

		l = stack[stack_idx]
		stack_idx--

		pivot := partition(numbers, l, r)

		if pivot-1 > l {
			stack_idx++
			stack[stack_idx] = l

			stack_idx++
			stack[stack_idx] = pivot - 1
		}

		if pivot+1 < r {
			stack_idx++
			stack[stack_idx] = pivot + 1

			stack_idx++
			stack[stack_idx] = r
		}
	}
}
