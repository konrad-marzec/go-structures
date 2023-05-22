package linkedlist

import (
	"testing"
)

func TestListNew(t *testing.T) {
	l := New[string]()

	if s := l.size; s != 0 {
		t.Errorf("Expected list size %v got %v", 0, s)
	}

	if h := l.head; h != nil {
		t.Errorf("Expected list head nil got %v", h)
	}

	if h := l.tail; h != nil {
		t.Errorf("Expected list tail nil got %v", h)
	}
}

func TestListInsert(t *testing.T) {
	l := New[string]()

	l.Insert("first")
	l.Insert("second")
	l.Insert("third")

	if s := l.size; s != 3 {
		t.Errorf("Expected list size %v got %v", 3, s)
	}

	if h := l.head; h.value != "first" {
		t.Errorf("Expected list head value to be first got %v", h)
	}

	if h := l.tail; h.value != "third" {
		t.Errorf("Expected list tail value to be third got %v", h)

	}
}

func TestListRemove(t *testing.T) {
	l := New[string]()

	l.Insert("first")
	l.Insert("second")
	l.Insert("third")

	l.Remove("random")
	if s := l.size; s != 3 {
		t.Errorf("Expected list size %v got %v", 3, s)
	}

	l.Remove("second")
	if s := l.size; s != 2 {
		t.Errorf("Expected list size %v got %v", 2, s)
	}

	l.Remove("third")
	if s := l.size; s != 1 {
		t.Errorf("Expected list size %v got %v", 1, s)
	}

	if h := l.head; h.value != "first" {
		t.Errorf("Expected list head value to be first got %v", h)
	}

	if h := l.tail; h.value != "first" {
		t.Errorf("Expected list tail value to be first got %v", h)
	}

	l.Remove("first")
	if s := l.size; s != 0 {
		t.Errorf("Expected list size %v got %v", 0, s)
	}

	if h := l.head; h != nil {
		t.Errorf("Expected list head nil got %v", h)
	}

	if h := l.tail; h != nil {
		t.Errorf("Expected list tail nil got %v", h)
	}
}

func TestListShift(t *testing.T) {
	l := New[string]()

	l.Insert("first")
	l.Insert("second")
	l.Insert("third")

	if f := l.Shift(); f.value != "first" {
		t.Errorf("Expected list shifted value to be first got %v", f.value)
	}

	if s := l.size; s != 2 {
		t.Errorf("Expected list size %v got %v", 2, s)
	}

	if h := l.head; h.value != "second" {
		t.Errorf("Expected list head value to be second got %v", h)
	}

	if f := l.Shift(); f.value != "second" {
		t.Errorf("Expected list shifted value to be second got %v", f.value)
	}

	if f := l.Shift(); f.value != "third" {
		t.Errorf("Expected list shifted value to be third got %v", f.value)
	}

	if f := l.Shift(); f != nil {
		t.Errorf("Expected list shifted element to be nil got %v", f)
	}

	if s := l.size; s != 0 {
		t.Errorf("Expected list size %v got %v", 0, s)
	}

	if h := l.head; h != nil {
		t.Errorf("Expected list head nil got %v", h)
	}

	if h := l.tail; h != nil {
		t.Errorf("Expected list tail nil got %v", h)
	}
}

func TestListPop(t *testing.T) {
	l := New[string]()

	l.Insert("first")
	l.Insert("second")
	l.Insert("third")

	l.Pop()
	if s := l.size; s != 2 {
		t.Errorf("Expected list size %v got %v", 2, s)
	}

	if h := l.tail; h.value != "second" {
		t.Errorf("Expected list tail value to be second got %v", h)
	}

	l.Pop()
	l.Pop()
	l.Pop()
	if s := l.size; s != 0 {
		t.Errorf("Expected list size %v got %v", 0, s)
	}

	if h := l.head; h != nil {
		t.Errorf("Expected list head nil got %v", h)
	}

	if h := l.tail; h != nil {
		t.Errorf("Expected list tail nil got %v", h)
	}
}
