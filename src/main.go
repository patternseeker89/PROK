package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
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
	ParentKey string
	Key       string
	Name      string
	Data      string
}

var rootNode node = node{key: "root", name: "Storage", data: "Root node in storage"}

type colors struct {
	red    string
	green  string
	yellow string
	blue   string
	reset  string
}

var consoleColors colors = colors{red: "\033[31m", green: "\033[32m", yellow: "\033[33m", blue: "\033[34m", reset: "\033[0m"}

type consoleCommands struct {
	insertData    string
	findNode      string
	addNode       string
	showNode      string
	showStorage   string
	help          string
	storageStatus string
	exit          string
}

var consoleCommandList consoleCommands = consoleCommands{
	findNode:      "find node",
	addNode:       "add node",
	showNode:      "show node",
	showStorage:   "show storage",
	help:          "help",
	storageStatus: "storage status",
	exit:          "exit",
}

func main() {
	loadStorageFromFile()
	introInformation()
	console()

	// colorPurple := "\033[35m"
	// colorCyan := "\033[36m"
	// colorWhite := "\033[37m"
	// fmt.Println(string(colorWhite), "test")
	// fmt.Println(string(colorCyan), "test", string(colorReset))

}

/**
 *	CONSOLE PART
 **/

func introInformation() {
	fmt.Println(string(consoleColors.blue), "PROK 2021 email: porfirovskiy@gmail.com", string(consoleColors.reset))
	fmt.Println()
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
	case consoleCommandList.help:
		help()
	case consoleCommandList.storageStatus:
		fmt.Println("Records - 12031")
		fmt.Println("Size: - 2.3 Mb")
	case consoleCommandList.exit:
		saveStorageIntoFile()
		os.Exit(0)
	default:
		handleCommandWithParams(command)
	}
}

func handleCommandWithParams(command string) {
	if strings.Contains(command, consoleCommandList.findNode) {
		handleConsoleCommandFindNode(command)
	} else if strings.Contains(command, consoleCommandList.addNode) {
		handleConsoleCommandAddNode(command)
	} else if strings.Contains(command, consoleCommandList.showNode) {
		handleConsoleCommandShowNode(command)
	} else if strings.Contains(command, consoleCommandList.showStorage) {
		print()
	} else {
		fmt.Println("Unknown command!")
	}
}

func handleConsoleCommandFindNode(command string) {
	params := getParamsFromCommand(command, consoleCommandList.findNode)

	if len(params) > 0 {
		key := params[0]
		node := findNodeInTree(key, &rootNode)
		if node != nil {
			fmt.Println()
			showSuccessMessage("Node found:")
			fmt.Println("key: ", node.key)
			fmt.Println("name: ", node.name)
			fmt.Println()
		} else {
			showErrorMessage("Node not found!")
		}
	} else {
		showErrorMessage("Empty node key!")
	}
}

func handleConsoleCommandAddNode(command string) {
	params := getParamsFromCommand(command, consoleCommandList.addNode)

	if len(params) > 2 {
		addNode(params[0], params[1], params[2])
		showSuccessMessage("Node added")
	} else {
		fmt.Println(string(consoleColors.red), "Wrong node params!", string(consoleColors.reset))
	}
}

func handleConsoleCommandShowNode(command string) {
	params := getParamsFromCommand(command, consoleCommandList.showNode)

	if len(params) > 0 {
		key := params[0]
		showNode(key)
	} else {
		fmt.Println(string(consoleColors.red), "Empty node key!", string(consoleColors.reset))
	}
}

func getParamsFromCommand(command string, commandName string) []string {
	var length = len([]rune(command))
	var lengthSub = len([]rune(commandName))

	var params []string
	if length > lengthSub {
		data := command[lengthSub+1 : length]
		params = strings.Split(data, " ")
	}

	return params
}

func showErrorMessage(message string) {
	fmt.Println(string(consoleColors.red), message, string(consoleColors.reset))
}

func showSuccessMessage(message string) {
	fmt.Println(string(consoleColors.green), message, string(consoleColors.reset))
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
	if node != nil {
		fmt.Println()
		fmt.Println("key: " + node.key)
		fmt.Println("name: " + node.name)
		fmt.Println("data: " + node.data)
		fmt.Print("childs: ")
		fmt.Print(len(node.childs))
		fmt.Println()
		fmt.Println()
	} else {
		fmt.Println(string(consoleColors.red), "Node does not exist for this key!", string(consoleColors.reset))
	}
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
	fmt.Println(data[0])

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(jsonData)

	result, err := file.Write(jsonData)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(result)
}

func loadStorageFromFile() {
	file, err := os.Open("storage.dat")
	if err != nil {
		fmt.Println("Couldn't open storage file!")
	}

	defer file.Close()

	data := make([]byte, 64)
	for {
		numberOfBytes, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print("data: ", string(data[:numberOfBytes]))
	}

	var serializedData serializedNode
	error := json.Unmarshal(data, &serializedData)
	if error != nil {
		fmt.Println("error:", error)
	}
	fmt.Println(serializedData)
}

func serializeStorage() []serializedNode {
	return serializeTree(&rootNode, "")
}

func serializeTree(currentNode *node, parentKey string) []serializedNode {
	var serializedData []serializedNode

	fmt.Println("parentKey: " + parentKey)

	var currentSerializedNode = serializedNode{
		ParentKey: parentKey,
		Key:       currentNode.key,
		Name:      currentNode.name,
		Data:      currentNode.data}
	serializedData = append(serializedData, currentSerializedNode)

	if currentNode.childs != nil {
		for _, childNode := range currentNode.childs {
			currentSerializedNode := serializedNode{
				ParentKey: currentNode.key,
				Key:       childNode.key,
				Name:      childNode.name,
				Data:      childNode.data}
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

		if serializedNode.ParentKey == "" {
			nodes[serializedNode.Key] = node{key: serializedNode.Key, name: serializedNode.Name, data: serializedNode.Data}
			//parentKey = serializedNode.key
		} else {
			newNode := node{key: serializedNode.Key, name: serializedNode.Name, data: serializedNode.Data}
			parentNode := nodes[serializedNode.ParentKey]
			fmt.Println("parentKey: " + serializedNode.ParentKey)
			parentNode.childs = append(parentNode.childs, &newNode)
			nodes[serializedNode.ParentKey] = parentNode
			nodes[serializedNode.Key] = newNode

			fmt.Println("parentKey childs: ")
			fmt.Println(nodes[serializedNode.ParentKey].childs)
			//parentKey = serializedNode.key
		}
	}

	fmt.Println("nodes: ")
	fmt.Println(nodes)
	fmt.Println("root: ")
	//fmt.Println(nodes["root"].childs[0].childs[0])

	return nodes["root"]
}
