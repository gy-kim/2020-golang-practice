package main

import "fmt"

func main() {

	N := 5
	A := []int{3, 4, 4, 6, 1, 4, 4}

	result := Solution(N, A)
	fmt.Println("result:", result)
}

func Solution(N int, A []int) []int {
	arr := make([]int, N)
	max, base := 0, 0

	for _, x := range A {
		if x == N+1 {
			base = max
			continue
		}
		cnt := arr[x-1]
		if cnt < base {
			cnt = base
		}
		cnt++
		arr[x-1] = cnt

		if max < arr[x-1] {
			max = arr[x-1]
		}
	}

	for idx, v := range arr {
		if v < base {
			arr[idx] = base
		}
	}

	return arr
}
