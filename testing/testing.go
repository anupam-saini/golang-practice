package main

import (
	"fmt"
	"github.com/noti5/practice/stack"
	"github.com/noti5/practice/queue"
)

func main() {

	/* stack */
	stack := stack.New()

	stack.Push(11)
	fmt.Println("stack.Push(11)")
	val, _ := stack.Peek()
	fmt.Println("stack.Peek()", val)

	stack.Push(22)
	fmt.Println("stack.Push(22)")
	val, _ = stack.Peek()
	fmt.Println("stack.Peek()", val)

	val, _ = stack.Pop()
	fmt.Println("stack.Pop()", val)

	val, _ = stack.Pop()
	fmt.Println("stack.Pop()", val)
	fmt.Println("stack.Size()", stack.Size())

	/* queue */
	queue := queue.New()

	queue.Push(11)
	fmt.Println("queue.Push(11)")
	val, _ := queue.Peek()
	fmt.Println("queue.Peek()", val)

	queue.Push(22)
	fmt.Println("queue.Push(22)")
	val, _ = queue.Peek()
	fmt.Println("queue.Peek()", val)

	val, _ = queue.Pop()
	fmt.Println("queue.Pop()", val)

	val, err := queue.Pop()
	fmt.Println("queue.Pop()", val, err)
	fmt.Println("queue.Size()", queue.Size())
}