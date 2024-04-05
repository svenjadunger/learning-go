# Pointers

A **pointer** is a variable that *points* to another variable.

To declare a pointer, add an asterisk (`*`) before the type.
The type is crucial: a pointer can only point to variables of the defined type.

```go
alice := "Alice"
bob := "Bob"

var current *string
current = &alice
```

At this point, `current` points to `alice`. However, when you print `current`, you won't see `Alice`.
Instead, you'll see something like `0x14000010290`.

This is the *address in memory* of the variable `alice`.
See how it was assigned: the ampersand symbol (`&`) returns a variable's address.

```go
current = &alice
```

To access the actual variable, add an asterisk (`*`) before the variable.
It's called *dereferencing*.

```go
fmt.Println(current)  // Prints: 0x14000010290
fmt.Println(*current) // Prints: Alice
```

You can change what the pointer points to at any time.

```go
current = &bob
```

Pointers can be used as other types: as struct fields or in collections.

When reading complex types, go left-to-right. Pointers always refer to the type immediately to the right of the asterisk (`*`) symbol.

```go
// a map of string to pointers to integers
var counters map[string]*int
// a slice of pointers to booleans
var choices []*bool
// a pointer to slice of strings
var dynamicList *[]string
```

## Exercise

File: `10-memory/01-pointers/main.go`

Referencing other variables by pointers refers to the exact variable, not only the data it holds.

This exercise shows a `User` struct referencing `Order` structs. Thanks to using pointers, it will always reference
up-to-date structures, even if an order changes after being assigned.

Add the missing `User` structure with a `Name` (a `string`), and `Orders` (a slice of pointers to `Order`).
