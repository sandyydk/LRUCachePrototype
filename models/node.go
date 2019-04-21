package models

// Node represents individual entry or unit of the cache
type Node struct {
	Next       *Node
	PageNumber int
	Previous   *Node
}

func NewNode(pgNum int) *Node {
	var node Node

	node.PageNumber = pgNum
	node.Next = nil
	node.Previous = nil

	return &node
}
