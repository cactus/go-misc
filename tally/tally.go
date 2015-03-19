package tally

import (
	"sync"
	"sync/atomic"
)

// Tally is a counter structure.
type Tally struct {
	c   uint64
	mux sync.Mutex
}

// Increments the counter by one.
// This operation is atomic.
func (t *Tally) Inc() {
	atomic.AddUint64(&t.c, 1)
}

/*
// Decrements the counter by one.
// This operation is atomic.
// Note that decrement will wrap.
func (t *Tally) Dec() {
	atomic.AddUint64(&t.c, ^uint64(0))
}
*/

// Returns the current value of the counter.
// This operation is atomic.
func (t *Tally) Get() uint64 {
	return atomic.LoadUint64(&t.c)
}

// Sets the counter to value.
// This operation is atomic.
func (t *Tally) Set(i uint64) {
	atomic.StoreUint64(&t.c, i)
}

// Sets the counter to 0.
// This operation is atomic.
func (t *Tally) Reset() {
	atomic.StoreUint64(&t.c, 0)
}

// Increments and returns the next counter value.
// This uses locking, to ensure that values are
// safely incremented and returned.
func (t *Tally) SafeNext() uint64 {
	t.mux.Lock()
	t.c++
	i := t.c
	t.mux.Unlock()
	return i
}
