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

	var c int = 3
	var d uint = 2
	//error we cannot add two different types
	fmt.Printf("%v, %T \n", uint(c)+d, uint(c)+d)

	// bitwise operators
	e := 8
	f := 3
	fmt.Println(e & f)
	fmt.Println(e | f)
	fmt.Println(e ^ f)
	fmt.Println(e &^ f)

	// bit shifting
	// shift left or right the bits of the number
	fmt.Println(e << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(e >> 3) // 2^3 / 2^3 == 2^0 =1

	// floating points are float32 and float64 auto initialization is always float64
	// floating points have ho bitwise or bitshift operations
	g := 3.14
	h := 13.7E72 // or 13.7e72
	fmt.Printf("%v, %T \n", g, g)
	fmt.Printf("%v, %T \n", h, h)

	// complex type   complex64 complex128
	// parsed by a float32 + float32 or float64+float64 for the real and imaginary parts
	// complex types have ho bitwise or bitshift operations
	var com complex64 = 1 + 2i
	fmt.Printf("%v, %T \n", com, com)
	var comTwo complex64 = 1 + 2i
	fmt.Printf("%v, %T \n", com+comTwo, com+comTwo)
	// i we want to extract the real and the imaginary
	fmt.Printf("%v, %T \n", real(com), real(com))
	fmt.Printf("%v, %T \n", imag(com), imag(com))

	// text types

	z := "this is a string"
	fmt.Printf("%v, %T \n", z, z)

	// a string is a collection of chars so it will return the char in byte
	fmt.Printf("%v, %T \n", z[2], z[2])
	fmt.Printf("%v, %T \n", string(z[2]), z[2])

	// strings are immutable so you cant change a letter in the string with the bellow example
	//z[2]="b"

	// string concatenation
	fmt.Printf("%v, %T \n", z+z, z+z)

	// string to array of bytes else known as slice in go
	fmt.Printf("%v, %T \n", []byte(z), []byte(z))
	// i can convert the bytes back to letters using string(byteValue)

	// runes they represent UTF32 chars so they are of type int32
	// we declare a rune using single quotes

	r := 'a'
	fmt.Printf("%v, %T \n", r, r)

}
