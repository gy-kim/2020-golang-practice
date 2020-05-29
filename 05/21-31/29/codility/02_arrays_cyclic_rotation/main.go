package main

import "fmt"

func main() {
	A := []int{3, 8, 9, 7, 6}
	K := 3

	result := Solution(A, K)
	fmt.Println("result:", result)
}

func Solution(A []int, K int) []int {
	if len(A) == 0 || K == 0 || len(A) == K {
		return A
	}

	if len(A) < K {
		K = K % len(A)
	}

	return append(A[len(A)-K:], A[:len(A)-K]...)
}
