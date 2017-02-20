package queue

import (
	"github.com/noti5/practice/value"
	"github.com/noti5/practice/deque"
)

type Value value.Value

type Queue struct {
	deque *deque.Deque
}

func New() *Queue {
	queue := Queue{deque.New()}
	return &queue
}

func (queue *Queue) Push(item Value) {
	queue.deque.PushFirst(item)
}

func (queue *Queue) Pop() (Value, error) {
	return queue.deque.PopLast()
}

func (queue *Queue) Peek() (Value, error) {
	return queue.deque.PeekFirst()	
}

func (queue *Queue) Size() uint32 {
	return queue.deque.Size()		
}