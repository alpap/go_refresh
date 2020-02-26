package main

import (
	"fmt"
)

func main() {
	// contants
	// naming convension
	// const follows the uppercase export standart
	const myConst int = 43
	fmt.Printf("%v, %T \n", myConst, myConst)
	// constants must be assigned at compile time so the bellow will give an error
	// const a float32 = math.Sin(1.57)
	// this happens cause this contant is assigned at runtime which is not allowed
	// shadowing is allowed with constants
	const b = 43 // int
	var c int16 = 27
	fmt.Printf("%v, %T \n", b+c, b+c)
	// the compiler check the usage of the b and then puts the type automatically so it matches the addition operation
	// Enumerated constants
	const (
		_     = iota // the _ symbol what it says to the program is that dont allocate memory to this value just throw it away
		enum         // gives the value 0
		enum2        // automatically applies the same formula so it will be 1
	)
	fmt.Printf("%v, %T \n", enum, enum)
	fmt.Printf("%v, %T \n", enum2, enum2)

	// enumerated expressions
	const (
		_  = iota
		KB = 1 << (10 * iota) // shifts the bit 0001 ten times to the left so 1*2^10-1024
		MB
		GB
	)
	fmt.Printf("%v, %T \n", KB, KB)
	fmt.Printf("%v, %T \n", MB, MB)
	fmt.Printf("%v, %T \n", GB, GB)

	const (
		isAdmin = 1 << iota
		isHq    // shifts the bit 0001 ten times to the left so 1*2^10-1024
		canSeeFinancials
		canSeeAmerica
		canSeeEurope
	)
	var roles byte = isAdmin | canSeeFinancials | isHq
	fmt.Printf("roles %v\n", roles)
	fmt.Printf("is Admin %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("is hq %v\n", isHq&roles == isHq)

}
