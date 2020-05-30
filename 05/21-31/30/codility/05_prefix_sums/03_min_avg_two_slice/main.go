package main

import "fmt"

func main() {
	A := []int{4, 2, 2, 5, 1, 5, 8}

	result := Solution(A)

	fmt.Println("result:", result)
}

func Solution(A []int) int {

	if len(A) < 2 {
		return -1
	}
	if len(A) == 2 {
		return 0
	}

	minIdx := 0
	minAvg := float64(0)

	for i := 0; i < len(A)-1; i++ {
		twoSum := A[i] + A[i+1]
		avg2 := float64(twoSum) / 2

		if i == 0 {
			minAvg = avg2
		}

		if avg2 < minAvg {
			minAvg = avg2
			minIdx = i
		}

		if i < len(A)-2 {
			avg3 := float64(twoSum+A[i+2]) / 3
			fmt.Println("i", i, "avg3:", avg3, "minAvg:", minAvg)

			if avg3 < minAvg {
				minAvg = avg3
				minIdx = i
			}
		}
	}
	return minIdx
}
