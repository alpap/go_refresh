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

	// slices are like arrays but using copy by reference insted of by value
	a := []int{1, 2, 3}
	b := a
	fmt.Println(a)
	fmt.Println(b)
	b[1] = 4
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%v, length %v \n", a, len(a))
	fmt.Printf("%v, capacity %v \n", a, cap(a))

	//creating slices even if an array is sliced it still uses reference
	ab := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //...
	bb := ab[:]
	cb := ab[3:]
	db := ab[:6]
	eb := ab[3:6] // from the third elemnet up until nut not including the 6th element
	fmt.Println(bb)
	fmt.Println(cb)
	fmt.Println(db)
	fmt.Println(eb)

	eb[1] = 123
	fmt.Println(eb)
	fmt.Println(ab)

	// make
	makeSlice := make([]int, 3)
	fmt.Printf("%v, length %v \n", makeSlice, len(makeSlice))
	fmt.Printf("%v, capacity %v \n", makeSlice, cap(makeSlice))
	// make a slice with a higher capacity

	makeSlice2 := make([]int, 3, 100)
	fmt.Printf("%v, length %v \n", makeSlice2, len(makeSlice2))
	fmt.Printf("%v, capacity %v \n", makeSlice2, cap(makeSlice2))
	// functions
	newSlice := []int{}

	fmt.Printf("%v, length %v \n", newSlice, len(newSlice))
	fmt.Printf("%v, capacity %v \n", newSlice, cap(newSlice))
	newSlice = append(newSlice, 1)

	fmt.Printf("%v, length %v \n", newSlice, len(newSlice))
	fmt.Printf("%v, capacity %v \n", newSlice, cap(newSlice))

	newSlice = append(newSlice, 2, 3, 4, 5)
	fmt.Printf("%v, length %v \n", newSlice, len(newSlice))
	fmt.Printf("%v, capacity %v \n", newSlice, cap(newSlice))

	// watch out cause the capacity is doubled here this can be memory inefficient
	newSlice = append(newSlice, 2, 3, 4, 5)
	fmt.Printf("%v, length %v \n", newSlice, len(newSlice))
	fmt.Printf("%v, capacity %v \n", newSlice, cap(newSlice))
	// concatinating slices using th spread operator
	newSlice = append(newSlice, []int{2, 3, 4, 5}...)
	fmt.Printf("%v, length %v \n", newSlice, len(newSlice))
	fmt.Printf("%v, capacity %v \n", newSlice, cap(newSlice))

	// pop elements from the stack
	// remove the first element from the slice
	newCutSlice := newSlice[1:]
	fmt.Printf("%v, length %v \n", newCutSlice, len(newCutSlice))
	fmt.Printf("%v, capacity %v \n", newCutSlice, cap(newCutSlice))
	// remove the last element from the slice
	newCutSlice2 := newSlice[:len(newSlice)-1]
	fmt.Printf("%v, length %v \n", newCutSlice2, len(newCutSlice2))
	fmt.Printf("%v, capacity %v \n", newCutSlice2, cap(newCutSlice2))
	// remove a middle element from the slice
	newCutSlice3 := append(newSlice[2:4], newSlice[7:10]...)
	fmt.Printf("%v, length %v \n", newCutSlice3, len(newCutSlice3))
	fmt.Printf("%v, capacity %v \n", newCutSlice3, cap(newCutSlice3))
	// watchout the original slice is geeting affected by this
	fmt.Printf("%v, length %v \n", newCutSlice, len(newCutSlice))
	fmt.Printf("%v, capacity %v \n", newCutSlice, cap(newCutSlice))
}
