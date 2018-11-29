package smap

import "testing"

func TestIntMap(t *testing.T) {
	m := NewIntMap()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")
	m.Set(4, "four")
	m.Set(5, "five")

	m.Delete(1)
	l := len(m.Index())
	if l != 4 {
		t.Errorf("Expected 4 entries, got %d", l)
	}

	m.Delete(2)
	l = len(m.Index())
	if l != 3 {
		t.Errorf("Expected 3 entries, got %d", l)
	}

	if !m.Contains(3) {
		t.Error("Map does not contain 3.")
	}

	n := m.Get(5)
	if n != "five" {
		t.Errorf("Expected 'five' from 5, got %d", n)
	}
}
