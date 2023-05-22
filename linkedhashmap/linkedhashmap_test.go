package linkedhashmap

import (
	"testing"
)

func TestLRUAddNewItems(t *testing.T) {
	lru := New[string](3)

	lru.Set("first", "first")
	lru.Set("second", "second")
	lru.Set("third", "third")

	if s, _ := lru.Get("first"); s == *new(string) {
		t.Errorf("Expected list size %v got %v", 0, s)
	}

	lru.Set("fourth", "fourth")

	if s, _ := lru.Get("first"); s == *new(string) {
		t.Errorf("Expected list size %v got %v", 0, s)
	}

	if s, _ := lru.Get("fourth"); s == *new(string) {
		t.Errorf("Expected list size %v got %v", 0, s)
	}

	if s, _ := lru.Get("second"); s != *new(string) {
		t.Errorf("Expected list size %v got %v", 0, s)
	}

	lru.Set("fifth", "fifth")

	if s, _ := lru.Get("first"); s == *new(string) {
		t.Errorf("Expected list size %v got %v", 0, s)
	}

	if s, _ := lru.Get("fourth"); s == *new(string) {
		t.Errorf("Expected list size %v got %v", 0, s)
	}

	if s, _ := lru.Get("third"); s != *new(string) {
		t.Errorf("Expected list size %v got %v", 0, s)
	}

}

func TestLRUAddExistingOneManyTimes(t *testing.T) {
	lru := New[string](3)

	lru.Set("first", "first")
	lru.Set("first", "first")
	lru.Set("first", "first")
	lru.Set("second", "second")
	lru.Set("third", "third")
	lru.Set("third", "third")
	lru.Set("third", "third")
	lru.Set("third", "third")
	lru.Set("third", "third")

	if s, _ := lru.Get("first"); s == *new(string) {
		t.Errorf("Expected list size %v got %v", 0, s)
	}

	if s, _ := lru.Get("second"); s == *new(string) {
		t.Errorf("Expected list size %v got %v", 0, s)
	}

	if s, _ := lru.Get("third"); s == *new(string) {
		t.Errorf("Expected list size %v got %v", 0, s)
	}
}
