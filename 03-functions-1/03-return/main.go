package main

import "fmt"

func main() {
	result := Sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}

func Sum(x int,y int, z int, h int, s int) int {
	sum:= x+y+z+h+s
	return sum
}