package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSetAdd(t *testing.T) {
	t.Parallel()
	s := NewHashSet(nil)
	s.Add("btest")
	s.Add("atest")
	assert.Equal(t, []string{"atest", "btest"}, s.Items())
}

func TestHashSetDuplicateAdd(t *testing.T) {
	t.Parallel()
	s := NewHashSet(nil)
	s.Add("btest")
	s.Add("atest")
	s.Add("atest")
	s.Add("btest")
	assert.Equal(t, []string{"atest", "btest"}, s.Items())
}

func TestHashSetNewAdd(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	assert.Equal(t, []string{"atest", "btest"}, s.Items())
}

func TestHashSetDelete(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	s.Del("btest")
	assert.Equal(t, []string{"atest"}, s.Items())
}

func TestHashSetClear(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	s.Clear()
	assert.Equal(t, []string{}, s.Items())
	assert.Equal(t, 0, s.Count())
}

func TestHashSetCount(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	assert.Equal(t, 2, s.Count())
	s.Add("ctest")
	assert.Equal(t, 3, s.Count())
	s.Del("ctest")
	assert.Equal(t, 2, s.Count())
	s.Clear()
	assert.Equal(t, 0, s.Count())
}

func TestHashSetContains(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	assert.Equal(t, true, s.Contains("atest"))
	assert.Equal(t, true, s.Contains("btest"))
	assert.Equal(t, false, s.Contains("ctest"))
}
