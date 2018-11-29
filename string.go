// Package smap provides ordered maps for fast lookup and sorted indices.
package smap

import (
	"sort"
)

// SortedStringMap wraps a map with string keys to keep the keys sorted.
type SortedStringMap struct {
	keys  []string               // Index, sorted as needed
	data  map[string]interface{} // Actual map
	dirty bool                   // Set if sorting is needed
}

// NewStringMap returns an initialised string map.
func NewStringMap() *SortedStringMap {
	s := SortedStringMap{
		keys: make([]string, 0),
		data: make(map[string]interface{}),
	}

	return &s
}

// Set data. Replacing data for an existing key is a cheap operation, as it doesn't force reindexing.
func (s *SortedStringMap) Set(key string, value interface{}) {
	if !s.Contains(key) {
		s.keys = append(s.keys, key)
		s.dirty = true
	}
	s.data[key] = value
}

// Get data.
func (s *SortedStringMap) Get(key string) interface{} {
	return s.data[key]
}

// Delete key. Fast, as sorting is delayed until required.
func (s *SortedStringMap) Delete(key string) {
	if !s.Contains(key) {
		return
	}

	var n int
	var x string
	for n, x = range s.keys {
		if x == key {
			break
		}
	}

	if len(s.data) == 1 {
		s.keys = []string{}
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
func (s *SortedStringMap) Contains(key string) bool {
	_, ok := s.data[key]
	return ok
}

// Index returns the sorted list of keys.
func (s *SortedStringMap) Index() []string {
	if s.dirty {
		sort.Strings(s.keys)
	}
	return s.keys
}
