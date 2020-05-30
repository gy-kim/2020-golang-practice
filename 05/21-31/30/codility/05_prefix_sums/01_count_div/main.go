package main

import "fmt"

func main() {

	A := 6
	B := 11
	K := 2

	result := Solution(A, B, K)

	fmt.Println("result:", result)
}

func Solution(A int, B int, K int) int {

	if A == 0 {
		return B/K + 1
	}

	return B/K - (A-1)/K
}
