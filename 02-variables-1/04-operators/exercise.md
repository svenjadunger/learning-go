# Operators

Go supports all usual math operators like `+`, `-`, `*`, `/`, and `%` (modulo).

They always work on two values (or variables).

```go
ten := 5 + 5
five := 10 / 2
```

You can join strings with `+`.

```go
fullName := firstName + " " + lastName
```

Go supports short forms of incrementing variables.

```go
// Same thing
i = i + 1
i += 1
i++
```

Same thing works for other operators.

```go
counter--
multiplier *= 2
divider /= 10
```

## Exercise

File: `02-variables-1/04-operators/main.go`

Add the missing `totalCost` variable that calculates the total order cost.

Multiply the product cost by the number of products and add the shipping cost.
