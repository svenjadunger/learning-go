package main

import "fmt"

var price = 40.0

func main() {
	tax := taxRate * price
	fmt.Println("tax:", tax)
}
