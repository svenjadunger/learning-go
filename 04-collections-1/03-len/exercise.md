# Len

The number of elements in a slice (or array) is called *length*. The `len` function returns the current length of a slice or array.

```go
methods := []string{"GET", "POST", "DELETE"}
numberOfMethods := len(methods)
```

## Exercise

File: `04-collections-1/03-len/main.go`

Add missing functions: `NumberOfColors` and `NumberOfSystems`. Both should take no arguments and return an integer.
They should return the length of `colors` and `systems`.


<div class="accordion" id="hints-accordion">

<div class="accordion-item">
	<h3 class="accordion-header" id="hints-accordion-header-1">
	<button class="accordion-button fs-4 fw-semibold collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#hints-accordion-body-1" aria-expanded="false" aria-controls="hints-accordion">
		Hint #1
	</button>
	</h3>
	<div id="hints-accordion-body-1" class="accordion-collapse collapse" aria-labelledby="hints-accordion-header-1" data-bs-parent="#hints-accordion">
	<div class="accordion-body">

Don't forget the return value type.

```go
func NumberOfColors() int {
	// ...
}
```

</div>
	</div>
	</div>

</div>
