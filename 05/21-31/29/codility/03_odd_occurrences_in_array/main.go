package main

import "fmt"

func main() {

	A := []int{9, 3, 9, 3, 9, 7, 9}

	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {
	list := make(map[int]interface{})
	for _, v := range A {
		if _, ok := list[v]; ok {
			delete(list, v)
			continue
		}
		list[v] = true
	}
	for k := range list {
		return k
	}

	return -1
}
