# New

Because a declared pointer is equal to `nil`, you need an extra step to assign it to an empty variable.

```go
var number int
pointer := &number
```

To make it shorter, use the built-in `new` function that creates a zeroed variable and returns a pointer to it.

```go
pointer := new(int)
// *pointer is equal 0
```

You can also use `new` to initialize structs.

```go
alice := new(User)

// Same thing
bob := &User{}
```

## Exercise

File: `10-memory/03-new/main.go`

Implement `AllocateBuffer` function that returns a new string buffer — a pointer to an empty string.

There should be a limit of 3 buffers the user can allocate. Any subsequent calls to the function should return `nil`.

You can use a global variable to count allocated buffers.