package main

import "fmt"

func main() {
	// primitives

	// an empty variable is always initialized
	var emptyBool bool
	fmt.Printf("%v, %T \n", emptyBool, emptyBool)

	var boolData bool = true
	fmt.Printf("%v, %T \n", boolData, boolData)

	// equals operator

	result := 1 == 2
	fmt.Printf("%v, %T \n", result, result)

	// int types
	// int8
	// int16
	// int34
	// int64
	// every signed integer has a unsigned integer

	var uns uint16 = 33
	fmt.Printf("%v, %T \n", uns, uns)

	// arithmetic
	// + - * / %

	a := 10
	b := 3
	// gives an int because it an integer division
	fmt.Printf("%v, %T \n", a/b, a/b)

	// var c int = 3
	// var d uint = 2
	// error we cannot add two different types
	// fmt.Printf("%v, %T \n", c+d, c+d)

}
