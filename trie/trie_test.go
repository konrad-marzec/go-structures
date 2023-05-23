package trie

import (
	"testing"
)

func TestTrieNew(t *testing.T) {
	tr := New()

	if l := len(tr.node.children); l != 0 {
		t.Errorf("Expected children size %v got %v", 0, l)
	}

	if l := tr.node.value; l != "" {
		t.Errorf("Expected root node value '' got %v", l)
	}

}

func TestTrieInsert(t *testing.T) {
	tr := New()

	tr.Insert("tom")
	tr.Insert("hemburger")
	tr.Insert("hembulger")
	tr.Insert("hello")
	tr.Insert("tomas")
	tr.Insert("hellada")
	tr.Insert("hell")
}

func TestTrieSearch(t *testing.T) {
	tr := New()

	tr.Insert("tom")
	tr.Insert("hello")
	tr.Insert("tomas")
	tr.Insert("hell")

	if l := tr.Search("tom"); l != true {
		t.Errorf("Expected children size %v got %v", 0, l)
	}

	if l := tr.Search("toma"); l != false {
		t.Errorf("Expected children size %v got %v", 0, l)
	}

	if l := tr.Search("tomass"); l != false {
		t.Errorf("Expected children size %v got %v", 0, l)
	}
}

func TestTrieStartsWith(t *testing.T) {
	tr := New()

	tr.Insert("tom")
	tr.Insert("hello")
	tr.Insert("tomas")
	tr.Insert("hell")
	tr.Insert("to")

	if l := tr.StartsWith("tom"); l != true {
		t.Errorf("Expected tom got: %v", l)
	}

	if l := tr.StartsWith("toma"); l != true {
		t.Errorf("Expected toma got: %v", l)
	}

	if l := tr.StartsWith("tomas "); l != false {
		t.Errorf("Not expected tomas : got %v", l)
	}
}

func TestTrieCombo(t *testing.T) {
	tr := New()
	tr.Insert("apple")

	if l := tr.Search("apple"); l != true {
		t.Errorf("Expected apple got %v", l)
	}

	if l := tr.Search("app"); l != false {
		t.Errorf("Not expected app got %v", l)
	}

	if l := tr.StartsWith("app"); l != true {
		t.Errorf("Expected app got %v", l)
	}

	tr.Insert("app")
	if l := tr.Search("app"); l != true {
		t.Errorf("Expected app got %v", l)
	}
}
