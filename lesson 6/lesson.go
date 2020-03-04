package main

func main() {
	// if statements
	if true {
		println("ok")
	}

	numbers := map[string]int{
		"one": 1,
		"two": 2,
	}

	if pop, ok := numbers["one"]; ok {
		println(pop)
	}

	// boolean operators
	// == eq
	// != neq
	// >= gt or eeq
	// <= lt or eq
	// < gt
	// > lt

	// or operator
	// uses shortcircuiting so if the first if true i wont check the rest of them
	if true || false {
		println("ok")
	}

	// and operator and not operator
	// uses shortcircuiting so if the first if false i wont check the rest of them
	if true && !false {
		println("ok")
	}

	// if else example
	if true {
		println("ok")
	} else {
		println("not ok")
	}

	// if else if example
	if true {
		println("ok")
	} else if true {
		println("not ok")

	} else {
		println("not ok")
	}

	// switch

}
