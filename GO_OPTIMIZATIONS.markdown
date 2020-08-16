# Go Optimizations

When I discover a way to optimize.

## Don't use pointers to a slice

It may sound dumb but when I [implemented quicksort in Go](go/src/quicksort/quicksort.go), I started with fixed size arrays (stuff like `var numbers [1000]int`). As you may be aware of, arrays are pass-by-value in Go, which means that if you do a function call like `f(numbers)`, the entire array is going to be copied into another one. So I used pointers to avoid that and the calls looked like `f(&numbers)`.

Then slices came into play. I replaced my declaration with `var numbers := make([]int, size)`. `numbers` isn't an array anymore, it's a slice on top of an array of length 1000 and capacity 1000.

I changed my functions definition to handle a pointer to a slice. Which is kindof dumb because slices aren't pass-by-value, they are pass-by-reference.

When I realised that, I [changed the code](https://github.com/TanguyAndreani/algorithms/commit/7bcf485e00aa2a254b663156eb14318fb9f06353) so that my functions would take in a slice, not a pointer to a slice.

As you can see in the commit's diff, it cut the running time in half!
