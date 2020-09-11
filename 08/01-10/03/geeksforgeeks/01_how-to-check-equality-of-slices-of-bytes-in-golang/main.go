package main

import (
	"bytes"
	"fmt"
)

// https://www.geeksforgeeks.org/how-to-check-equality-of-slices-of-bytes-in-golang/

func main() {
	slice_1 := []byte{'A', 'N', 'M', 'A', 'P', 'A', 'A', 'W'}
	slice_2 := []byte{'A', 'N', 'M', 'A', 'P', 'A', 'A', 'W'}

	res := bytes.Equal(slice_1, slice_2)
	if res == true {
		fmt.Println("Slice_1 is equal to Slice_2")
	} else {
		fmt.Println("Slice_1 is not equal to slice_2")
	}
}
