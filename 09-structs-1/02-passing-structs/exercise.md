# Passing Structs as Arguments

Structs come especially useful for passing a group of variables as a single argument.

```go
type Money struct {
	Amount   int
	Currency string
}

func Format(m Money) string {
	return fmt.Sprintf("%v %v", m.Amount, m.Currency)
}
```

### Initializing structs

You can initialize structs without field names, passing values in the order they are defined.

```go
func main() {
	m := Money{100, "USD"}
	fmt.Println(Format(m))
}
```

## Exercise

File: `09-structs-1/02-passing-structs/main.go`

Write an `Area` function that takes an `Rectangle` as an argument.
It should return the rectangle's area (width multiplied by length).
