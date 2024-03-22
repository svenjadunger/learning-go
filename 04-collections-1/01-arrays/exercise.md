# Arrays

Arrays keep a set of variables of the same type. They have a fixed size defined within brackets.

```go
var contactMethods [3]string
```

You can initialize an array when declaring it:

```go
contactMethods := [3]string{"email", "phone", "sms"}
```

You can reference array items by index (starting from `0`). Follow the array name with an index within brackets.

```go
// Retrieve array item
email := contactMethods[0]

// Set array item
contactMethods[1] = "mail"
```

## Exercise

File: `04-collections-1/01-arrays/main.go`

Extend the `roles` array with two more roles: `moderator` and `admin`.

Notice that the comma is required after each element when an array is defined across multiple lines.