# Go Learning Course Notes

Welcome to your Go learning journey! This document summarizes all the key concepts, code patterns, and pro tips from your interactive notes website. Use it to quickly recall and review Go basics and best practices.

---

## 1. Hello World

**Purpose:** Your first Go program! Prints a message to the console and introduces the `main` function and `fmt` package.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

💡 **Pro Tip:** Every Go program starts with `package main` and a `main()` function.

---

## 2. Simple Values

**Types covered:** Integer, String, Boolean, Float. Go is statically typed, so each value has a type.

```go
//integer
fmt.Println(1)
fmt.Println(1 + 1)

//string
fmt.Println("Hello, World!")

//bool
fmt.Println(true)
fmt.Println(false)

//floats
fmt.Println(10.5)
fmt.Println(7.0 / 2.3)
```

💡 **Pro Tip:** Use `//` for single-line comments to explain your code.

---

## 3. Variables

**Ways to declare:** `var` (with or without type), shorthand `:=`, and uninitialized variables. Variables must be declared before use. Types can be inferred.

```go
var name string = "Shahid"
var name2 = "Amin" // type inferred
var isAdult bool = true
var age int = 30
var height float64 = 5.9
var country string // defaults to ""
city := "Biratnagar" // shorthand
```

💡 **Pro Tip:** Use `:=` for quick variable declaration inside functions.

---

## 4. Constants

**Constants** are immutable values declared with `const`. They cannot be changed after assignment.

```go
const age = 30
const name string = "golang"
// name = "javascript" // error: cannot reassign constant
const (
    port = 5000
    host = "localhost"
)
```

💡 **Pro Tip:** Use constants for values that never change, like configuration or magic numbers.

---

## 5. For Loops

**Go's only loop:** `for` is the only looping construct. It can be used in several ways:

### While-style For

```go
i := 1
for i <= 3 {
    fmt.Println(i)
    i++
}
```

_Use this for loops with unknown iteration count._

### Classic For

```go
for j := 0; j < 3; j++ {
    fmt.Println(j)
}
```

_Classic C-style for loop, good for counting._

### Range Loop

```go
for i := range 3 {
    fmt.Println(i)
}
```

_Use `range` to iterate over arrays, slices, maps, or channels._

### Infinite Loop

```go
for {
    fmt.Println("Infinite Loop")
}
```

_Break out of infinite loops with `break` or `return`._

---

## 6. If/Else

**Conditional logic:** Go supports several forms of `if` and `else` for decision making.

### Simple If

```go
age := 20
if age >= 18 {
    fmt.Println("You are an adult.")
}
```

_Use for single, clear conditions._

### If-Else

```go
age := 16
if age >= 18 {
    fmt.Println("You are an adult.")
} else {
    fmt.Println("You are not an adult.")
}
```

_Provides an alternative path if the condition is false._

### If-Else If-Else

```go
percentage := 40
if percentage > 80 {
    fmt.Println("A")
} else if percentage > 65 {
    fmt.Println("B+")
} else {
    fmt.Println("Fail")
}
```

_Chain multiple conditions for complex logic._

### Inline Variable Declaration

```go
if age := 15; age >= 18 {
    fmt.Println("You are an adult.")
} else {
    fmt.Println("You are not an adult.")
}
```

_Declare a variable just for the if/else scope._

### Logical Operators

```go
role := "admin"
hasPermission := true
if role == "admin" || hasPermission {
    fmt.Println("User has admin permission")
}
```

_Use `&&` (and), `||` (or) to combine conditions._

### No Ternary Operator

```go
// Go does not support: result := condition ? a : b
// Use:
if condition {
    result = a
} else {
    result = b
}
```

_Go does **not** have a ternary operator. Use if/else instead._

---

> _This document is auto-generated from your interactive Go notes website. Add more sections as you learn!_
