# Structs

Struct are groups of variables. The variables held by a struct are called "fields".
They let you treat many values as a single unit.

```go
type Money struct {
	Amount   int
	Currency string
}
```

You can create a struct with all fields set to zero-values:

```go
m := Money{}
```

Or you can fill in all or some of them:

```go
m := Money{
	Amount:   10,
	Currency: "USD",
}
```

To access a field, use the struct variable followed by a dot (`.`) and the field name.

```go
// Read a struct's field
amount := m.Amount

// Write to a struct's field
m.Currency = "EUR"
```

## Exercise

File: `09-structs-1/01-structs/main.go`

Fill in missing fields of the `Point` structure.
