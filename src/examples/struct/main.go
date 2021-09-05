package main

import "fmt"

type bigNode struct {
	key    string
	value  string
	childs []*bigNode
}

type storage struct {
	tree       bigNode
	size       int
	nodesCount int
}

func main() {
	type node struct {
		value int
		next  *node
	}

	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}
	four := node{value: 30}

	first.next = &second
	second.next = &third
	third.next = &four

	var current *node = &first
	for current != nil {
		fmt.Print(current.value, " -> ")
		current = current.next
	}
	fmt.Print("nil ")

	//--------------------

	node1 := bigNode{value: "Text 1"}
	node2 := bigNode{value: "Text 2"}
	node3 := bigNode{value: "Text 3"}
	node4 := bigNode{value: "Text 4"}

	node1.childs = []*bigNode{&node2, &node3}
	node2.childs = []*bigNode{&node4}

	// var currentBigNode *bigNode = &node1
	// for currentBigNode != nil {
	// 	fmt.Print(currentBigNode.value, " -> ")
	// 	//currentBigNode = currentBigNode.childs[1]
	// }
	fmt.Println()
	fmt.Println()
	printNodesTree(&node1)

	fmt.Println(findNodeInTree("Text 4", &node1))
}

func printNodesTree(node *bigNode) {
	fmt.Println(node.value)

	if node.childs != nil {
		for _, childNode := range node.childs {
			printNodesTree(childNode)
		}
		fmt.Println()
	}
}

func findNodeInTree(value string, node *bigNode) *bigNode {
	var neededNode *bigNode

	if node.value != value {
		if node.childs != nil {
			for _, childNode := range node.childs {
				neededNode = findNodeInTree(value, childNode)
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

}

func deleteNode(key string, depth int) {

}

func updateNode(key string, name, data string) {

}
