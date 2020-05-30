package main

import (
	"fmt"
	"sort"
)

func main() {

	// A := []int{10, 2, 5, 1, 8, 20}
	A := []int{3, 3, 5}
	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {

	if len(A) < 3 {
		return 0
	}

	sort.Ints(A)
	for i := 0; i < len(A)-2; i++ {
		if A[i]+A[i+1] > A[i+2] {
			return 1
		}
	}
	return 0
}
