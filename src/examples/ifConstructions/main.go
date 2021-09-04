package main

import "fmt"

func main() {
	/**
	* -- IF-chiki --
	 */

	fmt.Println("-- IF-chiki --")

	a := 10
	b := 10

	if a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a === b")
	}

	// -- switch --

	aa := 10

	switch aa {
	case 9:
		fmt.Println("9")
	case 10:
		fmt.Println("10")
	default:
		fmt.Println("def")
	}
}
