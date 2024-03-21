# HTTP Server Project

You're now going to start a new *project*. You'll improve it over a few exercises.

Some syntax will be new to you. Don't worry about it yet.
For now, the point is to get you familiar with the general structure of an HTTP server.

Follow the instructions and type code from the snippets. Eventually, you'll grasp all of it.

## The HTTP server

Go comes with a production-ready HTTP library. It's straightforward to run your own website or web application.

To start an HTTP server, run in your `main` function:

```go
http.HandleFunc("/your-endpoint", yourFunction)

log.Fatal(http.ListenAndServe(":8080", nil))
```

You need to import `net/http` and `log` packages.
Instead of writing one `import` statement per line, you can create a group:

```go
import (
	"net/http"
	"log"
)
```

Then, define the handler function:

```go
func yourFunction(w http.ResponseWriter, r *http.Request) {
	// Handler function body
}
```

To return a response, use `fmt.Fprint` with `http.ResponseWriter` as the first argument: 

```go
func yourFunction(w http.ResponseWriter, r *http.Request) {
	name := "Peter"
    fmt.Fprint(w, "Hello, my name is ", name)
}
```

The server is started on the port you pass to `http.ListenAndServe`. It will run forever, until you close the application (by pressing Ctrl+C).
In the example above, it will be available at [http://localhost:8080](http://localhost:8080).

See detailed documentation at [pkg.go.dev](https://pkg.go.dev/net/http).

## Exercise

File: `project-http-hello/main.go`

Implement an HTTP server that runs on port 8080 and returns `Hello, [name]` from the `/hello` endpoint with `[name]` provided by the `name` query parameter.

For example, visiting `http://localhost:8080/hello?name=Emma` should display `Hello, Emma`.

To get the `name` parameter, use:

```go
name := r.URL.Query().Get("name")
```

You don't need to define the query parameter when calling `HandleFunc`.

The server should be running on port `8080` (like in the example above).
To see your server in action, visit [http://localhost:8080/hello](http://localhost:8080/hello) in your browser or in the terminal using `curl`.

Add a `?name=` parameter with a name, and you should see a greeting.
