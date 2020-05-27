package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func worker2() <-chan bool {
	done := make(chan bool)

	go func() {
		defer close(done)
		fmt.Println("working2...")
		time.Sleep(time.Second)
		fmt.Println("done")
	}()

	return done
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done

	done2 := worker2()
	<-done2
}
