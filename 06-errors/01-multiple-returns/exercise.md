# Multiple Return Values

Functions can return **any number of values**.

You need to add another pair of parentheses to define such functions. 
Separate the types with commas (`,`).

```go
func FullName() (string, string) {
	return "Alice", "Smith"
}
```

```go
func MonthAndYear() (string, int) {
	month := "April"
	year := 2012
	return month, year
}
```

When calling the function, assign the same number of variables, separated by commas

```go
func main() {
	firstName, lastName := FullName()
}
```

## Exercise

File: `06-errors/01-multiple-returns/main.go`

Write the missing `Swap` function that takes two `string` arguments and returns them unchanged
but in reverse order.


<div class="accordion" id="hints-accordion">

<div class="accordion-item">
	<h3 class="accordion-header" id="hints-accordion-header-1">
	<button class="accordion-button fs-4 fw-semibold collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#hints-accordion-body-1" aria-expanded="false" aria-controls="hints-accordion">
		Hint #1
	</button>
	</h3>
	<div id="hints-accordion-body-1" class="accordion-collapse collapse" aria-labelledby="hints-accordion-header-1" data-bs-parent="#hints-accordion">
	<div class="accordion-body">

Remember to define both the arguments list: `(string, string)` and the return value: `(string, string)`.

</div>
	</div>
	</div>

</div>
