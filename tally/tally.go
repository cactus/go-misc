// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package tally

import (
	"sync/atomic"
)

// Tally is a counter structure.
type Tally struct {
	c uint64
}

// Inc increments the counter by one.
// This operation is atomic.
func (t *Tally) Inc() {
	atomic.AddUint64(&t.c, 1)
}

// Dec decrements the counter by one.
// This operation is atomic.
// Note that decrement will wrap.
func (t *Tally) Dec() {
	atomic.AddUint64(&t.c, ^uint64(0))
}

// Get returns the current value of the counter.
// This operation is atomic.
func (t *Tally) Get() uint64 {
	return atomic.LoadUint64(&t.c)
}

// Set sets the counter to value.
// This operation is atomic.
func (t *Tally) Set(i uint64) {
	atomic.StoreUint64(&t.c, i)
}

// Reset sets the counter to 0.
// This operation is atomic.
func (t *Tally) Reset() {
	atomic.StoreUint64(&t.c, 0)
}

// Next increments and returns the next counter value.
func (t *Tally) Next() uint64 {
	return atomic.AddUint64(&t.c, 1)
}
