# Else

Use `else` to execute code when an `if` condition is not met.

```go
if role == "admin" { 
	// Access granted 
} else { 
	// Access denied
}
```

You can combine any number of conditions with `else if`.

```go
if role == "admin" {
    // Admin access granted 
} else if role == "user" {
    // User access granted
} else {
    // Access denied
}
```

## Exercise

File: `05-conditionals/02-else/main.go`

Change the implementation of `DescribeNumber` to return:

* `"zero"` if the provided number is `0`,
* `"negative"` if the provided number is less than `0`,
* and `"positive"` if the provider number is greater than `0`.