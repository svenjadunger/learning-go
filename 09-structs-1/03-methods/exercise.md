# Methods

Methods are functions bound to a single struct.

To write a method, write a regular function, but add an extra section in parentheses before its name.
It should contain a name and the type.

```go
type Money struct {
	Amount   int
	Currency string
}

func (m Money) Format() string {
	return fmt.Sprintf("%v %v", m.Amount, m.Currency)
}
```

The *receiver* part (`(m Money)`) is how you reference the variable inside the method.
The name is usually a one letter and should be the same for all methods on a single type.

When you consider the previous example, it's clear that methods are just a more convenient way to write functions working with a struct.

```go
func Format(m Money) string {
	return fmt.Sprintf("%v %v", m.Amount, m.Currency)
}

func (m Money) Format() string {
	return fmt.Sprintf("%v %v", m.Amount, m.Currency)
}
```

You call methods similarly to regular functions. Prefix the method name with the struct variable and a dot (`.`).

```go
formattedByFunction := Format(m)

formattedByMethod := m.Format()
```

## Exercise

File: `09-structs-1/03-methods/main.go`

Add an `Area` method to `Rectangle`. It should do the same thing as in the previous exercise.