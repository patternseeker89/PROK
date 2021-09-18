package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type node struct {
	key    string
	name   string
	data   string
	childs []*node
}

type storage struct {
	tree       node
	size       int
	nodesCount int
}

var rootNode node = node{key: "root", name: "Storage", data: "Root node in storage"}

var storage1 []string

func main() {
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

func saveStorageInFile1() {

	file, err := os.Create("data.txt")

	if err != nil {
		//log.Fatal(err)
	}

	defer file.Close()

	for _, value := range storage1 {
		_, err2 := file.WriteString(value + "\n")

		if err2 != nil {
			//log.Fatal(err2)
		}
	}
}

func loadStorageFromFile1() {
	file, err := os.Open("data.txt")
	if err != nil {
		//return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		storage1 = append(storage1, scanner.Text())
	}
}

func showData() {
	for _, value := range storage1 {
		fmt.Println(value)
	}
}

/**
 *	CONSOLE PART
 **/

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
		saveStorageIntoFile()
		os.Exit(0)
	default:
		handleCompoundCommand(command)
	}
}

func handleCompoundCommand(command string) {
	if strings.Contains(command, "insert data") {
		var length = len([]rune(command))
		var lengthSub = len([]rune("insert data"))
		data := command[lengthSub+1 : length]
		storage1 = append(storage1, data)
		fmt.Println("Data saved.")
	} else if strings.Contains(command, "find node") {
		var length = len([]rune(command))
		var lengthSub = len([]rune("find node"))
		key := command[lengthSub+1 : length]
		fmt.Println(key)
		fmt.Println(findNodeInTree(key, &rootNode).key)
	} else if strings.Contains(command, "add node") {
		var length = len([]rune(command))
		var lengthSub = len([]rune("add node"))
		data := command[lengthSub+1 : length]
		fmt.Println(data)
		params := strings.Split(data, " ")
		fmt.Println(params)
		addNode(params[0], params[1], params[2])
	} else if strings.Contains(command, "show storage") {
		print()
	} else {
		fmt.Println("Unknown command!")
	}
}

func help() {
	fmt.Println("Help information of commands:")
	fmt.Println("1. show shorage")
	fmt.Println("2. storage status")
	fmt.Println("3. insert node {{parentKey}} {{name}} {{data}}")
	fmt.Println("3. find node {{parentKey}}")
	fmt.Println("4. delete node {{key}}")
	fmt.Println("5. update node {{key}} {{name}} {{data}}")
	fmt.Println("6. move node {{key}} {{newParentKey}}")
	fmt.Println("7. exit")
}

/**
 *	STORAGE PART
 **/

func findNodeInTree(key string, parentNode *node) *node {
	var neededNode *node = nil

	if parentNode.key != key {
		if rootNode.childs != nil {
			for _, childNode := range parentNode.childs {
				neededNode = findNodeInTree(key, childNode)
				if neededNode != nil {
					break
				}
			}
		} else {

		}
	} else {
		neededNode = parentNode
	}

	return neededNode
}

func addNode(parentKey string, name, data string) {
	var parentNode *node = findNodeInTree(parentKey, &rootNode)
	newNode := node{key: strconv.FormatInt(int64(generateNodekey()), 10), name: name, data: data}
	parentNode.childs = append(parentNode.childs, &newNode)
}

func generateNodekey() int {
	return rand.Intn(10000000)
}

func deleteNode(key string, depth int) {
	//var neededNode = findNodeInTree(key, &rootNode)
}

func updateNode(key string, name, data string) {
	var neededNode = findNodeInTree(key, &rootNode)
	neededNode.name = name
	neededNode.data = data
}

func print() {
	fmt.Println(".")
	fmt.Println("|")
	printTree(&rootNode, "|")
}

func printTree(node *node, separator string) string {
	separator = separator + "----"
	fmt.Println(separator, "#"+node.key+" "+node.name)
	if node.childs != nil {
		for _, childNode := range node.childs {
			if childNode.childs != nil {
				printTree(childNode, separator)
			} else {
				separator = separator + "----"
				fmt.Println(separator, "#"+childNode.key+" "+childNode.name)
				separator = separator[0 : len(separator)-4]
			}
		}
	}

	return separator
}

type data struct {
	str1 int32
	str2 int32
}

func saveStorageIntoFile() {
	file, err := os.Create("storage.dat")
	if err != nil {
		fmt.Println("Couldn't open storage file!")
	}
	defer file.Close()

	var data = serializeStorage()
	fmt.Println(data)

	err = binary.Write(file, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("Save storage file failed!")
	}
}

func loadStorageFromFile() {
	file, err := os.Open("storage.dat")
	if err != nil {
		fmt.Println("Couldn't open storage file!")
	}

	defer file.Close()

	var data1 = data{}

	binary.Read(file, binary.LittleEndian, &data1.str1)
	binary.Read(file, binary.LittleEndian, &data1.str2)
	// binary.Read(file, binary.LittleEndian, &rootNode.data)
	// binary.Read(file, binary.LittleEndian, &rootNode.childs)

	fmt.Println(data1)
}

func serializeStorage() []node {
	return serializeTree(&rootNode)
}

func serializeTree(currentNode *node) []node {
	var serializedData []node

	serializedData = append(serializedData, *currentNode)
	if currentNode.childs != nil {
		for _, childNode := range currentNode.childs {
			serializedData = append(serializedData, *currentNode)
			if childNode.childs != nil {
				serializeTree(childNode)
			}
		}
	}

	// var node = node{key: "root", name: "Storage", data: "Root node in storage"}
	// serializedData = append(serializedData, node)

	return serializedData
}
