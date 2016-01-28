//  ---------------------------------------------------------------------------
//
//  Uint.go
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
    "bytes"
    "fmt"
    "sync"
)

// Uint implements the Counter interface for the uint data type.
type Uint struct {
    lock sync.RWMutex
    name string
    val  uint64
}

// NewUint returns a reference to a new Uint counter instance.
func NewUint(name string) *Uint {
    newObj := &Uint {
        name : name,
    }

    return newObj
}

// Add type-asserts the given value as a uint64 and adds it to the current
// counter value.
func (this *Uint) Add(val interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.val += val.(uint64)
}

// Get returns the current counter value as a string.
func (this *Uint) Get() string {
    this.lock.RLock()
    defer this.lock.RUnlock()

    return fmt.Sprintf("%d", this.val)
}

// GetRaw returns the current counter value as a uint64.
func (this *Uint) GetRaw() interface{} {
    this.lock.RLock()
    defer this.lock.RUnlock()

    return this.val
}

func (this *Uint) MarshalJSON() ([]byte, error) {
    this.lock.RLock()
    defer this.lock.RUnlock()

    var buffer bytes.Buffer

    buffer.WriteString("{")
    buffer.WriteString(fmt.Sprintf("\"key\" : \"%s\",", this.name))
    buffer.WriteString(fmt.Sprintf("\"value\" : %d", this.val))
    buffer.WriteString("}")

    return buffer.Bytes(), nil
}

// Name returns the name of this counter instance.
func (this *Uint) Name() string {
    return this.name
}

// Set type-asserts the given value as uint64 and sets the current counter
// value to that value.
func (this *Uint) Set(val interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.val = val.(uint64)
}
