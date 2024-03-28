# Variadic arguments

There's a way to pass *any* number of arguments to a function. Such functions are called *variadic*.

The list of arguments is marked with three dots (`...`) before the type's name.

```go
func ShowNames(names ...string) {
	
}
```

The function can take other arguments as well, but they need to come before the variadic list.

```go
func ShowNames(howMany int, names ...string) {

}
```

The argument becomes a regular slice inside the function.

```go
func PrintCount(names ...string) {
	fmt.Println("Found", len(names), "names")
}
```

Actually, you've been using a variadic function from the beginning of the training.
`fmt.Println` takes any number of arguments to print.

You can "unpack" a slice of the same type into variadic arguments using the same three dots syntax.

```go
letters := []string{"a", "b", "c"}
PrintLetters(letters...)
```

In a sense, variadic arguments are no different from passing a slice to a function.
It's simply a more convenient way to call the function, as you don't need to create the slice before.

### Append

`append` is also a variadic function.

```go
names = append(names, "Alice", "Bob", "Charlie")
```

A handy pattern to join two slices is by using the append function and *unpacking* one of them:

```go
names := append(originalNames, newNames...)
```

## Exercise

File: `07-collections-2/03-variadic/main.go`

Fill in three functions: `DebugLog`, `InfoLog`, and `ErrorLog`.

We want to use them to prepend a log level to our log messages.

They should take any numbers of string arguments, and return them unchanged, with one additional `string` **added as the first element**: `"[DEBUG]"`, `"[INFO]"`, or `"[ERROR]"`.

For example, for a function call:

```go
ErrorLog("could not create user:", err.Error())
```

The expected result would be:

```go
[]string{"[ERROR]", "could not create user", err.Error()}
```
