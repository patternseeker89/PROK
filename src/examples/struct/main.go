package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type bigNode struct {
	key    string
	name   string
	data   string
	childs []*bigNode
}

type storage struct {
	tree       bigNode
	size       int
	nodesCount int
}

var node bigNode = bigNode{
	key:  "jj38",
	name: "root",
	data: "Using Join() function: This function concatenates all the elements present in the slice of string into a single string. This function is available in string package.",
}

func main() {
	fmt.Println(findNodeInTree("jj38", &node).key)

	addNode("jj38", "Banks", "List of all my banks")
	addNode("8498081", "Privatbank", "very big")
	addNode("8498081", "PUMB", "Sister work")
	addNode("9727887", "Deposits", "List of my deposits")
	addNode("9727887", "Accounts", "List of my accounts")
	addNode("1902081", "FOP cabinet", "---")

	fmt.Println("----------------------------")
	fmt.Println()
	print()
}

func findNodeInTree(key string, node *bigNode) *bigNode {
	var neededNode *bigNode

	if node.key != key {
		if node.childs != nil {
			for _, childNode := range node.childs {
				neededNode = findNodeInTree(key, childNode)
				if neededNode != nil {
					break
				}
			}
		} else {

		}
	} else {
		neededNode = node
	}

	return neededNode
}

func addNode(parentKey string, name, data string) {
	var parentNode *bigNode = findNodeInTree(parentKey, &node)
	newNode := bigNode{key: strconv.FormatInt(int64(generateNodekey()), 10), name: name, data: data}
	parentNode.childs = append(parentNode.childs, &newNode)
}

func generateNodekey() int {
	return rand.Intn(10000000)
}

func deleteNode(key string, depth int) {

}

func updateNode(key string, name, data string) {

}

func print() {
	fmt.Println(".")
	fmt.Println("|")
	printTree(&node, "|")
}

func printTree(node *bigNode, separator string) string {
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
