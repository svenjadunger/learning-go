# Globals

When you declare a variable inside a function, it's called a **local variable**.
It means other functions can't access it.

You can also declare **global variables** outside of functions.
They have their uses, but are rarely needed — stick to local variables when you can.

```go
var globalName = "Alice"

func main() {
	localName := "Bob"	
}
```

## var blocks

There's a short syntax for declaring many variables at once.
You can use it in both local and global scopes.

Variables in the same `var` block are not related to each other in any way.

```go
var (
	name  = "Alice"
	hours = 10
)
```

The same works for constants.

```go
const (
	pi         = 3.1415
	hoursInDay = 24
)
```

## Exercise

File: `02-variables-1/05-globals/main.go`

Add the missing `taxRate` global constant and set it to `0.1`.
