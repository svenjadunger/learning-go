# Function Arguments

Functions can take arguments. Arguments are variables that are passed to the function.

Similarly to variables, you have to specify the argument's name and type.
You do it between the parentheses after the function's name.

```go
func PrintNumber(number int) {
	fmt.Println("My favorite number is", number)
}

func main() {
	// Prints: My favorite number is 5
	// number is equal to 5 inside the function
	PrintNumber(5)
}
```

## Exercise

File: `03-functions-1/02-arguments/main.go`

Write the missing `Greet` function. Call it once again in the `main`, so that the output is:

```bash
Hello, Alice
Hello, Bob
```


<div class="accordion" id="hints-accordion">

<div class="accordion-item">
	<h3 class="accordion-header" id="hints-accordion-header-1">
	<button class="accordion-button fs-4 fw-semibold collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#hints-accordion-body-1" aria-expanded="false" aria-controls="hints-accordion">
		Hint #1
	</button>
	</h3>
	<div id="hints-accordion-body-1" class="accordion-collapse collapse" aria-labelledby="hints-accordion-header-1" data-bs-parent="#hints-accordion">
	<div class="accordion-body">

Function names are *case-sensitive*: `greet` won't work, it needs to be `Greet`.

</div>
	</div>
	</div>

</div>
