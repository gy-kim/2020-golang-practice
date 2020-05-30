package main

import (
	"fmt"
	"sort"
)

func main() {

	A := []int{4, 1, 2}

	result := Solution(A)

	fmt.Println("result:", result)
}

func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}

	sort.Ints(A)

	if A[0] != 1 {
		return 0
	}

	if A[len(A)-1] != len(A) {
		return 0
	}

	for i := 0; i < len(A)-1; i++ {
		if A[i+1]-A[i] != 1 {
			return 0
		}
	}
	return 1
}
