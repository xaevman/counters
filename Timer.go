//  ---------------------------------------------------------------------------
//
//  Timer.go
//
//  Copyright (c) 2015, Jared Chavez. 
//  All rights reserved.
//
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.
//
//  -----------

package counters

import (
    "fmt"
    "sync"
    "time"
)

// Timer implements the Counter interface to supply a continuously updating
// timer counter.
type Timer struct {
    lock      sync.RWMutex
    name      string
    startTime time.Time
}

// NewTimer returns a reference to a new Timer instance. The timer is not started
// until the user explicitly starts it.
func NewTimer(name string) *Timer {
    newObj := &Timer {
        name : name,
    }

    return newObj
}

// Add is not implemented in the Timer counter.
func (this *Timer) Add(val interface{}) {}

// Get returns the current number of seconds elapsed since the timer was started
// as a string.
func (this *Timer) Get() string {
    this.lock.RLock()
    defer this.lock.RUnlock()

    elapsed := time.Since(this.startTime).Seconds()

    return fmt.Sprintf("%f", elapsed)
}

// Get returns the current number of seconds elapsed since the timer was started
// as a float64.
func (this *Timer) GetRaw() interface{} {
    this.lock.RLock()
    defer this.lock.RUnlock()

    return time.Since(this.startTime).Seconds()
}

// Name returns the name of this counter instance.
func (this *Timer) Name() string {
    return this.name
}

// Set starts this timer object.
func (this *Timer) Set(val interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.startTime = time.Now()
}
