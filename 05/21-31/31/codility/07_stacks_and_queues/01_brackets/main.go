package main

import "fmt"

func main() {

	S := "{[()()]}"
	result := Solution(S)

	fmt.Println("result:", result)

}

func Solution(S string) int {

	if len(S) == 0 {
		return 1
	}

	chars := map[string]int{"{": -3, "[": -2, "(": -1, ")": 1, "]": 2, "}": 3}

	stack := []int{}

	for _, v := range S {
		ch := string(v)
		num := chars[ch]

		if num < 0 {
			stack = append(stack, num)
		} else {
			if len(stack) == 0 {
				return 0
			}

			lastNum := stack[len(stack)-1]

			if (lastNum + num) != 0 {
				return 0
			}

			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return 0
	}

	return 1
}
