package node

type Value interface {}

type Node struct {
	Next *Node
	Prev *Node
	Val Value
}