package main

import "fmt"

func main() {
	X := 10
	Y := 83
	D := 30

	result := Solution(X, Y, D)

	fmt.Println("result:", result)
}

func Solution(X int, Y int, D int) int {
	if X == Y {
		return 0
	}

	Y -= X
	cnt := Y / D
	if Y%D != 0 {
		cnt++
	}
	return cnt

}
