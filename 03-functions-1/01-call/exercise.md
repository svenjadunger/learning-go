# Functions

You've already worked with two functions: `fmt.Println` and `main`.
Let's see how you can create your own.

You can use functions to avoid writing the same thing twice,
and to make your code more readable.

To define a function, use the `func` keyword, followed by the function's name and parentheses (`()`).

```go
func PrintMessage() {
	// Your code goes here
}
```

You can define functions before or after the `main` function. It's up to you.

To call a function, use its name followed by parentheses. Similarly to how you've used `fmt.Println()` before.
If the function is in the same package, you don't need to add the package prefix, like `fmt.`.

```go
PrintMessage()
```

## Exercise

File: `03-functions-1/01-call/main.go`

Call the `Hello` function twice in the `main` function, so the output is:

```bash
Hello, World!
Hello, World!
```
