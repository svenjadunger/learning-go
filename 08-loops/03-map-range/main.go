package main

import "fmt"

var products = map[int]string{
	1: "Book",
	2: "Video Course",
	3: "Lecture",
	4: "Talk",
	5: "Training",
}

func main() {
	ids := Keys(products)
	names := Values(products)

	fmt.Println("Prouct IDs:", ids)
	fmt.Println("Product names:", names)
}

func Keys(products map[int]string) []int {
    var ids []int
    for id := range products {
        ids = append(ids, id)
    }
    return ids
}

func Values(products map[int]string) []string {
    var names []string
    for _, name := range products {
        names = append(names, name)
    }
    return names
}
