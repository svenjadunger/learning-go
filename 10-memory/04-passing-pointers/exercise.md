# Passing pointers

If a function modifies its arguments, it doesn't really access the original variable. 
What it changes is a *copy* of the passed variable.

```go
func Increase(number int) {
	number++	
}

func main() {
	x := 0
	Increase(x) 
	// x is still 0
}
```

Pointers let you modify the actual passed variable.

```go
func Increase(number *int) {
	*number++
}

func main() {
	x := 0
	Increase(&x) 
	// x is now 1
}
```

## Exercise

File: `10-memory/04-passing-pointers/main.go`

The `Deduplicate` function uses maps for *deduplicating* a slice of strings.

First, it fills the map, using slice values as keys.
Then, it converts the keys to a slice of strings.

The duplicates are gone, because setting the same key multiple times in a map has no effect.

Instead of returning the result, `Deduplicate` saves it back to the input argument.
That's why it should accept a pointer to a slice of strings as the argument.

Your task is to:

* Fill the missing argument in the function definition.
* Assign the result (`keys`) to the input argument.