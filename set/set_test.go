// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

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

func TestHashSetUnion(t *testing.T) {
	t.Parallel()
	a := NewHashSet([]string{"1", "2", "3"})
	b := NewHashSet([]string{"3", "4", "5"})
	c := a.Union(b)

	items := c.Items()
	assert.Equal(t, items, []string{"1", "2", "3", "4", "5"})
}

func TestHashSetIntersection(t *testing.T) {
	t.Parallel()
	a := NewHashSet([]string{"1", "2", "3"})
	b := NewHashSet([]string{"3", "4", "5"})
	c := a.Intersection(b)

	items := c.Items()
	assert.Equal(t, items, []string{"3"})
}

func TestHashSetDifference(t *testing.T) {
	t.Parallel()
	a := NewHashSet([]string{"1", "2", "3"})
	b := NewHashSet([]string{"3", "4", "5"})
	c := a.Difference(b)

	items := c.Items()
	assert.Equal(t, items, []string{"1", "2"})
}
