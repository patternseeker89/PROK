package main

import (
	"fmt"
)

var storage []string

func main() {
	/**
	* Variables
	 */
	fmt.Println("-- Variables --")
	var text string
	text = "Hello world!"
	fmt.Println(text)

	// ------------------------

	var text2 string = "Hello world! 2"
	fmt.Println(text2)

	// ------------------------

	var (
		message string = "go to"
		age     int    = 32
	)
	fmt.Println(message)
	fmt.Println(age)

	// --------------------

	amount := 10
	fmt.Println(amount)

	var name1, age1 = "Tom", 27

	fmt.Println(name1)
	fmt.Println(age1)
}
