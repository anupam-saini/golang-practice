package queue

import (
	"github.com/noti5/practice/deque"
)

type Queue struct {
	deque *deque.Deque
}

func New() *Queue {
	queue := Queue{deque.New()}
	return &queue
}

func (queue *Queue) Push(item interface{}) {
	queue.deque.PushFirst(item)
}

func (queue *Queue) Pop() (interface{}, error) {
	return queue.deque.PopLast()
}

func (queue *Queue) Peek() (interface{}, error) {
	return queue.deque.PeekFirst()	
}

func (queue *Queue) Size() uint32 {
	return queue.deque.Size()		
}