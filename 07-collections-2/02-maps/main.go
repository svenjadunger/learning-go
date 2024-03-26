package main

import "fmt"

var (
	Stats = map[string]int{}
)

func CreateUser(user string) {
	fmt.Println("Creating user", user)
}

func UpdateUser(user string) {
	fmt.Println("Updating user", user)
}

func PurgeStats() {
}
