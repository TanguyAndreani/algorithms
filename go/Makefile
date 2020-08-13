ROOTDIR=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

all: quicksort_benchmark.bin

quicksort_benchmark.bin: quicksort_benchmark.go src/quicksort/quicksort.go
	GOPATH=${ROOTDIR} go build -o $@ $*
