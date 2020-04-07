package main

func main() {
	// defer -> run the statement right before the function exits
	// if we defer all functions then we have a LIPO last in first out that means that the order of execution will be reversed
	println("start")
	defer println("middle")
	println("end")
	// panic()
	// we use this to indicate that the application cant run without the resource specified
	// panics happen after defers
	// defer panic return value
	panicking()
	println("ok")
}

func panicking() {

	// Recover

	// this should be before panic else it exits before it has the chance to declate the defer function
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}()

	panic("time to panic")

	// this doesn't get executed cause we paniced that means that we wont be able to continue operation so the functions exits
	a, b := 1, 0
	println(a)
	println(b)
	ans := a / b //this will panic you cant divide with 0
	println(ans)

}
