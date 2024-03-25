# The "ls" Project

The standard library provides the `strings` package that exposes functions to manipulate strings.
For example, [`strings.Join`](https://pkg.go.dev/strings#Join).
Click the link and take a look at what other functions there are.

If you want to check a function from the standard library, visit [pkg.go.dev](https://pkg.go.dev/).
You can also look up definitions directly in your IDE.

## Exercise

File: `project-ls/main.go`

Your task is to implement a simple `ls` command that lists the contents of a directory.
For now, you need to list just the `testdata` directory.

There's a ready to use `listFiles` function that returns a list of files from a directory.
You need to call it with a proper argument.

`listFiles` returns file names as a slice of strings.
You need to print the returned list in one line in the following format:

```bash
backup.zip contacts.csv cv.doc photos
```

File names come correctly sorted. Use `strings.Join` to join the files with a space.