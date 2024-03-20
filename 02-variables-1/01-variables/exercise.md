# Variables

So far, you've used the text type called `string`. In Go, any text in double-quotes is a `string`.

Values are much more helpful as variables that you can modify.

When you first declare a variable, pick a name for it and assign a value using `:=`:

```go
name := "Alice"
```

You can then modify the variable using the `=` operator. (Note the difference between `:=` and `=`.)

```go
name = "Bob"
```

## Exercise

File: `02-variables-1/01-variables/main.go`

Add a `year` variable equal to 2012. It can be an integer (`2012`) or a string (`"2012"`).

Notice how `fmt.Println()` uses the two variables.
It adds a space between each variable it prints.

The program should print:

```bash
Go 1.0 was released in 2012
```