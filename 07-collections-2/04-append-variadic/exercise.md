# Variadic Append

## Selecting elements

You can select a part of a slice by specifying two indexes.
It will select elements between the indexes, omitting the last one.

```go
letters := []string{"A", "B", "C", "D", "E", "F"}

fmt.Println(letters[2:4]) // C, D
```

If you omit the first index, it defaults to `0`.

```go
fmt.Println(letters[0:4]) // [A B C D]
fmt.Println(letters[:4]) // [A B C D]
```

If you omit the last index, it defaults to the number of elements (the slice's length).

```go
fmt.Println(letters[4:len(letters)]) // [E F]
fmt.Println(letters[4:]) // [E F]
```

If you omit both indexes, you get the entire slice.

```go
fmt.Println(letters[0:len(letters)]) // [A B C D E F]
fmt.Println(letters[:]) // [A B C D E F]
```

You can use `append` together with indexing to do other operations on slices, like removing elements:

```go
withoutThirdElement := append(names[:2], names[3:]...)
```

Or changing the order:

```go
middle := len(names)/2
swappedInTheMiddle := append(names[middle:], names[:middle]...)
```

## Exercise

File: `07-collections-2/04-append-variadic/main.go`

Write three functions that help to work with slices:

* `Merge` - extends the first slice with elements from the second one.
* `Remove` - removes one element from the slice at a given index.
* `RemoveLast` - removes the last element from a given slice. Use `Remove` to keep the logic in one place!


<div class="accordion" id="hints-accordion">

<div class="accordion-item">
	<h3 class="accordion-header" id="hints-accordion-header-1">
	<button class="accordion-button fs-4 fw-semibold collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#hints-accordion-body-1" aria-expanded="false" aria-controls="hints-accordion">
		Hint #1
	</button>
	</h3>
	<div id="hints-accordion-body-1" class="accordion-collapse collapse" aria-labelledby="hints-accordion-header-1" data-bs-parent="#hints-accordion">
	<div class="accordion-body">

The index of the last element of a slice is the slice's length minus one.

```go
lastElementIndex := len(slice)-1
```

</div>
	</div>
	</div>

</div>
