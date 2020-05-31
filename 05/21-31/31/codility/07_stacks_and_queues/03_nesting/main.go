package main

import "fmt"

func main() {
	S := "(()(())())"

	result := Solution(S)
	fmt.Println("result:", result)

}

func Solution(S string) int {
	if len(S) == 0 {
		return 1
	}
	chars := map[string]int{"(": -1, ")": 1}
	stack := []int{}

	for _, v := range S {
		ch := string(v)
		num := chars[ch]

		if num < 0 {
			stack = append(stack, num)
			continue
		}
		lastNum := stack[len(stack)-1]
		if (num + lastNum) != 0 {
			return 0
		}
		stack = stack[:len(stack)-1]
	}

	if len(stack) != 0 {
		return 0
	}
	return 1
}
