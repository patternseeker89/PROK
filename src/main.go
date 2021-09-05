package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var storage []string

func main() {
	var number int = 5
	var p *int
	p = &number

	fmt.Println(*p)
	*p = 34
	fmt.Println(*p)

	loadStorageFromFile()

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
		fmt.Print("> ")
		command, _ := reader.ReadString('\n')
		command = strings.Replace(command, "\n", "", -1)
		handleCommand(command)
	}
}

func handleCommand(command string) {
	switch command {
	case "help":
		help()
	case "show data":
		showData()
	case "storage status":
		fmt.Println("Records - 12031")
		fmt.Println("Size: - 2.3 Mb")
	case "exit":
		saveStorageInFile()
		os.Exit(0)
	default:
		handleCompoundCommand(command)
	}
}

func handleCompoundCommand(command string) {
	// type consoleCommand struct {
	// 	name string
	// 	id   int
	// }

	if strings.Contains(command, "insert data") {
		var length = len([]rune(command))
		var lengthSub = len([]rune("insert data"))
		data := command[lengthSub+1 : length]
		storage = append(storage, data)
		fmt.Println("Data saved.")
	} else {
		fmt.Println("Unknown command!")
	}

}

func showData() {
	for _, value := range storage {
		fmt.Println(value)
	}
}

func help() {
	fmt.Println("Help information of commands:")
	fmt.Println("1. exit")
	fmt.Println("2. storage status")
	fmt.Println("3. insert data {{data}}")
	fmt.Println("4. show data")
}

func saveStorageInFile() {

	file, err := os.Create("data.txt")

	if err != nil {
		//log.Fatal(err)
	}

	defer file.Close()

	for _, value := range storage {
		_, err2 := file.WriteString(value + "\n")

		if err2 != nil {
			//log.Fatal(err2)
		}
	}
}

func loadStorageFromFile() {
	file, err := os.Open("data.txt")
	if err != nil {
		//return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		storage = append(storage, scanner.Text())
	}
}
