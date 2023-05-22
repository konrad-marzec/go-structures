package minstack

import "testing"

func TestMainStackNew(t *testing.T) {
	ms := New()

	if l := len(ms.stack); l != 0 {
		t.Errorf("Expected stack size %v got %v", 0, l)
	}

	if e := ms.Top(); e != -1 {
		t.Errorf("Expected %v got %v on stack top", -1, e)
	}

	if e := ms.Pop(); e != -1 {
		t.Errorf("Expected %v got %v", -1, e)
	}
}

func TestMainStackPush(t *testing.T) {
	ms := New()

	ms.Push(4)

	if l := len(ms.stack); l != 1 {
		t.Errorf("Expected stack size %v got %v", 1, l)
	}

	if e := ms.Top(); e != 4 {
		t.Errorf("Expected %v got %v on stack top", 4, e)
	}

	ms.Push(5)
	ms.Push(3)

	if l := len(ms.stack); l != 3 {
		t.Errorf("Expected stack size %v got %v", 3, l)
	}

	if e := ms.Top(); e != 3 {
		t.Errorf("Expected %v got %v on stack top", 3, e)
	}
}

func TestMainStackGetMin(t *testing.T) {
	ms := New()

	ms.Push(4)

	if e := ms.GetMin(); e != 4 {
		t.Errorf("Expected %v got %v as stack min value", 4, e)
	}

	ms.Push(5)

	if e := ms.GetMin(); e != 4 {
		t.Errorf("Expected %v got %v as stack min value", 4, e)
	}

	ms.Push(3)
	ms.Push(3)

	if e := ms.GetMin(); e != 3 {
		t.Errorf("Expected %v got %v as stack min value", 3, e)
	}

	ms.Pop()
	ms.Pop()

	if e := ms.GetMin(); e != 4 {
		t.Errorf("Expected %v got %v as stack min value", 4, e)
	}

	ms.Pop()
	ms.Pop()

	if e := ms.GetMin(); e != -1 {
		t.Errorf("Expected %v got %v as stack min value", -1, e)
	}
}
