package main

import "fmt"

func main() {
	X := 5
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}

	result := Solution(X, A)
	fmt.Println("result:", result)
}

func Solution(X int, A []int) int {

	list := make(map[int]bool)

	if len(A) == 0 || len(A) < X {
		return -1
	}

	for idx, i := range A {
		fmt.Println("list", list, "idx:", idx, "i:", i)
		if _, ok := list[i]; !ok {
			list[i] = true
		}
		if len(list) == X {
			return idx
		}
	}

	return -1
}
