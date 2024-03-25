# Skipping unused variables

Sometimes, you don't care about the returned value, just if the result was successful.

A common pattern is to use the `_` identifier to skip a variable

```go
func main() {
	_, err := OpenFile("results.csv")
	if err != nil {
		fmt.Println("Failed to open file")
	}
}

func OpenFile(name string) (File, error) { 
	// ...	
}
```

It's not optional. Go won't compile code with declared but unused variables.

## Exercise

File: `06-errors/05-skipping-unused/main.go`

Write a `CheckFile` function that takes a `string` argument and returns a `bool`.

It should:

* Call `os.ReadFile` passing the `string` input value.
* Check the `error` returned by `os.ReadFile`.
* Return `true` if the error is `nil`, and `false` otherwise.

Here's the signature of [`os.ReadFile`](https://pkg.go.dev/os#ReadFile):

```go
func ReadFile(name string) ([]byte, error)
```
