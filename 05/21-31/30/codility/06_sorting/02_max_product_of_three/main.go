package main

import (
	"fmt"
	"sort"
)

func main() {

	A := []int{-5, 5, -5, 4}
	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {
	sort.Ints(A)

	result1 := A[len(A)-3] * A[len(A)-2] * A[len(A)-1]

	if A[0] < 0 && A[1] < 0 && A[len(A)-1] >= 0 {
		result2 := A[0] * A[1] * A[len(A)-1]
		if result1 < result2 {
			return result2
		}
	}
	return result1
}
