package algo

import "fmt"

type Node struct {
	NextNode *Node
	Proprety int
}

type LinkedList struct {
	HeadNode *Node
}

func (ll *LinkedList) AddToHead(p int) {
	var node = &Node{}
	node.Proprety = p
	node.NextNode = nil

	if ll.HeadNode != nil {
		node.NextNode = ll.HeadNode
	}
	ll.HeadNode = node
}

func (ll *LinkedList) IterateList() {
	var node *Node
	for node = ll.HeadNode; node != nil; node = node.NextNode {
		fmt.Println(node.Proprety)
	}
}

func (ll *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node

	for node = ll.HeadNode; node != nil; node = node.NextNode {
		if node.NextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (ll *LinkedList) AddToEnd(p int) {
	var node = Node{}
	node.Proprety = p
	node.NextNode = nil
	var lastNode *Node
	lastNode = ll.LastNode()
	if lastNode != nil {
		lastNode.NextNode = &node
	}
}

func (ll *LinkedList) NodeWithValue(p int) *Node {
	var node *Node
	var nodeWith *Node
	for node = ll.HeadNode; node != nil; node = node.NextNode {
		if node.Proprety == p {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

func (ll *LinkedList) AddAfter(np int, p int) {
	var node = &Node{}
	node.Proprety = p
	node.NextNode = nil
	var nodeWith *Node
	nodeWith = ll.NodeWithValue(np)
	if nodeWith != nil {
		node.NextNode = nodeWith.NextNode
		nodeWith.NextNode = node
	}
}
