// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package set

import (
	"testing"

	"github.com/dropwhile/assert"
)

func TestHashSetAdd(t *testing.T) {
	t.Parallel()
	s := NewHashSet(nil)
	s.Add("btest")
	s.Add("atest")
	assert.Equal(t, s.Items(), []string{"atest", "btest"})
}

func TestHashSetDuplicateAdd(t *testing.T) {
	t.Parallel()
	s := NewHashSet(nil)
	s.Add("btest")
	s.Add("atest")
	s.Add("atest")
	s.Add("btest")
	assert.Equal(t, s.Items(), []string{"atest", "btest"})
}

func TestHashSetNewAdd(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	assert.Equal(t, s.Items(), []string{"atest", "btest"})
}

func TestHashSetDelete(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	s.Del("btest")
	assert.Equal(t, s.Items(), []string{"atest"})
}

func TestHashSetClear(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	s.Clear()
	assert.Equal(t, s.Items(), []string{})
	assert.Equal(t, s.Count(), 0)
}

func TestHashSetCount(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	assert.Equal(t, s.Count(), 2)
	s.Add("ctest")
	assert.Equal(t, s.Count(), 3)
	s.Del("ctest")
	assert.Equal(t, s.Count(), 2)
	s.Clear()
	assert.Equal(t, s.Count(), 0)
}

func TestHashSetContains(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	assert.Equal(t, s.Contains("atest"), true)
	assert.Equal(t, s.Contains("btest"), true)
	assert.Equal(t, s.Contains("ctest"), false)
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
