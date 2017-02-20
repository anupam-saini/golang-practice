package deque

import (
	"github.com/noti5/practice/node"
	"errors"
)

type Value node.Value

type Deque struct {
	Head *node.Node
	Tail *node.Node
	totalSize uint32
}

func New() *Deque {
	deque := Deque{nil, nil, 0}
	return &deque
}

func (deque *Deque) PushFirst(item Value) {
	node := node.Node{Val: item}
	currentHead := deque.Head
	if currentHead == nil {
		deque.Head, deque.Tail = &node, &node
	} else {
		node.Next, currentHead.Prev = currentHead, &node
		deque.Head = &node
	}
	deque.totalSize++
}

func (deque *Deque) PushLast(item Value) {
	node := node.Node{Val: item}
	currentTail := deque.Tail
	if currentTail == nil {
		deque.Head, deque.Tail = &node, &node
	} else {
		node.Prev, currentTail.Next = currentTail, &node
		deque.Tail = &node
	}
	deque.totalSize++
}

func (deque *Deque) PopFirst() (Value, error) {
	var val Value
	var err error
	if deque.Head == nil {
		err = errors.New("No element found in the deque")
	} else {
		node := deque.Head
		newHead := node.Next
		if newHead != nil {
			newHead.Prev = nil
		}
		node.Next = nil
		deque.Head = newHead
		val = node.Val
		deque.totalSize--
	}
	return val, err
}

func (deque *Deque) PopLast() (Value, error) {
	var val Value
	var err error
	if deque.Tail == nil {
		err = errors.New("No element found in the deque")
	} else {
		node := deque.Tail
		newTail := node.Prev
		if newTail != nil {
			newTail.Next = nil
		}
		node.Prev = nil
		deque.Tail = newTail
		val = node.Val
		deque.totalSize--
	}
	return val, err
}

func (deque *Deque) PeekFirst() (Value, error) {
	var val Value
	var err error
	if deque.Head == nil {
		err = errors.New("No element found in the deque")
	} else {
		val = deque.Head.Val
	}
	return val, err
}

func (deque *Deque) PeekLast() (Value, error) {
	var val Value
	var err error
	if deque.Tail == nil {
		err = errors.New("No element found in the deque")
	} else {
		val = deque.Tail.Val
	}
	return val, err
}

func (deque *Deque) Size() uint32 {
	return deque.totalSize
}