package main

import (
	"fmt"
)

func main() {
	// arrays
	// you have to specify at creating what type of data you will store
	grades := [3]int{97, 85, 93}
	fmt.Printf("%v, %T \n", grades, grades)
	// we dont need to say how big the array is if we are initializing it we can just say  ... which means expand to the required size
	grades2 := [...]int{97, 85, 93}
	fmt.Printf("%v, %T \n", grades2, grades2)
	// if we dont initialize hte array it will be automatically  intialized with 0s
	var grades3 [3]int
	fmt.Printf("%v, %T \n", grades3, grades3)
	grades3[0] = 12
	fmt.Printf("%v, %T \n", grades3, grades3)
	fmt.Printf("%v, %T \n", len(grades3), grades3)

	// array of array
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Printf("%v, %T \n", identityMatrix, identityMatrix)

	identityMatrix2 := identityMatrix // golang copies by values not by reference
	identityMatrix2[1][1] = 5
	fmt.Println(identityMatrix)
	fmt.Println(identityMatrix2)

	identityMatrix3 := &identityMatrix // here we pass the reference
	identityMatrix3[1][1] = 2

	fmt.Println(identityMatrix)
	fmt.Println(identityMatrix2)

	// slices
	a := []int{1, 2, 3}
	fmt.Printf("%v, length %v \n", a, len(a))
	fmt.Printf("%v, capacity %v \n", a, cap(a))
}
