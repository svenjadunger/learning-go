# Types

Go supports the types you can find in most programming languages, like:

* `int`, an integer number, e.g., `42`, `0`, or `-200`
* `bool`: `true` or `false`
* `float64`, a decimal, e.g., `3.1415`
* `string`, a text, e.g., `"Alice"`

## Declaring variables

**Before using a variable in Go, you have to declare it**.

To declare a variable, use the `var` keyword together with a name and a type.

```bash
var <name> <type>
```

For example:

```go
var firstName string
```

Such variables have no value assigned yet. They're equal to the **zero value**.
It depends on the type. For strings it's an empty string (`""`), for numbers it's `0`, and for booleans it's `false`.

To assign a value use the equal sign:

```go
firstName = "Alice"
```

You can also assign the value right after declaration:

```go
var firstName string = "Alice"
```

Go compiler knows that `"Alice"` is a `string`, so you can skip the type:

```go
var firstName = "Alice"
```

Finally, you can use a short form of the same thing. You already know it from the previous exercise.

```go
firstName := "Alice"
```

Remember:

* Use `=` to assign values to already declared variables and when using `var`.
* Use `:=` when declaring and assigning variable for the first time.

## Exercise

File: `02-variables-1/02-types/main.go`

Fill in the variables so that the output:

* Prints the number of days in December.
* Prints a more precise pi number, `3.1415`.
* Confirms that you're learning Go!

```bash
days in december: 31
pi: 3.1415
learning go: true
```