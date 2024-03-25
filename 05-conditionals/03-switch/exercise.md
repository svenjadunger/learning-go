# Switch

Sometimes, you need to write a long chain of `if-else` statements.

```go
if x == 0 {
	// Zero
} else if x == 1 { 
	// One
} else if x == 2 { 
	// Two
} else { 
	// Other
}
```

Use the `switch` statement to make this more readable and short.

```go
switch x {
case 0:
	// Zero	
case 1: 
	// One	
case 2: 
	// Two	
default: 
	// Other	
}
```

## Exercise

Fill in the `Direction` function.

The `d` arguments is a world direction. Make the function return:

* `north` for `N`,
* `east` for `E`,
* `south` for `S`,
* `west` for `W`,
* `unknown` for any other value.