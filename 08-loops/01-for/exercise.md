# Loops

Go has one loop — `for`. It comes in a few flavors.

Using `for` without arguments creates an infinite loop. The code inside braces will repeat forever.

```go
for {
	// Infinite loop
}
```

You can add a boolean expression to `for`, similar to `if`.

The loop bellow will run ten times, as `x` goes from `0` to `9`.

```go
x := 0
for x < 10 {
	fmt.Println("x =", x)
	x++
}
```

There's also a short form of this approach:

```go
for x := 0; x < 10; x++ {
	fmt.Println("x =", x)
}
```

* The first part (`x := 0`) is executed once, at the start of the loop.
* The second part (`x < 10`) is a condition checked before starting the next iteration.
* The third part (`x++`) is executed after each iteration.

## Exercise

File: `08-loops/01-for/main.go`

Fill the `Alphabet` function. The function should return a slice of English alphabet letters.
The length of the slice is based on the `length` argument.

For example:

* `Alphabet(2)` should return `[]string{"a", "b"}`.
* `Alphabet(5)` should return `[]string{"a", "b", "c", "d", "e"}`.

Use the `characterByIndex` function to get the next alphabet character.
For `0`, it returns `"a"`, for `1`, it returns `"b"`, and so on.

### Bonus: Runes

A `rune` is a type that represents a single character. `rune` literals use single-quotes:

```go
s := "this is a string"
r := 'r' // This is a rune
```

A `string` is a slice of runes (`[]rune`).
