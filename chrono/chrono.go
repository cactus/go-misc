// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package chrono

import (
	"sync"
	"sync/atomic"
	"time"
)

// TimeNow is a unix epoch time structure that is updated in 1 second
// intervals.
type TimeNow struct {
	mu          sync.RWMutex
	onceUpdater sync.Once
	t           int64
}

// Get current time value as unix epoc.
func (t *TimeNow) Get() int64 {
	return atomic.LoadInt64(&t.t)
}

// Force an update to the current time.
func (t *TimeNow) Update() {
	atomic.StoreInt64(&t.t, time.Now().UTC().Unix())
}

// Return a new TimeNow struct
func NewTimeNow() *TimeNow {
	t := &TimeNow{t: time.Now().UTC().Unix()}
	t.onceUpdater.Do(func() {
		go func() {
			for range time.Tick(1 * time.Second) {
				t.Update()
			}
		}()
	})
	return t
}

// internal/global TimeNow struct
var nowTimer = NewTimeNow()

// Gets the current TimeNow time in unix epoc.
func GetTime(d int64) int64 {
	t := nowTimer.Get()
	if d > 0 {
		t = t + d
	}
	return t
}
