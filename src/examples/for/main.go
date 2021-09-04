package main

import "fmt"

func main() {
	/**
	*	FOR
	 */

	fmt.Println("-- For --")

	for i := 1; i < 11; i++ {
		fmt.Println(i * i)
	}

	//----------------

	var i = 1
	for i < 3 {
		fmt.Println(i * i)
		i++
	}

	// ------------

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	// ---- for perebor
	var users = [3]string{"Yuriy", "Ron", "Nick"}
	for index, value := range users {
		fmt.Println(index, value)
	}

	// ------ without range
	var users2 = [3]string{"Yuriy", "Ron", "Nick"}
	for i := 0; i < len(users2); i++ {
		fmt.Println(i, users2[i])
	}

	//--- continue

	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	var sum = 0

	for _, value := range numbers {
		if value < 0 {
			continue // переходим к следующей итерации
		}
		sum += value
	}
	fmt.Println("Sum:", sum) // Sum: 27

	//---- break

	var numbers2 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum2 = 0

	for _, value := range numbers2 {
		if value > 4 {
			break // если число больше 4 выходим из цикла
		}
		sum2 += value
	}
	fmt.Println("Sum:", sum2) // Sum: 10

}
