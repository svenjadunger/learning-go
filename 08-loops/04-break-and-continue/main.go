package main

import (
	"fmt"
	"strings"
)

func CountCreatedEvents(events []string) int {
	count := 0

	for _, event := range events {
		if strings.HasSuffix(event, "_created") {
			count++ 
		} else if event == "client_deleted" {
			break 
		}
		
	}

	return count
}


func main() {
	events := []string{
		"product_created",
		"product_updated",
		"product_assigned",
		"order_created",
		"order_updated",
		"client_created",
		"client_updated",
		"client_refreshed",
		"client_deleted",
		"order_updated",
	}


	count := CountCreatedEvents(events)
	fmt.Println("Anzahl der 'created' Ereignisse:", count)
}

	

