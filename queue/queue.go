package queue

import (
	"github.com/noti5/practice/node"
	"github.com/noti5/practice/deque"
)

type Queue struct {
	deque *deque.Deque
}

func New() *Queue {
	queue := Queue{deque.New()}
	return &queue
}

func (queue *Queue) Push(item node.Value) {
	queue.deque.PushFirst(item)
}

func (queue *Queue) Pop() (node.Value, error) {
	return queue.deque.PopLast()
}

func (queue *Queue) Peek() (node.Value, error) {
	return queue.deque.PeekFirst()	
}

func (queue *Queue) Size() uint32 {
	return queue.deque.Size()		
}