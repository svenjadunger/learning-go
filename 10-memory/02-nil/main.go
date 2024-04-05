package main

var (
	Name     = "Alice"
	Password = "top-secret"
)

func UpdateUser(newName *string, newPassword *string) {
	Name = *newName
	Password = *newPassword
}

func main() {
	// Nothing to update
	UpdateUser(nil, nil)

	// Update password
	newPassword := "much-more-secure"
	UpdateUser(nil, &newPassword)

	// Update both
	newName := "Bob"
	UpdateUser(&newName, &newPassword)
}
