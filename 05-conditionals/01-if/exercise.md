# If

To make decisions, use the `if` statement and comparison operators.

```go
if speedInKnots >= 64 {
    // A Hurricane!
}
```

```go
if role == "admin" {
	// Access granted
}
```

### Operators

* `==`: equal (mind the double `=`)
* `!=`: not equal
* `<`: less than
* `>`: greater than
* `<=`: less than or equal to
* `>=`: greater than or equal to

## Exercise

File: `05-conditionals/01-if/main.go`

Fill in `DescribeNumber` implementation to return `"negative"` if the provided number is smaller than `0`.
In other case, the function should return `"zero or more"`. 
