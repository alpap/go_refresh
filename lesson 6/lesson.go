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
	state := 2
	switch state {
	case 1:
		println("only one")
	case 2, 3, 4:
		println("more than one")
	default:
		println("multiple")
	}

	// make a modification and then pass it to the case
	switch i := state + 3; i {
	case 1:
		println("only one")
	case 2, 3, 4:
		println("more than one")
	default:
		println("multiple")
	}

	// tagless syntax
	newState := 22
	switch {
	case newState < 1: // only this get evaluated it warks as an or
		println("only one")
		fallthrough // no matter what it will go and execute the next case
	case newState <= 2:
		println("more than one")
	default:
		println("multiple")
	}

	// type switch
	var tp interface{} = 1
	switch tp.(type) {
	case int:
		println("int")
	case float64:
		println("float")
	case string:
		println("string")
	case [3]int:
		println("array 3")
	case [2]int: // different type from the above
		break // break out of the switch
		println("array 2")
	}
}
