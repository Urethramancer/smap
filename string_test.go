package smap

import (
	"testing"
)

func TestStringMap(t *testing.T) {
	m := NewStringMap()
	m.Set("one", 1)
	m.Set("two", 2)
	m.Set("three", 3)
	m.Set("four", 4)
	m.Set("five", 5)

	m.Delete("one")
	l := len(m.Index())
	if l != 4 {
		t.Errorf("Expected 4 entries, got %d", l)
	}

	m.Delete("two")
	l = len(m.Index())
	if l != 3 {
		t.Errorf("Expected 3 entries, got %d", l)
	}

	if !m.Contains("three") {
		t.Error("Map does not contain 'three'.")
	}

	n := m.Get("five")
	if n != 5 {
		t.Errorf("Expected 5 from 'five', got %d", n)
	}
}
