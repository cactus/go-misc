// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package tally

import (
	"sort"
	"sync"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTally(t *testing.T) {
	t.Parallel()
	c := &Tally{}
	assert.Check(t, is.Equal(uint64(0), c.Get()))
	c.Inc()
	assert.Check(t, is.Equal(uint64(1), c.Get()))
	c.Inc()
	assert.Check(t, is.Equal(uint64(2), c.Get()))
	c.Set(42)
	assert.Check(t, is.Equal(uint64(42), c.Get()))
	c.Reset()
	assert.Check(t, is.Equal(uint64(0), c.Get()))
	c.Inc()
	assert.Check(t, is.Equal(uint64(1), c.Get()))
}

func TestTallyConcurrent(t *testing.T) {
	t.Parallel()
	c := &Tally{}

	var wg sync.WaitGroup
	rezchan := make(chan uint64, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			rezchan <- c.Next()
			wg.Done()
		}()
	}
	wg.Wait()

	x := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		y := <-rezchan
		x = append(x, int(y))
	}
	sort.Ints(x)
	assert.Check(t, is.DeepEqual([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, x))
}
