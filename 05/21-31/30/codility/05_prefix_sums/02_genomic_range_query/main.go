package main

import (
	"fmt"
)

func main() {

	S := "CAGCCTA"
	P := []int{2, 5, 0}
	Q := []int{4, 5, 6}
	result := Solution(S, P, Q)
	fmt.Println("result:", result)

}

func Solution(S string, P []int, Q []int) []int {

	arr := make([][]int, len(S))
	size := len(P)
	result := make([]int, size)

	for i := 0; i < len(S); i++ {
		row := make([]int, 4)
		c := string(S[i])

		switch c {
		case "A":
			row[0] = 1
		case "C":
			row[1] = 1
		case "G":
			row[2] = 1
		case "T":
			row[3] = 1
		}

		arr[i] = row
		fmt.Println(arr)
	}

	fmt.Println()

	for i := 1; i < len(S); i++ {
		for j := 0; j < 4; j++ {
			arr[i][j] += arr[i-1][j]
		}
		fmt.Println(arr)
	}

	for i := 0; i < size; i++ {
		minIdx := P[i]
		maxIdx := Q[i]

		for j := 0; j < 4; j++ {
			sub := 0
			if minIdx != 0 {
				sub = arr[minIdx-1][j]
			}
			fmt.Println("sub:", sub)
			if (arr[maxIdx][j] - sub) > 0 {
				result[i] = j + 1
				break
			}
		}
	}

	return result

	// dat := map[string]int{
	// 	"A": 1,
	// 	"C": 2,
	// 	"G": 3,
	// 	"T": 4,
	// }
	// size := len(P)

	// result := make([]int, size)

	// for i := 0; i < size; i++ {

	// 	fmt.Println("P:", P[i], "Q:", Q[i]+1)
	// 	slice := S[P[i] : Q[i]+1]
	// 	fmt.Println(slice)

	// 	min := 4
	// 	for _, s := range slice {
	// 		num := dat[string(s)]
	// 		if num == 1 {
	// 			min = 1
	// 			break
	// 		}
	// 		if num < min {
	// 			min = num
	// 		}
	// 	}

	// 	result[i] = min

	// 	min = 4
	// }

	// return result
}
