package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// output := strings.Join(listFiles("testdata"))

	// fmt.Println(output)

	output :=listFiles("testdata")
	fmt.Println(strings.Join(output," "))
}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := os.ReadDir(dirname)
	 
	if err!= nil {
		log.Fatal(err)
	}

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
