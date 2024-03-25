# Short form error handling

The usual error handling is a bit verbose.

```go
err := SendRequest()
if err != nil {
	return err	
}
```

If you'd like to make this code more concise, you can use a short form:

```go
if err := SendRequest(); err != nil {
	return err	
}
```

Important note: the function's return values are available only in the *scope* of the `if`.
You can't access them outside the `if` statement or its braces.
This form is mostly useful for functions that return no other values, just an `error`.

## Exercise

File: `06-errors/06-short-form/main.go`

`os.Stat` returns information about the given file:

```go
func Stat(name string) (FileInfo, error)
```

Change the `main` function to handle the [`os.Stat`](https://pkg.go.dev/os#File.Stat) call using the short form of error handling.

We want to check if the file exists. If an error happens, print it out. Otherwise, print "ok".

We don't care about the first argument. You can skip it.
