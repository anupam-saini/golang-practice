package deque

import (
	"errors"
	"sync"
)

type node struct {
	next *node
	prev *node
	value interface{}
}

type Deque struct {
	head *node
	tail *node
	totalSize uint32
	mux sync.Mutex
}

func New() *Deque {
	deque := new(Deque)
	return deque
}

func (deque *Deque) PushFirst(item interface{}) {
	deque.mux.Lock()
	newNode := node{value: item}
	currentHead := deque.head
	if currentHead == nil {
		deque.head, deque.tail = &newNode, &newNode
	} else {
		newNode.next, currentHead.prev = currentHead, &newNode
		deque.head = &newNode
	}
	deque.totalSize++
	deque.mux.Unlock()
}

func (deque *Deque) PushLast(item interface{}) {
	deque.mux.Lock()
	newNode := node{value: item}
	currentTail := deque.tail
	if currentTail == nil {
		deque.head, deque.tail = &newNode, &newNode
	} else {
		newNode.prev, currentTail.next = currentTail, &newNode
		deque.tail = &newNode
	}
	deque.totalSize++
	deque.mux.Unlock()
}

func (deque *Deque) PopFirst() (interface{}, error) {
	deque.mux.Lock()
	var value interface{}
	var err error
	if deque.head == nil {
		err = errors.New("No element found in the deque")
	} else {
		oldHead := deque.head
		newHead := oldHead.next
		if newHead != nil {
			newHead.prev = nil
		}
		oldHead.next = nil
		deque.head = newHead
		value = oldHead.value
		deque.totalSize--
	}
	defer deque.mux.Unlock()
	return value, err
}

func (deque *Deque) PopLast() (interface{}, error) {
	deque.mux.Lock()
	var value interface{}
	var err error
	if deque.tail == nil {
		err = errors.New("No element found in the deque")
	} else {
		oldTail := deque.tail
		newTail := oldTail.prev
		if newTail != nil {
			newTail.next = nil
		}
		oldTail.prev = nil
		deque.tail = newTail
		value = oldTail.value
		deque.totalSize--
	}
	defer deque.mux.Unlock()
	return value, err
}

func (deque *Deque) PeekFirst() (interface{}, error) {
	deque.mux.Lock()
	var value interface{}
	var err error
	if deque.head == nil {
		err = errors.New("No element found in the deque")
	} else {
		value = deque.head.value
	}
	defer deque.mux.Unlock()
	return value, err
}

func (deque *Deque) PeekLast() (interface{}, error) {
	deque.mux.Lock()
	var value interface{}
	var err error
	if deque.tail == nil {
		err = errors.New("No element found in the deque")
	} else {
		value = deque.tail.value
	}
	defer deque.mux.Unlock()
	return value, err
}

func (deque *Deque) Size() uint32 {
	deque.mux.Lock()
	defer deque.mux.Unlock()
	return deque.totalSize
}