package main

import (
	"errors"
)

var message = ""

func StoreMessage(m string) error {
	if m == "" {
		return errors.New("no message")
	}

	message = m

	return nil
}

func MustStoreMessage(message string) {
	err := StoreMessage(message)

	if err != nil {
		panic(err)
	}
}

func main() {
	MustStoreMessage("Hello!")
}
