package stack

import (
	"github.com/noti5/practice/value"
	"github.com/noti5/practice/deque"
)

type Value value.Value

type Stack struct {
	deque *deque.Deque
}

func New() *Stack {
	stack := Stack{deque.New()}
	return &stack
}

func (stack *Stack) Push(item Value) {
	stack.deque.PushFirst(item)
}

func (stack *Stack) Pop() (Value, error) {
	return stack.deque.PopFirst()
}

func (stack *Stack) Peek() (Value, error) {
	return stack.deque.PeekFirst()	
}

func (stack *Stack) Size() uint32 {
	return stack.deque.Size()		
}