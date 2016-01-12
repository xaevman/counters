//  ---------------------------------------------------------------------------
//
//  Float.go
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

// Float implements the Counter interface for the floating point data type.
type Float struct {
    lock sync.RWMutex
    name string
    val  float64
}

// NewFloat returns a reference to an instance of a new Float counter object.
func NewFloat(name string) *Float {
    newObj := &Float {
        name : name,
    }

    return newObj
}

// Add type-asserts teh supplied value as a float64 and adds it to the existing
// stored value.
func (this *Float) Add(val interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.val += val.(float64)
}

// Get returns the current stored value as a string.
func (this *Float) Get() string {
    this.lock.RLock()
    defer this.lock.RUnlock()

    return fmt.Sprintf("%f", this.val)
}

// GetRaw reutrns the current stored value as a float64.
func (this *Float) GetRaw() interface{} {
    this.lock.RLock()
    defer this.lock.RUnlock()

    return this.val
}

// Name returns the name of this counter instance.
func (this *Float) Name() string {
    return this.name
}

// Set type-asserts the supplied value as a float64 and sets the current
// value of the counter to that value.
func (this *Float) Set(val interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.val = val.(float64)
}
