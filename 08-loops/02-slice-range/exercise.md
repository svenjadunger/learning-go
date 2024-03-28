# Slice range

To iterate over collections, Go provides the `range` keyword. It has two forms.

The first uses one variable and assigns to it the *index* in each iteration (starting from zero).

```go
phoneTypes := []string{"home", "mobile", "work"}

for i := range phoneTypes {
	fmt.Println("Index", i)
	p := phoneTypes[i]
	fmt.Println("Call my", p, "number")
}
```

The second form uses two variables. The first variable is the index, just as in the previous form.
The second one is the current element of the slice.

```go
for i, p := range phoneTypes {
	fmt.Println("Index", i)
	fmt.Println("Call my", p, "number")
}
```

Most of the time, you won't need the index in the second form. You can just skip it using the `_` identifier.

```go
for _, p := range phoneTypes {
	fmt.Println("Call my", p, "number")
}
```

## Exercise

File: `08-loops/02-slice-range/main.go`

Write a `Sum` function that takes any number of `int` arguments (use the variadic form).
It should return a sum of all the arguments.