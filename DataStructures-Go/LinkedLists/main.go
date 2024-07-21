package main

import "fmt"

//LEARNING LINKED LISTS (SINGLY)
//We creating a linked list data structure that represent a node

type Node struct {
	data     int
	nextNode *Node
}

//SETTER Methods
func (n *Node) setData(data int) {
	n.data = data
}

func (n *Node) setNextNode(nextNode Node) {
	n.nextNode = &nextNode
}

//GETTER METHODS
func (n Node) getData() int {
	return n.data
}

func main() {
	//local instance of a node
	node := Node{}
	//creatng our first node with int data
	var data int
	var nextNode Node

	fmt.Println("Enter data for node 1: ")
	fmt.Scan(&data)

	node.setData(data)

	//second node
	fmt.Println("Enter data for node 2: ")
	fmt.Scan(&data)

	node.setNextNode(nil)
}
