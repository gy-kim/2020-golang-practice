package main

import "fmt"

func main() {

	A := []int{23171, 21011, 21123, 21366, 21013, 21367}
	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {
	
	if len(A) < 2 {
		return 0
	}

	max := 0
	minDate := A[0]

	for i := 1; i<len(A); i++ {
		current := A[i] - minDate
		if current > 0 {
			if max < current {
				max = current
			}
		} else {
			minDate = A[i]
		}
	}
	return max
}