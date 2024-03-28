# Map range

Similarly to slice ranges, you can use `range` to iterate over maps.

It also comes in two flavors. The one-value `range` goes over the *keys* in a map.

```go
interfaceColors := map[string]string{
	"alice": "blue",
	"bob":   "orange",
	"emma":  "green",
}

for person := range interfaceColors {
	color := interfaceColors[person]
	fmt.Println(person, "uses", color, "interface")
}
```

The two-values `range` goes over *key-value pairs* in a map. 

```go
for person, color := range interfaceColors {
	fmt.Println(person, "uses", color, "interface")
}
```

A key difference between slices and maps is that maps are not ordered.

Calling `range` on a map returns the values in a random order, so you can't rely on it.

## Exercise

File: `08-loops/03-map-range/main.go`

Implement the missing `Keys` and `Values` functions. Both should accept a map of integers to strings.

* `Keys` should return all keys in the map: IDs of the products (a slice of integers).
* `Values` should return all values in the map: names of the products (a slice of strings).