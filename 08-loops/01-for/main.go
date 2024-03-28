package main

import "fmt"

func Alphabet(length int) []string {

	var result []string

	for i:= 0; i < length; i++ {
		letter := characterByIndex(i)
		result = append(result, letter)

		
	}
	return result
}

func main() {
	alphabet := Alphabet(26)
	fmt.Println(alphabet)
}

func characterByIndex(i int) string {
	return string(rune('a' + i))
}
