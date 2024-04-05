# Nil

Declared pointers are equal to their zero-value: `nil`.

```go
var current *string // nil
```

A `nil` indicates that the pointer doesn't point to anything.

You can use this to indicate that a variable wasn't assigned yet.

```go
var isUpdated *bool

if isUpdated == nil {
	fmt.Println("isUpdated not set")
} else {
	fmt.Println("isUpdated set to", *isUpdated)	
}
```

An `if` like this is often called a *nil-check*. It's key when working with pointers.

If you try accessing a `nil` pointer's value (also called *dereferencing*) your application will panic.

## Exercise

File: `10-memory/02-nil/main.go`

The `UpdateUser` function updates the user's name and password.

Sometimes, the user wants to update just one of these.
The function uses string pointers to let the user pass a `nil` in case they don't want to update one of the fields.

However, the function is incomplete. If you pass a `nil` now, it panics. 

Your task is to add proper *nil-checks* to `UpdateUser`.
