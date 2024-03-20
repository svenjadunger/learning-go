# Main

Let's take another look at Go code structure. 

## Packages

The top line of every Go file starts with the `package` declaration.

`main` is a special package — this is where your program begins.

```go
package main
```

## The main function

The entry point to your application is always the `main` function in the `main` package.

```go
func main() { 
	// Your code goes here	
}
```

## Exercise

File: `01-start/02-main/main.go`

There are two mistakes you need to fix.

Can you see what needs to be changed to print `Hello, World!`, just like the previous program?
