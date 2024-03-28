# Break and Continue

There are ways to control the loop once it starts.

Use `break` to interrupt the loop immediately.

```go
found := false

for _, name := range names {
	if name == targetName {
		found = true
		break
	}
	
	// ... 	
}
```

Use `continue` to skip the current loop's iteration and go to the next.

```go
for _, status := range statuses {
	if status == "disabled" {
		continue
	}
	
	// ... 	
}
```

## Exercise

File: `08-loops/04-break-and-continue/main.go`

Fill in the `CountCreatedEvents` function.
It takes a list of events as an argument and returns the number of "created" events.

* Use a `for` loop and `range` to count events.
* To determine if an event's type is "created", check if it ends with `_created`. Use [`strings.HasSuffix`](https://pkg.go.dev/strings#HasSuffix).
* If the event's type doesn't match, use `continue` to skip the iteration.
* If the event's type is "deleted", stop the loop entirely using `break`.