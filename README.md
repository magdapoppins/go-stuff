# Hello, GO! 
Notes on a Go intro course with code samples.

## Go is...
- compiled (generates machine code that will run directly on the CPU, and thus perhaps having better performance than interpreted langs)
- garbage-collected (you don't have to manage memory as a developer, the runtime will handle the allocation and releasing of memory for you)
- using cuncurrency (doing multiple things at a time will be okay)

## Take a look
Some imports

```go 
package main

import (
    "sdl"
    "rand"
)
```
Stuff:

```go
const initBalance = 100
// This is a pointer:
var initShift *int = flag.Int("level", 1)

var specs = [...]string{
    specN,
    specT,
    specMirrored
}
```
A function:

```go
func drawBlock(x, y int, color TetrisBlockColor) {
    glx := gl.GLint(x)

    gl.Color3ub()
    gl.End()
}
```
No classes, but there are types, which can have extensionmethodlike "methods"

```go
type TetrisBlock struct {
    Filled bool
    Color TetrisBlockColor
}

func (self *TetrisBlock) Draw(x, y int) {
    if self.Filled{
        drawBlock(x, y, self.Color)
    }
}
```
Last but not least, there is a main()

```go
func main() {
    // Do everything
}
```

## Core concepts

### Packages
- partitioning global namespace
- similar to namespaces
- golang.org/pkg has all the built-in libs

### Imports 
- source code
- local packages
- remote packages

### Functions curiosity
A function looks like this, right: 
```go
func add(a int, b int) int {
    return a + b
}
```
If you got 2 params of the same type, you can just list them and add the type after the last one:
```go
func subtr(a, b int) int {
    return a - b
}
```
A function can return multiple things! 
```go
func swap(a, b string) (string, string) {
    return b, a
}
```
A "naked" return statement will returned the **named** values. This is only to be used in short functions, as it will harm readability of longer ones. 
```go
func split(sum int) (x, y int) {
    x = sum * 4/9
    y = sum - x
    return
}
```

### More syntactic sugar

The **var** statement declares a variable list. The type is last:
`var a, b, c bool`.
A var statement can include **initializers**, one per value. In this case the type can be omitted, as it will take the type of the initializer: `var i, j = 2, 5`.

Have you seen this strange thing? `:=`  
It is the short assignment statement. It can be used instead of a var declaration with implicit type: `k := 3`.

### Type conversions

```go
var i int = 100
var f float = float(i)
// Or, more simply put...
i := 42
f := float64(i)
```
There are also constants:
`const Pi = 3.14`

### The for loop
The for loop is the only looping construct in go.
```go
sum := 0
for 1 := 0; i < 10; i++ {
    sum += i
}    
```
For is go's "while". If you omit the looping conditions, an infinate loop is created:
```go
for {
 // 4ever
}
```
If statements. Again, the expression need not parenthesis, but braces are required: 
```go
if x > 10 {
    return math.Sqrt(x)
}
return 10
```






