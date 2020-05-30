package main

import "fmt"

func main() {

	A := []int{3, 1, 2, 4, 3}

	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {
	min := 0
	list := make(map[int]int, len(A))

	for i := 0; i < len(A); i++ {
		if i == 0 {
			list[i] = A[i]
			continue
		}
		list[i] = list[i-1] + A[i]
	}

	fmt.Println("list", list)
	for i := 0; i < len(list)-1; i++ {
		left := list[i]
		right := list[len(list)-1] - left
		fmt.Println("left:", left, "right:", right)
		gap := left - right
		if gap < 0 {
			gap *= -1
		}
		fmt.Println("gap:", gap)
		if i == 0 {
			min = gap
			continue
		}
		if min > gap {
			min = gap
		}
	}

	return min
}
