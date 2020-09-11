package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	str1 := "20"

	num1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("num1:", num1)

	str2 := "3600s"
	dur2, err := time.ParseDuration(str2)
	if err != nil {
		fmt.Println("ParseDuration err:", err)
	}
	fmt.Println(dur2)

}
