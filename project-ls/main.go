package main

import (
	"os"
)

func main() {
	// Your solution goes here. Good luck!
}

func listFiles(dirname string) []string {
	var dirs []string

	files, _ := os.ReadDir(dirname)

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
