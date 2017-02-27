package stack

import (
	"github.com/noti5/practice/deque"
)

type Stack struct {
	deque *deque.Deque
}

func New() *Stack {
	stack := Stack{deque.New()}
	return &stack
}

func (stack *Stack) Push(item interface{}) {
	stack.deque.PushFirst(item)
}

func (stack *Stack) Pop() (interface{}, error) {
	return stack.deque.PopFirst()
}

func (stack *Stack) Peek() (interface{}, error) {
	return stack.deque.PeekFirst()	
}

func (stack *Stack) Size() uint32 {
	return stack.deque.Size()		
}