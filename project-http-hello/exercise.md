# HTTP Server Stats Project

*This exercise continues your previous HTTP project. You will work on the same files.*

We want to add some statistics to our HTTP server.
For now, let's just print them out.

After calling the endpoints:

```bash
GET /hello?name=Bill
GET /hello?name=Eva
GET /hello?name=Bill
```

We expect to have output like:

```bash
calls: []string{"Bill"}
stats: map[string]int{"Bill":1}

calls: []string{"Bill", "Eva"}
stats: map[string]int{"Eva":1, "Bill":1}

calls: []string{"Bill", "Eva", "Bill"}
stats: map[string]int{"Eva":1, "Bill":2}
```

## Exercise

File: `project-http-hello/main.go`

Use global variables to store `calls` and `stats`.
Normally, using global variables is not a good idea, but in this case, we'll stick to a simple approach to have something working.

Store `calls` in a slice of strings (`[]string`) and `stats` in a map of strings to integers (`map[string]int`).

This approach isn't thread safe. We will take care of that later.

For printing of the values you can use `fmt.Printf` with the `%#v` placeholder.
It prints out the contents of the variable together with its type information.

```go
fmt.Printf("calls: %#v\n", calls)
fmt.Printf("stats: %#v\n\n", stats)
```

Keep in mind that maps in Go are not ordered. Printing maps is not deterministic as well.