package main

import "fmt"

func main() {

	A := []int{3, 2, -6, 4, 0}

	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {

	if len(A) == 1 {
		return A[0]
	}

	maxSum := A[0]
	partSum := A[0]

	for i := 1; i < len(A); i++ {
		compare := partSum + A[i]
		if A[i] > compare {
			compare = A[i]
		}

		if maxSum < compare {
			maxSum = compare
		}

		if compare < 0 {
			partSum = 0
		} else {
			partSum = compare
		}
	}
	return maxSum
}
