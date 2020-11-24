# Go Cheatsheet

For things you certainly don't need to learn by heart. (But it's always
useful to have them noted somewhere.)

- [Go Cheatsheet](#go-cheatsheet)
  * [Packages](#packages)
    + [Overriding a package name](#overriding-a-package-name)
  * [Basics](#basics)
    + [Declaring variables](#declaring-variables)
    + [Declaring multiple variables at once](#declaring-multiple-variables-at-once)
    + [For loops](#for-loops)
    + [Compile time constants](#compile-time-constants)
  * [Slices and arrays](#slices-and-arrays)
    + [Arrays](#arrays)
    + [Slices](#slices)
  * [Maps](#maps)
  * [I/O](#i-o)
    + [Reading a file](#reading-a-file)
    + [Opening a file](#opening-a-file)

<small><i><a href='http://ecotrust-canada.github.io/markdown-toc/'>Table of contents generated with markdown-toc</a></i></small>


## Packages

### Overriding a package name

```go
import "log"
import mylog "my/log"

// log.DoThis()
// mylog.DoThat()
```

## Basics

### Declaring variables

```go
a := 12 // short variable declaration
var a = 12
var a int
var a int = 12
```

### Declaring multiple variables at once

```go
var a, b int = 1, 3
```

### For loops

```go
for {}      // while (1) {}
for cond {} // while (cond) {}
for init; cond; post {} // for (init; cond; post) {}
for index, value := range myarray {}
for _, value := range myarray {}
for key, value := range mymap {}
```

### Compile time constants

```go
const (
  a = 12
  b = "hello"
)
```

## Slices and arrays

[Slices intro](https://blog.golang.org/slices-intro)

### Arrays

- An array variable denotes the entire array (pass-by-value)

```go
var a [4]int
a[0] = 1
i := a[0]
b := [2]string{"Penn", "Teller"}
```

### Slices

Just like arrays but size is left empty (`[]T`).

```go
x := [4]string{"Лайка", "Белка", "Стрелка", "Hello"}
s := x[:] // a slice referencing the storage of x
p := x[1:2] // a slice of length 2 and capacity 3
```

A little bit like an interface between you and a real array.

```go
package main

import "fmt"

func main() {
  x := [6]string{"hello", "world", "my", "name", "is", "Foo"}
  fmt.Println("x:", s, len(x), cap(x))
  // => x: [hello world my name is Foo] 6 6

  s := x[:]
  fmt.Println("s:", s, len(s), cap(s))
  // => s: [hello world my name is Foo] 6 6

  p := x[1:3]
  fmt.Println("p:", p, len(p), cap(p))
  // => p: [world my] 2 5

  // reslicing p to including every remaining items of x
  // but we can't go back to the first item...
  t := p[0:cap(p)]
  fmt.Println("t:", t, len(t), cap(t))
  // => t: [world my name is Foo] 5 5
}
```

## Maps

```go
mymap := make(map[string]int) // mymap[string] = int
```

## I/O

### Reading a file

```go
input := bufio.NewScanner(os.Stdin)
for input.Scan() { // puts line in input.Text() and return true or return false on EOF
  fmt.Println(input.Text())
}
```

### Opening a file

Check out the os package's [documentation](https://golang.org/pkg/os).

```go
f, err := os.Open(filepath) // readonly
f, err := os.OpenFile(filepath, flag, permissions)
```
