package main

import "fmt"

var allocatedBuffers int 

func AllocateBuffer() *string {
	if allocatedBuffers >= 3 {
	return nil
	
}

allocatedBuffers++

var buffer string
	return &buffer
}


func main() {
	var buffers []*string

	for {
		b := AllocateBuffer()
		if b == nil {
			break
		}

		buffers = append(buffers, b)
	}

	fmt.Println("Allocated", len(buffers), "buffers")
}

