package main

import "time"

func main() {
	// routines
	go routine()
	goSupportsClosures()
	time.Sleep(100 * time.Millisecond)
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
}
