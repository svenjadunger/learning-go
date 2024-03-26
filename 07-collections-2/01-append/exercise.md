# Append

To add a new element to a slice, use the `append` function.

The first argument of `append` is the slice to append to. Then, any number of arguments that should be appended at the end.

`append` returns the result slice.
It doesn't modify the passed slice in-place, so you have to assign the result to a variable.

```go
colors := []string{}

colors = append(colors, "red")
colors = append(colors, "blue", "green")

// colors is equal to []string{"red", "blue", "green"}
```

## Exercise

File: `07-collections-2/01-append/main.go`

Write the missing `AddUser` function. It should take one `string` argument and add it to the `users` slice.