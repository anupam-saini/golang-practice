package stack

import (
	"fmt"
	"testing"
)

func TestStack_Push(t *testing.T) {
	var s *Stack = New()
	s.Push(2)
	element, err := s.Pop()
	if err != nil {
		t.Errorf("Error when popping an element from stack error=%v", err)
	}
	if element != 2 {
		t.Error(fmt.Sprintf("Expecting value %d, got %d", 2, element))
	}
}

func TestStack_Pop_EmptyStack(t *testing.T) {
	var s *Stack = New()
	_, err := s.Pop()
	if err == nil {
		t.Errorf("Expecting error, but pop passed with no elements in stack")
	}
}

func TestStack_Peek(t *testing.T) {
	var s *Stack = New()
	s.Push(2);
	element, _ := s.Peek();
	if element!= 2 {
		t.Errorf("Expecting peek value %v, got %v", 2, element)
	}
	size := s.Size()
	if size != 1 {
		t.Errorf("Expecting size of stack %v, got %v", 1, size)
	}
}

func TestStack_Peek_EmptyStack(t *testing.T) {
	var s *Stack = New()
	element, err := s.Peek()
	if err == nil {
		t.Error("Expecting error when peek is called on empty stack, received element %v", element)
	}
}

func TestStack_Size(t *testing.T) {
	var s *Stack = New()
	size := s.Size()
	if size!=0 {
		t.Error("Expecting stack to be emtpy when initialized")
	}

	s.Push(2)
	size = s.Size()
	if size != 1 {
		t.Error("Expecting stack size to be %v, got %v", 1, size)
	}
}