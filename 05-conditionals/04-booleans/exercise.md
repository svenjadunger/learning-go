# Complex booleans

Sometimes, a single condition is not enough to make a decision.

You can combine conditions with `&&` (and) and `||` (or).

```go
if a > 0 && b > 0 {
	fmt.Println("a and b are positive")
}

if a == 0 || b == 0 {
	fmt.Println("a or b is equal to zero")
}
```

You can assign the condition result to a boolean variable. 

```go
aNegative := a < 0
bNegative := b < 0
bothNegative := a && b
```

You can skip operators when using boolean variables in the `if` statement. Use the `!` operator to negate (flip) a boolean.

```go
if bothNegative == true {
	
}
// Same thing
if bothNegative {
	
}

if bothNegative == false {
	
}
// Same thing
if !bothNegative {
	
}
```

## Exercise

File: `05-conditionals/04-booleans/main.go`

Fill in the `In20thCentury` function. It should return `true` if the passed `year` is in the 20th century (years 1901 to 2000).
