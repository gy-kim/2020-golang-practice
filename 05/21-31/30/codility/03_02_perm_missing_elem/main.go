package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{2, 3, 1, 5}

	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {
	if len(A) == 0 {
		return 1
	}

	sort.Ints(A)
	for i := 0; i < len(A); i++ {
		if A[i] != i+1 {
			return i + 1
		}
	}
	return len(A) + 1
}
