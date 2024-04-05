package main

import "fmt"

func main() {
	rect := Rectangle{
		Width:  10,
		Length: 5,
	}

	area := rect.Area()

	fmt.Println("Area:", area)
}

func (rect Rectangle) Area() int {
return rect.Width*rect.Length
}

type Rectangle struct {
	Width  int
	Length int
}
