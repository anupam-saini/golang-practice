package deque

import (
	"github.com/noti5/practice/value"
	"errors"
)

type Value value.Value

type node struct {
	Next *node
	Prev *node
	Val Value
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

func (deque *Deque) PushFirst(item Value) {
	newNode := node{Val: item}
	currentHead := deque.Head
	if currentHead == nil {
		deque.Head, deque.Tail = &newNode, &newNode
	} else {
		newNode.Next, currentHead.Prev = currentHead, &newNode
		deque.Head = &newNode
	}
	deque.TotalSize++
}

func (deque *Deque) PushLast(item Value) {
	newNode := node{Val: item}
	currentTail := deque.Tail
	if currentTail == nil {
		deque.Head, deque.Tail = &newNode, &newNode
	} else {
		newNode.Prev, currentTail.Next = currentTail, &newNode
		deque.Tail = &newNode
	}
	deque.TotalSize++
}

func (deque *Deque) PopFirst() (Value, error) {
	var val Value
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
		val = oldHead.Val
		deque.TotalSize--
	}
	return val, err
}

func (deque *Deque) PopLast() (Value, error) {
	var val Value
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
		val = oldTail.Val
		deque.TotalSize--
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
	return deque.TotalSize
}