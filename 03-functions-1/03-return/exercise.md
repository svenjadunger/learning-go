# Returning a Value

Functions can return values. You need to define the type that the function returns.
You do it after the parentheses.

You can return either explicit values (like `12`) or a variable with matching type.

```go
func NumberOfMonths() int {
	return 12
}

func DaysInAWeek() int {
	days := 7
	return days
}
```

A function can take any number of arguments, separated by commas (`,`).

```go
func Add(x int, y int) int {
	sum := x + y
	return sum
}

func main() {
	five := Add(2, 3)
}
```

If arguments next to each other are of the same type, you can use a shorter form.
You can also write equations following the `return` immediately, with no additional variables.

```go
// Exactly the same thing
func Add(x, y int) int {
	return x + y
}
```

## Exercise

File: `03-functions-1/03-return/main.go`

Add the missing `Sum` function.

It should take 5 `int` arguments and return another integer, a sum of all arguments.


<div class="accordion" id="hints-accordion">

<div class="accordion-item">
	<h3 class="accordion-header" id="hints-accordion-header-1">
	<button class="accordion-button fs-4 fw-semibold collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#hints-accordion-body-1" aria-expanded="false" aria-controls="hints-accordion">
		Hint #1
	</button>
	</h3>
	<div id="hints-accordion-body-1" class="accordion-collapse collapse" aria-labelledby="hints-accordion-header-1" data-bs-parent="#hints-accordion">
	<div class="accordion-body">

Don't forget the return type (`int`) before the opening brace of the function.

</div>
	</div>
	</div>

<div class="accordion-item">
	<h3 class="accordion-header" id="hints-accordion-header-2">
	<button class="accordion-button fs-4 fw-semibold collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#hints-accordion-body-2" aria-expanded="false" aria-controls="hints-accordion">
		Hint #2
	</button>
	</h3>
	<div id="hints-accordion-body-2" class="accordion-collapse collapse" aria-labelledby="hints-accordion-header-2" data-bs-parent="#hints-accordion">
	<div class="accordion-body">

Argument names need to start with a letter, so you can't just use `1`, `2`, `3`, `4`, `5`.
Try `a`, `b`, `c`, `d`, `e`, or something similar.

</div>
	</div>
	</div>

</div>
