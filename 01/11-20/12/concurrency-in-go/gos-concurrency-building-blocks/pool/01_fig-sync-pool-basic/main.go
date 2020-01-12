package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	fmt.Println("first")
	myPool.Get()
	fmt.Println("second")
	instance := myPool.Get()
	myPool.Put(instance)
	fmt.Println("third")
	myPool.Get()
}
