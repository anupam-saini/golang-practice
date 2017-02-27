package deque

import (
	"errors"
)

type node struct {
	Next *node
	Prev *node
	Value interface{}
}

type Deque struct {
	Head *node
	Tail *node
	TotalSize uint32
}

func New() *Deque {
	deque := new(Deque)
	return deque
}

func (deque *Deque) PushFirst(item interface{}) {
	newNode := node{Value: item}
	currentHead := deque.Head
	if currentHead == nil {
		deque.Head, deque.Tail = &newNode, &newNode
	} else {
		newNode.Next, currentHead.Prev = currentHead, &newNode
		deque.Head = &newNode
	}
	deque.TotalSize++
}

func (deque *Deque) PushLast(item interface{}) {
	newNode := node{Value: item}
	currentTail := deque.Tail
	if currentTail == nil {
		deque.Head, deque.Tail = &newNode, &newNode
	} else {
		newNode.Prev, currentTail.Next = currentTail, &newNode
		deque.Tail = &newNode
	}
	deque.TotalSize++
}

func (deque *Deque) PopFirst() (interface{}, error) {
	var value interface{}
	var err error
	if deque.Head == nil {
		err = errors.New("No element found in the deque")
	} else {
		oldHead := deque.Head
		newHead := oldHead.Next
		if newHead != nil {
			newHead.Prev = nil
		}
		oldHead.Next = nil
		deque.Head = newHead
		value = oldHead.Value
		deque.TotalSize--
	}
	return value, err
}

func (deque *Deque) PopLast() (interface{}, error) {
	var value interface{}
	var err error
	if deque.Tail == nil {
		err = errors.New("No element found in the deque")
	} else {
		oldTail := deque.Tail
		newTail := oldTail.Prev
		if newTail != nil {
			newTail.Next = nil
		}
		oldTail.Prev = nil
		deque.Tail = newTail
		value = oldTail.Value
		deque.TotalSize--
	}
	return value, err
}

func (deque *Deque) PeekFirst() (interface{}, error) {
	var value interface{}
	var err error
	if deque.Head == nil {
		err = errors.New("No element found in the deque")
	} else {
		value = deque.Head.Value
	}
	return value, err
}

func (deque *Deque) PeekLast() (interface{}, error) {
	var value interface{}
	var err error
	if deque.Tail == nil {
		err = errors.New("No element found in the deque")
	} else {
		value = deque.Tail.Value
	}
	return value, err
}

func (deque *Deque) Size() uint32 {
	return deque.TotalSize
}