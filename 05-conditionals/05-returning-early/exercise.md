# Returning early

You can use `return` to end a function early, even if it doesn't return any value.

This helps to avoid nested `if` statements and keeps your code readable.

```go
func Greet(name string) {
	if name == "" {
		return
	}

	fmt.Println("Hello", name)
}
```

## Exercise

File: `05-conditionals/05-returning-early/main.go`

Add an early return to the `ResetPassword` function. If the `code` isn't equal to `2022`, the function shouldn't continue.