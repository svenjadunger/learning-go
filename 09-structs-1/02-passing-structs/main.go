package main

import "fmt"

func main() {
	rect := Rectangle{
		Width:  10,
		Length: 5,
	}

	area := Area(rect)

	fmt.Println("Area:", area)
}

func Area(rect Rectangle) int {
return rect.Width*rect.Length
}

type Rectangle struct {
	Width  int
	Length int
}
