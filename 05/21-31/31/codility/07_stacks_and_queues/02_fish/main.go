package main

import "fmt"

func main() {
	A := []int{4, 3, 2, 1, 5}
	B := []int{0, 1, 0, 0, 0}

	result := Solution(A, B)
	fmt.Println("result:", result)
}

func Solution(A []int, B []int) int {
	aliveCnt := 0
	down := []int{}
	size := len(A)

	for i := 0; i < size; i++ {
		aliveCnt++
		if B[i] == 1 {
			down = append(down, i)
			continue
		}
		if len(down) == 0 {
			continue
		}

		lastDownFish := down[len(down)-1]

		for {
			aliveCnt--

			if A[lastDownFish] < A[i] {
				down = down[:len(down)-1]
				if len(down) == 0 {
					break
				}
				lastDownFish = down[len(down)-1]
			} else {
				break
			}
		}
	}
	return aliveCnt
}
