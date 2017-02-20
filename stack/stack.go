package stack

import (
	"github.com/noti5/practice/node"
	"github.com/noti5/practice/deque"
)

type Stack struct {
	deque *deque.Deque
}

func New() *Stack {
	stack := Stack{deque.New()}
	return &stack
}

func (stack *Stack) Push(item node.Value) {
	stack.deque.PushFirst(item)
}

func (stack *Stack) Pop() (node.Value, error) {
	return stack.deque.PopFirst()
}

func (stack *Stack) Peek() (node.Value, error) {
	return stack.deque.PeekFirst()	
}

func (stack *Stack) Size() uint32 {
	return stack.deque.Size()		
}