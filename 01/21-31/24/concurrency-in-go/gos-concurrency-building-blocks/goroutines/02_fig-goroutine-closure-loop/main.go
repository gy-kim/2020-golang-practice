package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"Hello", "greeting", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}
