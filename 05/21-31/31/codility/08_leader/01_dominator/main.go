package main

import "fmt"

func main() {

	A := []int{3, 4, 3, 2, 3, -1, 3, 3}

	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {
	size := len(A)
	if size == 0 {
		return -1
	}
	if size == 1 {
		return 0
	}

	condition := (size / 2) + 1

	m := map[int]int{A[0]: 1}

	for i := 1; i < size; i++ {
		v := A[i]

		if cnt, ok := m[v]; ok {
			cnt++
			m[v] = cnt
		} else {
			m[v] = 1
		}

		if m[v] >= condition {
			return i
		}
	}

	return -1
}
