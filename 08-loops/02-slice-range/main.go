package main

import "fmt"

func main() {
	result := Sum(1, 2, 3, 4, 5)
    fmt.Println("Die Summe ist:", result)

    sum := []string{"a", "b", "c"}
    Print(sum)
}

func Print(strings []string) {
    for i, str := range strings {
        fmt.Println("Index", i, "Wert", str)
    }
}

func Sum(nums ...int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}

