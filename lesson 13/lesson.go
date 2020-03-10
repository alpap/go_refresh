package main

import (
	"runtime"
	"sync"
)

func main() {
	// routines
	go routine()
	goSupportsClosures()
	correctGoRoutine()
	withWaitGroup()
	usingMutex()
}

// by putting go in front it spins off a green thread
// this creates an abstraction of a routine
// then its getting scheduled by the runtime
func routine() {
	println("hello")
}

func goSupportsClosures() {
	var msg = "closure"
	go func() {
		println(msg)
	}()
	// the above method creates race contitions because the value will change before the completion of the routine
	msg = "salam"
}

func correctGoRoutine() {
	go func(message string) {
		println(message)
	}("correct way")
}

var wg = sync.WaitGroup{}

func withWaitGroup() {
	wg.Add(1) // adds one to the wait counter
	go func(message string) {
		println(message)
		wg.Done() // decrements one from the wait counter
	}("correct wait way")
	wg.Wait() // wait till counter is empty
}

// synchronize routines using a mutex

var m = sync.RWMutex{}
var counter = 0

func usingMutex() {
	print(runtime.GOMAXPROCS(-1)) // if we pass -1 the values don't change
	//by default golang uses as many threads as there are cores on the machine -if we change this value for ex to 1 then the program will run single threaded
	println(" cores")
	for i := 0; i <= 10; i++ {
		m.RLock()
		go printNumber()
		m.Lock()
		go incrementNumber()
	}
	wg.Wait()
}

func printNumber() {
	wg.Add(1)
	// m.RLock() // if i want serial data the locks should be outside the routine in a single context
	println(counter)
	m.RUnlock()
	wg.Done()
}

func incrementNumber() {
	wg.Add(1)
	// m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}

// best practices
// don't create go routines if you are working with a library except if necessary
// go routines should always end
// check for race conditions at compile time run go run -race or go build -race
