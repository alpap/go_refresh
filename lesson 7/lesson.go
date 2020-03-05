package main

import "fmt"

func main() {
	// simple loops
	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("(i => %v)\n", i)
	}

	// multiple variables
	for i, j := 0, 0; i < count; i, j = i+1, j+1 {
		fmt.Printf("i => %v   ", i)
		fmt.Printf("j => %v \n", j)
	}

	counter := 0 // scoped to the main function

	for ; counter <= 100; counter++ {
		fmt.Printf("%v   ", counter)
	}
	// do while

	for counter > 0 {
		counter--
		fmt.Printf("%v   ", counter)
	}

	// run forever
	for {

		counter++
		fmt.Printf("%v   ", counter)
		if counter == 5 {
			break // get out of the loop
			// continue // continue the loop
		}
	}

	// using a label to break out to
Loop:
	for {
		for {
			counter--
			fmt.Printf("%v   ", counter)
			if counter == 0 {
				break Loop // get out of the loop to the label
			}
		}
	}
	fmt.Println()
	// working with for loops and collections
	s := []int{6, 7, 8}
	for k, v := range s { // you can skip the keys by putting an underscore ex _,v

		fmt.Println(k, v)
	}

}
