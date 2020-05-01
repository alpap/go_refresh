package main

import (
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// exampleOne()
	// exampleTwo()
	exampleThree()
}

func exampleOne() {
	// create channel
	ch := make(chan int)

	// create a wait routine
	wg.Add(2)
	go func() {
		// get data from the channel
		i := <-ch
		println(i)
		wg.Done()
	}()

	go func() {
		// put data to the channel
		ch <- 42
		wg.Done()
	}()

	wg.Wait()
}

//buffered channel
func exampleTwo() {
	// create channel
	// we make this into a buffer to handle multiple inputs into the channel
	ch := make(chan int, 100)

	// create a wait routine
	wg.Add(2)
	// only a receiving
	go func(ch <-chan int) {
		// get data from the channel
		i := <-ch
		println(i)
		i = <-ch
		println(i)
		wg.Done()
	}(ch)

	wg.Add(2)
	go func(ch chan<- int) {
		// put data to the channel
		ch <- 42
		ch <- 23
		wg.Done()
	}(ch)
	// only sending

	wg.Wait()
}

// run as long as there is data
func exampleThree() {
	// create channel
	// we make this into a buffer to handle multiple inputs into the channel
	ch := make(chan int, 100)

	// create a wait routine
	wg.Add(1)
	// only a receiving
	go func(ch <-chan int) {
		// get data from the channel until the channel closes
		// else it will loop forever and cause a deadlock
		for i := range ch {
			println(i)
		}
		wg.Done()
	}(ch)

	wg.Add(1)
	// only sending
	go func(ch chan<- int) {
		// put data to the channel
		for i := 0; i < 50; i++ {
			ch <- i
		}
		// if we want to close the channel
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()

	// We can get another value from the channel which is boolean and signifies if the channel is open
	// if i,ok := <-ch ;ok {}
}

func exampleFour() {
	// create channel
	// we make this into a buffer to handle multiple inputs into the channel
	ch := make(chan int, 100)

	// create a wait routine
	wg.Add(1)
	// only a receiving
	go func(ch <-chan int) {
		// get data from the channel until the channel closes
		// else it will loop forever and cause a deadlock
		for i := range ch {
			println(i)
		}
		wg.Done()
	}(ch)

	wg.Add(1)
	// only sending
	go func(ch chan<- int) {
		// put data to the channel
		for i := 0; i < 50; i++ {
			ch <- i
		}
		// if we want to close the channel
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()

	// We can get another value from the channel which is boolean and signifies if the channel is open
	// if i,ok := <-ch ;ok {}
}

// signal only channels are channels that don't allocate memory and are only used to notify that something is send
var ch = make(chan int)
var doneCh = make(chan struct{})

//select

func selectChannel() {
	doneCh <- struct{}{}
	for {
		// select will wait till the selection is done
		select {
		case entry := <-ch:
			print(entry)
		case <-doneCh:
			break

		}
	}
}

// basic for channel monotoring
