package tally

import (
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTally(t *testing.T) {
	t.Parallel()
	c := &Tally{}
	assert.Equal(t, uint64(0), c.Get())
	c.Inc()
	assert.Equal(t, uint64(1), c.Get())
	c.Inc()
	assert.Equal(t, uint64(2), c.Get())
}

func TestTallyConcurrent(t *testing.T) {
	t.Parallel()
	c := &Tally{}

	var wg sync.WaitGroup
	rezchan := make(chan uint64, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			rezchan <- c.SafeNext()
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
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, x)
}
