# Panic

In any moment, you can use the `panic` function to immediately stop your application with an error message.

```go
panic("something terrible happened")
```

A more useful message usually comes from an error:

```go
if err != nil {
	panic(err)
}
```

A panic prints out a *stack trace*. It's a list of function calls that happened just before the panic.
It should help you understand what happened.

Most of the time, your code shouldn't `panic`. Immediately stopping your application is rarely a good design.

Sometimes it comes handy, like when writing a quick program to validate an idea. For now, just be aware it exists.

## Exercise

File: `06-errors/04-panic/main.go`

Write a `MustStoreMessage` function that takes a `string` argument.

It should call `StoreMessage` and panic if it returns an error.
