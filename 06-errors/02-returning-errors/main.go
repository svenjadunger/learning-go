package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(100, 50)
	fmt.Println("Result:", result, "Error:", err)
}

func Divide(firstNumber float64, secondNumber float64) (float64,error) {
if  secondNumber==0 {
	return 0, errors.New("false")
}

return firstNumber/secondNumber, nil

}