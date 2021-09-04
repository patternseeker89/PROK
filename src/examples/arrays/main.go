package main

import "fmt"

func main() {
	/**
	 * Arrays
	 */
	fmt.Println("-- Arrays --")

	var numbers [5]int
	fmt.Println(numbers)

	var numbers2 [5]int = [5]int{3, 24, 3454, 24, 55}
	fmt.Println(numbers2)

	fmt.Println(numbers2[3])
	numbers2[3] = 33445
	fmt.Println(numbers2[3])

	numbers3 := [5]int{1, 2}
	fmt.Println(numbers3)
}
