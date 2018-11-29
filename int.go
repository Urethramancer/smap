package smap

import (
	"sort"
)

// SortedIntMap wraps a map with integer keys to keep the keys sorted.
type SortedIntMap struct {
	keys  []int               // Index, sorted as needed
	data  map[int]interface{} // Actual map
	dirty bool                // Set if sorting is needed
}

// NewIntMap returns an initialised integer map.
func NewIntMap() *SortedIntMap {
	s := SortedIntMap{
		keys: make([]int, 0),
		data: make(map[int]interface{}),
	}

	return &s
}

// Set data. Replacing data for an existing key is a cheap operation, as it doesn't force reindexing.
func (s *SortedIntMap) Set(key int, value interface{}) {
	if !s.Contains(key) {
		s.keys = append(s.keys, key)
		s.dirty = true
	}
	s.data[key] = value
}

// Get data.
func (s *SortedIntMap) Get(key int) interface{} {
	return s.data[key]
}

// Delete key. Fast, as sorting is delayed until required.
func (s *SortedIntMap) Delete(key int) {
	if !s.Contains(key) {
		return
	}

	var n, x int
	for n, x = range s.keys {
		if x == key {
			break
		}
	}

	if len(s.data) == 1 {
		s.keys = []int{}
	} else {
		if len(s.keys) > n {
			s.keys[n] = s.keys[len(s.keys)-1]
		}
		s.keys = s.keys[0 : len(s.keys)-1]
	}

	s.dirty = true
	delete(s.data, key)
}

// Contains checks if a key exists.
func (s *SortedIntMap) Contains(key int) bool {
	_, ok := s.data[key]
	return ok
}

// Index returns the sorted list of keys.
func (s *SortedIntMap) Index() []int {
	if s.dirty {
		sort.Ints(s.keys)
	}
	return s.keys
}
