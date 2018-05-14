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

