package main

import (
	"fmt"
	"sort"
)

func main() {

	A := []int{2, 1, 1, 2, 3, 1}
	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {

	arr := make(map[int]bool)

	sort.Ints(A)
	for _, v := range A {
		if _, ok := arr[v]; !ok {
			arr[v] = true
		}
	}

	return len(arr)
}
