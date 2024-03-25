# HTTP Bad Request Project

*This exercise continues your previous HTTP project. You will work on the same files.*

By default, an HTTP handler returns the `200 OK` status when any content is written to `http.ResponseWriter`.

You can write custom status codes using `w.WriteHeader`. 
You need to do it before writing any content, otherwise it will have no effect.

See the list of HTTP codes on [pkg.go.dev](https://pkg.go.dev/net/http#pkg-constants).

### Returning from handlers

`w.WriteHeader` doesn't interrupt the execution of the handler function.

If you need to finish an HTTP request early, you need to explicitly use `return`.

```go
func hello(w http.ResponseWriter, r *http.Request) {
	// ...
	
	if thisShouldNotBeEmpty == "" {
		w.WriteHeader(http.StatusBadRequest)
		return // <- this is important!
	}
	
	// With no return in place, this would be executed
}
```

## Exercise

File: `project-http-hello/main.go`

Return `http.StatusBadRequest` if the `name` query param is empty (equal to `""`).
