# The "ls" Project: Errors

*This exercise continues your previous "ls" project. You will work on the same files.*

## Exercise

File: `project-ls/main.go`

`os.ReadDir` returns two values, the second one is an `error`. The current version ignores it.

```go
files, _ := os.ReadDir(dirname)
```

Assign this error to an `err` variable and handle with `log.Fatal`.

```go
if err != nil {
	log.Fatal(err)
}
```

[`log.Fatal`](https://pkg.go.dev/log#Fatal) prints out the message and stops the application.
It's a way to handle errors from which the application can't recover.
