# Returning errors

The functions you wrote so far didn't have side effects. There was an input and an output.
Functions like these are called *pure functions*.

Often, functions attempt things that can fail, such as accessing files, connecting over the network,
or validating the arguments. In such scenarios, the function needs to tell the caller somehow
if it succeeded.

In Go, errors are values returned from the function. If a function returns more than one value,
the last one is usually the error, with the type `error`. A `nil` value means no error.

Create new errors with `errors.New()` (from the `errors` package).

```go
func CheckUser(username string) error {
	if UserExists(username) {
		return nil
	} else {
		return errors.New("user not found")
	}
}
```

If a function returns multiple values and an error occurs, other arguments are usually returned as zero-value.

```go
func JoinName(firstName string, lastName string) (string, error) {
	if firstName == "" || lastName == "" {
		return "", errors.New("missing first name or last name")
	}
	
	return firstName + " " + lastName, nil
}
```

## Exercise

File: `06-errors/02-returning-errors/main.go`

Write a `Divide` function that takes two `float64` arguments.
It should return two values: the result of dividing them and an `error`.

If the second argument is `0`, the function should return an error, because division by zero is not allowed.

If the argument is valid, the `error` should be `nil`.