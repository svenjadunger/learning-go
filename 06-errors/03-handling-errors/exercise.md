# Handling errors

To tell if a function was successful, you have to check the error value.

```go
err := EmailExists("alice@example.com")
if err != nil {
	fmt.Println("Error:", err)
}
```

In Go, you will see this snippet a lot:

```go
if err != nil {
	return err
}
```

This means the error is returned again, propagated "up the stack".
This can happen multiple times, until one handler decides what to do with it.

One way of handling an error is simply logging it.

```go
if err != nil {
	fmt.Println(err)
}
```

Other times, you might choose to quit the program or show the error to the user.

## Exercise

File: `06-errors/03-handling-errors/main.go`

Extend the `CreateDirectory` function so that it:

* Properly checks the `error` returned by `os.Mkdir`.
* Returns `true` if the error is `nil`.
* Returns `false` otherwise.