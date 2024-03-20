# Imports

Go code can be grouped into packages. You can use other packages by importing them.

For now, let's focus on the standard library that comes with Go.
There's a lot of helpful code there that's ready to use, saving your time.

One of these packages is `fmt` (short for "format") used to format text and display it.

You import packages by using the `import` keyword. It needs to be at the top of the file after the `package` declaration.

```go
import "fmt"
```

You can use the `Println` function from the `fmt` package to print messages on the screen.

To call a function, start with the package name, followed by a dot, and the function name.

```go
fmt.Println("This is a text")
```

Did you notice the `ln` suffix in `Println`? This function adds a "newline" at the end of the text.
There's also an `fmt.Print` function that doesn't do it.

## Exercise

File: `01-start/03-imports/main.go`

First, import the `fmt` package. 

Then, use the `fmt.Println` function to display a text within two lines, like this:

```bash
Hello
world
```