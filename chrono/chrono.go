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
	t           int64
	onceUpdater sync.Once
	mu          sync.RWMutex
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

// TimeNowString is a formatted utc time string that is updated in 1 second
// intervals.
type TimeNowString struct {
	format      string
	dateValue   atomic.Value
	onceUpdater sync.Once
}

// Get current time value string
func (t *TimeNowString) String() string {
	stamp := t.dateValue.Load()
	return stamp.(string)
}

// Force an update to the current time.
func (t *TimeNowString) Update() {
	t.dateValue.Store(time.Now().UTC().Format(t.format))
}

// Return a new TimeNowString struct
func NewTimeNowString(format string) *TimeNowString {
	t := &TimeNowString{format: format}
	t.Update()
	t.onceUpdater.Do(func() {
		go func() {
			for range time.Tick(1 * time.Second) {
				t.Update()
			}
		}()
	})
	return t
}
