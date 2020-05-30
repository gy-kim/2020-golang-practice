package main

import "fmt"

func main() {
	A := []int{0, 1, 0, 1, 1}

	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {

	result := 0
	zeros := 0

	for _, v := range A {
		if v == 0 {
			zeros++
		}
		if v == 1 {
			result += zeros
		}

		if result > 1000000000 {
			return -1
		}
	}
	return result
}
