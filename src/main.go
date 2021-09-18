package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type serializedNode struct {
	parentKey string
	key       string
	name      string
	data      string
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
	} else if strings.Contains(command, "show node") {
		var length = len([]rune(command))
		var lengthSub = len([]rune("find node"))
		key := command[lengthSub+1 : length]
		showNode(key)
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
	fmt.Println("3. show node {{key}}")
	fmt.Println("4. insert node {{parentKey}} {{name}} {{data}}")
	fmt.Println("5. find node {{parentKey}}")
	fmt.Println("6. delete node {{key}}")
	fmt.Println("7. update node {{key}} {{name}} {{data}}")
	fmt.Println("8. move node {{key}} {{newParentKey}}")
	fmt.Println("9. exit")
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

func showNode(key string) {
	var node = findNodeInTree(key, &rootNode)
	fmt.Println("key: " + node.key)
	fmt.Println("name: " + node.name)
	fmt.Println("data: " + node.data)
	fmt.Print("childs: ")
	fmt.Print(len(node.childs))
	fmt.Println()
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
	fmt.Println("data: ")
	fmt.Println(data)
	// var unData = unserializeStorage(data)
	// fmt.Println(unData)
	// printTree(&unData, "|")

	var map1 = make(map[string]serializedNode)
	mapp := serializedNode{key: "test"}
	map1["root"] = mapp

	var dataBytes []byte
	dataBytes, err = json.Marshal(map1)
	if err != nil {
		fmt.Println("11111!")
	}

	fmt.Println("dataBytes: ")
	fmt.Println(dataBytes)

	err = ioutil.WriteFile("storage.dat", dataBytes, 0777)
	if err != nil {
		fmt.Println("22222!")
	}

	// _, err = file.WriteString(fmt.Sprintln(data))
	// if err != nil {
	// 	fmt.Println("Couldn't open storage file!")
	// }

	// n2, err := file.Write(data)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = binary.Write(file, binary.LittleEndian, data)
	// if err != nil {
	// 	fmt.Println("Save storage file failed!")
	// 	fmt.Println(err)
	// }
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

	// _, err = file.Read(fmt.Sprintln(data))
	// if err != nil {
	// 	fmt.Println("Couldn't open storage file!")
	// }

	fmt.Println(data1)
}

func serializeStorage() []serializedNode {
	return serializeTree(&rootNode, "")
}

func serializeTree(currentNode *node, parentKey string) []serializedNode {
	var serializedData []serializedNode

	fmt.Println("parentKey: " + parentKey)

	var currentSerializedNode = serializedNode{
		parentKey: parentKey,
		key:       currentNode.key,
		name:      currentNode.name,
		data:      currentNode.data}
	serializedData = append(serializedData, currentSerializedNode)

	if currentNode.childs != nil {
		for _, childNode := range currentNode.childs {
			currentSerializedNode := serializedNode{
				parentKey: currentNode.key,
				key:       childNode.key,
				name:      childNode.name,
				data:      childNode.data}
			serializedData = append(serializedData, currentSerializedNode)
			if childNode.childs != nil {
				serializedData = append(serializedData, serializeTree(childNode.childs[0], childNode.key)...)
			}
		}
	}

	return serializedData
}

func unserializeStorage(serializedData []serializedNode) node {
	//var parentKey string
	//var rootNode node
	//var currentNode node
	//var lastNode *node

	nodes := make(map[string]node)

	for _, serializedNode := range serializedData {

		if serializedNode.parentKey == "" {
			nodes[serializedNode.key] = node{key: serializedNode.key, name: serializedNode.name, data: serializedNode.data}
			//parentKey = serializedNode.key
		} else {
			newNode := node{key: serializedNode.key, name: serializedNode.name, data: serializedNode.data}
			parentNode := nodes[serializedNode.parentKey]
			fmt.Println("parentKey: " + serializedNode.parentKey)
			parentNode.childs = append(parentNode.childs, &newNode)
			nodes[serializedNode.parentKey] = parentNode
			nodes[serializedNode.key] = newNode

			fmt.Println("parentKey childs: ")
			fmt.Println(nodes[serializedNode.parentKey].childs)
			//parentKey = serializedNode.key
		}

		// if serializedNode.parentKey == "" {
		// 	//rootNode = node{key: serializedNode.key, name: serializedNode.name, data: serializedNode.data}
		// 	currentNode = node{key: serializedNode.key, name: serializedNode.name, data: serializedNode.data}
		// 	parentKey = serializedNode.key
		// } else {

		// 	newNode := node{key: serializedNode.key, name: serializedNode.name, data: serializedNode.data}
		// 	if serializedNode.parentKey == parentKey {
		// 		currentNode.childs = append(currentNode.childs, &newNode)
		// 		lastNode = &newNode
		// 	} else {
		// 		lastNode.childs = append(lastNode.childs, &newNode)
		// 		parentKey = serializedNode.key
		// 	}

		// }
	}

	fmt.Println("nodes: ")
	fmt.Println(nodes)
	fmt.Println("root: ")
	//fmt.Println(nodes["root"].childs[0].childs[0])

	return nodes["root"]
}
