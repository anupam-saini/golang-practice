package deque

import "testing"

func TestNew(t *testing.T) {
	deque := New()
	if deque == nil {
		t.Error("deque.New() == nil")
	}
}


func TestDeque_PushFirst_PopFirst(t *testing.T) {
	deque := New()
	val := 2
	deque.PushFirst(val);
	result, err := deque.PopFirst()
	if err != nil {
		t.Errorf("Error when popping element from the deque. Error: %v", err)
		return
	}
	if result != val {
		t.Errorf("Expected %v, found %v", val, result)
	}

	size := deque.Size()
	if size != 0 {
		t.Errorf("Error in size. Expected 0 found %v", size)
	}
}

func TestDeque_PushLast_PopLast(t *testing.T) {
	deque := New()
	val := 2
	deque.PushLast(val);
	result, err := deque.PopLast()
	if err != nil {
		t.Errorf("Error when popping element from the deque. Error: %v", err)
		return
	}
	if result != val {
		t.Errorf("Expected %v, found %v", val, result)
	}

	size := deque.Size()
	if size != 0 {
		t.Errorf("Error in size. Expected 0 found %v", size)
	}
}

func TestDeque_PushFirst_PopLast(t *testing.T) {
	deque := New()
	val := 2
	deque.PushFirst(val);
	result, err := deque.PopLast()
	if err != nil {
		t.Errorf("Error when popping element from the deque. Error: %v", err)
		return
	}
	if result != val {
		t.Errorf("Expected %v, found %v", val, result)
	}

	size := deque.Size()
	if size != 0 {
		t.Errorf("Error in size. Expected 0 found %v", size)
	}
}

func TestDeque_PushFirst_PopLast2(t *testing.T) {
	deque := New()
	val1 := 4
	val2 := 5
	deque.PushFirst(val1);
	deque.PushFirst(val2);
	result, err := deque.PopLast()
	if err != nil {
		t.Errorf("Error when popping element from the deque. Error: %v", err)
		return
	}
	if result != val1 {
		t.Errorf("Expected %v, found %v", val1, result)
	}

	result, err = deque.PopLast()
	if err != nil {
		t.Errorf("Error when popping element from the deque. Error: %v", err)
		return
	}
	if result != val2 {
		t.Errorf("Expected %v, found %v", val2, result)
	}

	size := deque.Size()
	if size != 0 {
		t.Errorf("Error in size. Expected 0 found %v", size)
	}
}

func TestDeque_PushLast_PopFirst(t *testing.T) {
	deque := New()
	val := 2
	deque.PushLast(val);
	result, err := deque.PopFirst()
	if err != nil {
		t.Errorf("Error when popping element from the deque. Error: %v", err)
		return
	}
	if result != val {
		t.Errorf("Expected %v, found %v", val, result)
	}

	size := deque.Size()
	if size != 0 {
		t.Errorf("Error in size. Expected 0 found %v", size)
	}
}

func TestDeque_PushFirst_Peek(t *testing.T) {
	deque := New()
	val1 := 4
	val2 := 5
	deque.PushFirst(val1);
	deque.PushFirst(val2);
	result, err := deque.PeekFirst()
	if err != nil {
		t.Errorf("Error when popping element from the deque. Error: %v", err)
		return
	}
	if result != val2 {
		t.Errorf("Expected %v, found %v", val2, result)
	}

	result, err = deque.PeekLast()
	if err != nil {
		t.Errorf("Error when popping element from the deque. Error: %v", err)
		return
	}
	if result != val1 {
		t.Errorf("Expected %v, found %v", val1, result)
	}

	size := deque.Size()
	if size != 2 {
		t.Errorf("Error in size. Expected 0 found %v", size)
	}
}

func TestDeque_PushLast_Peek(t *testing.T) {
	deque := New()
	val1 := 4
	val2 := 5
	deque.PushLast(val1);
	deque.PushLast(val2);
	result, err := deque.PeekFirst()
	if err != nil {
		t.Errorf("Error when popping element from the deque. Error: %v", err)
		return
	}
	if result != val1 {
		t.Errorf("Expected %v, found %v", val1, result)
	}

	result, err = deque.PeekLast()
	if err != nil {
		t.Errorf("Error when popping element from the deque. Error: %v", err)
		return
	}
	if result != val2 {
		t.Errorf("Expected %v, found %v", val2, result)
	}

	size := deque.Size()
	if size != 2 {
		t.Errorf("Error in size. Expected 0 found %v", size)
	}
}


func TestDeque_Size(t *testing.T) {
	deque := New()
	limit := uint32(10)
	for i := uint32(0); i < limit; i++ {
		deque.PushFirst(i)
	}
	size := deque.Size()
	if size != limit {
		t.Errorf("Error in size. Expected 0 found %v", size)
	}
}

