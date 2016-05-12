// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package set

import (
	"sort"
	"sync"
)

// HashSet is a simple Set structure.
type HashSet struct {
	m   map[string]struct{}
	mux sync.RWMutex
}

// NewHashSet creates and returns a new HashSet
// items is a list of initial set members.
func NewHashSet(items []string) *HashSet {
	h := &HashSet{m: make(map[string]struct{})}
	if items != nil {
		for _, x := range items {
			h.m[x] = struct{}{}
		}
	}
	return h
}

// Add adds an item to the HashSet.
func (h *HashSet) Add(item string) {
	h.mux.Lock()
	h.m[item] = struct{}{}
	h.mux.Unlock()
}

// Del removes item from the HashSet.
func (h *HashSet) Del(item string) {
	h.mux.Lock()
	delete(h.m, item)
	h.mux.Unlock()
}

// Clear empties the HashSet.
func (h *HashSet) Clear() {
	h.mux.Lock()
	h.m = make(map[string]struct{})
	h.mux.Unlock()
}

// Items returns the items as a slice
func (h *HashSet) Items() []string {
	h.mux.RLock()
	keys := make([]string, 0, len(h.m))
	for k := range h.m {
		keys = append(keys, k)
	}
	h.mux.RUnlock()
	sort.Strings(keys)
	return keys
}

// Count returns the current count of items
func (h *HashSet) Count() int {
	h.mux.RLock()
	l := len(h.m)
	h.mux.RUnlock()
	return l
}

// Contains returns set membership of item in the HashSet.
func (h *HashSet) Contains(item string) bool {
	h.mux.RLock()
	_, exists := h.m[item]
	h.mux.RUnlock()
	return exists
}
