package main

import (
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// create channel
	ch := make(chan int)

	// create a wait routine
	wg.Add(2)
	go func() {
		i := <-ch
		println(i)
		wg.Done()
	}()

	go func() {
		ch <- 42
		wg.Done()
	}()

	wg.Wait()

}
