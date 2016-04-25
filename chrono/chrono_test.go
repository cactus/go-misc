// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package chrono

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeNow(t *testing.T) {
	t.Parallel()
	c := NewTimeNow()
	c.Update()

	now := time.Now().UTC()
	cnow := time.Unix(c.Get(), 0)
	assert.WithinDuration(t, now, cnow, 2*time.Second)
}

func BenchmarkTimeNow(b *testing.B) {
	c := NewTimeNow()
	c.Update()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get()
		}
	})
}

func TestTimeNowString(t *testing.T) {
	t.Parallel()
	c := NewTimeNowString("Mon, 02 Jan 2006 15:04 GMT")
	c.Update()

	now := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04 GMT")
	cnow := c.String()
	assert.Equal(t, now, cnow)
}

func TestTimeNowStringDelay(t *testing.T) {
	t.Parallel()
	c := NewTimeNowString("Mon, 02 Jan 2006 15:04:05 GMT")
	c.Update()
	n := c.String()
	time.Sleep(2 * time.Second)
	c.Update()
	l := c.String()

	assert.NotEqual(t, n, l,
		"Date did not update as expected: %s == %s", n, l)
}

func BenchmarkTimeNowString(b *testing.B) {
	c := NewTimeNowString("Mon, 02 Jan 2006 15:04:05 GMT")
	c.Update()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.String()
		}
	})
}
