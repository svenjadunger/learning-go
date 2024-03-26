# Maps

Use maps to store key-value pairs. Map declaration uses brackets with the keys type inside it, followed by the values type.

```go
// A map of strings to ints
var limits map[string]int
```

A declared map is equal to `nil`. Accessing it in this state will crash your application.

Maps have to be initialized using `make`:

```go
limits = make(map[string]int)
```

Or using the `:=` syntax and a *map literal*:

```go
limits := map[string]int{}
```

You can assign values to initialized maps:

```go
limits["create"] = 10
```

You can fill maps with values when initializing them:

```go
limits := map[string]int{
	"create": 10,
	"update": 15,
	"delete": 1,
}
```

To retrieve values, use a syntax similar to accessing elements of a slice:

```go
update := limits["update"]
```

If a key doesn't exist, the map returns a zero value, depending on the type.

It's often important to know if a value exists at all. You can use a special two-value syntax for this.

```go
update, ok := limits["update"]
if !ok { 
	// update limits not found
}

fmt.Println("Updates limited to", update)
```

If you just need to check if a key exists, you can skip the value altogether using the underscore (`_`).

```go
_, ok := limits["updates"]
if !ok { 
	// update limits not found
}
```

To delete a key-value pair, use the `delete` keyword. If the key doesn't exist, it has no effect.

```go
delete(limits, "update")
```

## Exercise

File: `07-collections-2/02-maps/main.go`

There are two functions we use to create and update users: `CreateUser` and `UpdateUser`.

We'd like to track how many times each of them is called.

Use the global `Stats` map to track these numbers. Increase counts using keys `"create"` and `"update"` in the two functions.

Finally, fill the `PurgeStats` function that should clear all saved statistics.


<div class="alert alert-dismissible bg-light-primary d-flex flex-column flex-sm-row p-7 mb-10">
    <div class="d-flex flex-column">
        <h3 class="mb-5 text-dark">
			<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-lightbulb text-primary" viewBox="0 0 16 16">
			  <path d="M2 6a6 6 0 1 1 10.174 4.31c-.203.196-.359.4-.453.619l-.762 1.769A.5.5 0 0 1 10.5 13a.5.5 0 0 1 0 1 .5.5 0 0 1 0 1l-.224.447a1 1 0 0 1-.894.553H6.618a1 1 0 0 1-.894-.553L5.5 15a.5.5 0 0 1 0-1 .5.5 0 0 1 0-1 .5.5 0 0 1-.46-.302l-.761-1.77a1.964 1.964 0 0 0-.453-.618A5.984 5.984 0 0 1 2 6zm6-5a5 5 0 0 0-3.479 8.592c.263.254.514.564.676.941L5.83 12h4.342l.632-1.467c.162-.377.413-.687.676-.941A5 5 0 0 0 8 1z"/>
			</svg>
			Tip
		</h3>
        <span>

To increase or decrease numbers, you can use a few operators:

```go
x := 0

// All increase x by 1
x = x + 1
x += 1
x++

// All decrease x by 1
x = x - 1
x -= 1
x--
```

</span>
	</div>
	</div>
