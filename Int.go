//  ---------------------------------------------------------------------------
//
//  Int.go
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
)

// Int implements the Counter interface for the integer data type.
type Int struct {
    lock sync.RWMutex
    name string
    val  int64
}

// NewInt returns a reference to a new Int counter object.
func NewInt(name string) *Int {
    newObj := &Int {
        name : name,
    }

    return newObj
}

// Add type-asserts the supplied value as an int64 and adds it to the current
// counter value.
func (this *Int) Add(val interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.val += val.(int64)
}

// Get returns the current counter value as a string.
func (this *Int) Get() string {
    this.lock.RLock()
    defer this.lock.RUnlock()

    return fmt.Sprintf("%d", this.val)
}

// GetRaw returns the current counter value as an int64.
func (this *Int) GetRaw() interface{} {
    this.lock.RLock()
    defer this.lock.RUnlock()

    return this.val
}

// Name returns the name of this counter instance.
func (this *Int) Name() string {
    return this.name
}

// Set type-asserts the supplied value as an int64 and stores it as the new
// counter value.
func (this *Int) Set(val interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.val = val.(int64)
}
