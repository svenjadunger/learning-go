package main

import (
	"fmt"
	"os"
)

func CreateDirectory(name string) bool {
	os.Mkdir(name, 0755)
}

func main() {
	ok := CreateDirectory("my-directory")
	if ok {
		fmt.Println("Directory created")
	} else {
		fmt.Println("Failed to create directory")
	}
}
