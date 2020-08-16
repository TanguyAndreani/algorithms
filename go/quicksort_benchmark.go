package main

import (
  "fmt"
  "quicksort"
  "benchmark"
  "math/rand"
)

func fillWithRandomNumbers(numbers *[]int, count int) {
  for i := 0; i < count; i++ {
    (*numbers)[i] = rand.Intn(1000)
  }
}

func copyInto(dest *[]int, src *[]int, count int) {
  for i := 0; i < count; i++ {
    (*dest)[i] = (*src)[i]
  }
}

func main() {
  size := 300000

  numbers1 := make([]int, size)
  numbers2 := make([]int, size)

  fillWithRandomNumbers(&numbers1, size)
  copyInto(&numbers2, &numbers1, size)

  fmt.Println("Benchmark for slices of", size, "items sorted randomly")

  fmt.Println("Iterative Quicksort:", benchmark.Measure(func () {
    quicksort.QuicksortIterative(&numbers1, 0, size - 1)
  }))

  fmt.Println("Recursive Quicksort:", benchmark.Measure(func () {
    quicksort.QuicksortRecursive(&numbers2, 0, size - 1)
  }))
}
