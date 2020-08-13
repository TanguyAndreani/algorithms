all: quicksort_benchmark.bin

quicksort_benchmark.bin: quicksort_benchmark.go src/quicksort/quicksort.go
	GOPATH=${PWD} go build -o $@ $*
