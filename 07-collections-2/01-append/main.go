package main

var users []string

func main() {
	AddUser("Alice")
	AddUser("Bob")
}

func AddUser(u string) {
	users = append(users, u)
}


