package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	introInformation()
	console()

	// fmt.Println("Hello, world!")

	// colorReset := "\033[0m"

	// colorRed := "\033[31m"
	// colorGreen := "\033[32m"
	// colorYellow := "\033[33m"
	// colorBlue := "\033[34m"
	// colorPurple := "\033[35m"
	// colorCyan := "\033[36m"
	// colorWhite := "\033[37m"

	// fmt.Println(string(colorRed), "test")
	// fmt.Println(string(colorGreen), "test")
	// fmt.Println(string(colorYellow), "test")
	// fmt.Println(string(colorBlue), "test")
	// fmt.Println(string(colorPurple), "test")
	// fmt.Println(string(colorWhite), "test")
	// fmt.Println(string(colorCyan), "test", string(colorReset))

}

func introInformation() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "PROK 2021 email: porfirovskiy@gmail.com", string(colorReset))
	fmt.Println("---------------------")
}

func console() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		operations(text)
	}
}

func operations(text string) {
	if strings.Compare("hi", text) == 0 {
		fmt.Println("hello, Yourself")
	}

	if strings.Compare("exit", text) == 0 {
		os.Exit(0)
	}

	if strings.Compare("show data", text) == 0 {
		fmt.Println("1. Yuriy")
		fmt.Println("2. Koly")
		fmt.Println("1. Ivan")
	}
}
