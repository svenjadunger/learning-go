# Slices

Most of the time, you won't work with **arrays** directly. Instead, you'll use **slices**.

You can think of slices as dynamic arrays. You don't declare how many elements they have.
They can expand and shrink.

When declaring a slice, omit the size from brackets.

```go
var contacts []string
```

A declared slice is equal to `nil`. `nil` means "nothing" or "no value" — the slice is not initialized yet.

If you try to access elements of a `nil` value, your application will crash.

```go
var contacts []string

// This will crash your application with message:
// panic: runtime error: index out of range [0] with length 0
contacts[0] = "Jenny"
```

You can use the *slice literal* to initialize an empty slice:

```go
contacts := []string{}
```

Or preallocate some memory for a slice using `make`:

```go
// A slice of five empty strings
contacts := make([]string, 5)
```

Similarly to arrays, you can initialize a slice with a set of values.

```go
contacts := []string{
	"Alice",
	"John",
	"Emma",
}
```

## Exercise

File: `04-collections-1/02-slices/main.go`

Create the missing `actions` **global variable**: a `string` slice.

Initialize it with four elements: `create`, `read`, `update`, and `delete`.


<div class="accordion" id="hints-accordion">

<div class="accordion-item">
	<h3 class="accordion-header" id="hints-accordion-header-1">
	<button class="accordion-button fs-4 fw-semibold collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#hints-accordion-body-1" aria-expanded="false" aria-controls="hints-accordion">
		Hint #1
	</button>
	</h3>
	<div id="hints-accordion-body-1" class="accordion-collapse collapse" aria-labelledby="hints-accordion-header-1" data-bs-parent="#hints-accordion">
	<div class="accordion-body">

Global variables are defined **outside** the `main` function.

Use the `var` keyword.

Remember, when using `var`, you need to use `=` to assign values, not `:=`.

```go
var globalSlice = []string{"one", "two", "three"}
```

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

Don't modify the slice inside the `main` function. The test code doesn't run `main` at all.

</div>
	</div>
	</div>

</div>
