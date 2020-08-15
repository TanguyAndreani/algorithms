# Go Cheatsheet

For things you certainly don't need to learn by heart. (But it's always
useful to have them noted somewhere.)

## Overriding a package name

```go
import "log"
import mylog "my/log"

// log.DoThis()
// mylog.DoThat()
```

## Declaring variables

```go
a := 12
var a = 12
var a int
var a int = 12
```

## Declaring multiple variables at once

```go
var a, b int = 1, 3
```

## For loops

```go
for {}      // while (1) {}
for cond {} // while (cond) {}
for init; cond; post {} // for (init; cond; post) {}
for index, value := range array {}
for _, value := range array {}
```

## Maps

```go
mymap := make(map[string]int) // mymap[string] = int
```

## Compile time constants

```go
const (
  a = 12
  b = "hello"
)
```

## Reading a file

```go
input := bufio.NewScanner(os.Stdin)
for input.Scan() { // puts line in input.Text() and return true or return false on EOF
  fmt.Println(input.Text())
}
```
