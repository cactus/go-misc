// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package set

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestHashSetAdd(t *testing.T) {
	t.Parallel()
	s := NewHashSet(nil)
	s.Add("btest")
	s.Add("atest")
	assert.Check(t, is.DeepEqual([]string{"atest", "btest"}, s.Items()))
}

func TestHashSetDuplicateAdd(t *testing.T) {
	t.Parallel()
	s := NewHashSet(nil)
	s.Add("btest")
	s.Add("atest")
	s.Add("atest")
	s.Add("btest")
	assert.Check(t, is.DeepEqual([]string{"atest", "btest"}, s.Items()))
}

func TestHashSetNewAdd(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	assert.Check(t, is.DeepEqual([]string{"atest", "btest"}, s.Items()))
}

func TestHashSetDelete(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	s.Del("btest")
	assert.Check(t, is.DeepEqual([]string{"atest"}, s.Items()))
}

func TestHashSetClear(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	s.Clear()
	assert.Check(t, is.DeepEqual([]string{}, s.Items()))
	assert.Check(t, is.Equal(0, s.Count()))
}

func TestHashSetCount(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	assert.Check(t, is.Equal(2, s.Count()))
	s.Add("ctest")
	assert.Check(t, is.Equal(3, s.Count()))
	s.Del("ctest")
	assert.Check(t, is.Equal(2, s.Count()))
	s.Clear()
	assert.Check(t, is.Equal(0, s.Count()))
}

func TestHashSetContains(t *testing.T) {
	t.Parallel()
	s := NewHashSet([]string{"btest", "atest"})
	assert.Check(t, is.Equal(true, s.Contains("atest")))
	assert.Check(t, is.Equal(true, s.Contains("btest")))
	assert.Check(t, is.Equal(false, s.Contains("ctest")))
}

func TestHashSetUnion(t *testing.T) {
	t.Parallel()
	a := NewHashSet([]string{"1", "2", "3"})
	b := NewHashSet([]string{"3", "4", "5"})
	c := a.Union(b)

	items := c.Items()
	assert.Check(t, is.DeepEqual(items, []string{"1", "2", "3", "4", "5"}))
}

func TestHashSetIntersection(t *testing.T) {
	t.Parallel()
	a := NewHashSet([]string{"1", "2", "3"})
	b := NewHashSet([]string{"3", "4", "5"})
	c := a.Intersection(b)

	items := c.Items()
	assert.Check(t, is.DeepEqual(items, []string{"3"}))
}

func TestHashSetDifference(t *testing.T) {
	t.Parallel()
	a := NewHashSet([]string{"1", "2", "3"})
	b := NewHashSet([]string{"3", "4", "5"})
	c := a.Difference(b)

	items := c.Items()
	assert.Check(t, is.DeepEqual(items, []string{"1", "2"}))
}
