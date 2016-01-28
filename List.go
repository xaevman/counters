//  ---------------------------------------------------------------------------
//
//  List.go
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
    "encoding/json"
    "fmt"
    "sync"
)

// List is a centralized container for counter objects. It supplies methods
// for adding, removing, and operating on a set of counters registered with it.
type List struct {
    lock     sync.RWMutex
    counters map[string]Counter
}

// NewList returns a reference to a new List object.
func NewList() *List {
    newObj := &List {
        counters : make(map[string]Counter),
    }

    return newObj
}

// Add registers a new counter with this counter list. If an existing counter
// with the same name already exists, an error is returned.
func (this *List) Add(c Counter) error {
    this.lock.Lock()
    defer this.lock.Unlock()

    _, ok := this.counters[c.Name()]
    if ok {
        return fmt.Errorf("Counter already exists")
    }

    this.counters[c.Name()] = c

    return nil
}

// Do performs function f on each counter instance currently registered with the
// counter list.
func (this *List) Do(f func(Counter) error) []error {
    this.lock.RLock()
    defer this.lock.RUnlock()

    errors := make([]error, 0)

    for k, _ := range this.counters {
        err := f(this.counters[k])
        if err != nil {
            errors = append(errors, err)
        }
    }

    return errors
}

// Get attempts to retreive a counter by the given name from the counter list. If
// it succeeds, a reference to the counter and an ok value of true are returned. 
// If not, the counter reference will be nil, and an ok value of false is also 
// returned.
func (this *List) Get(counterName string) (Counter, bool) {
    this.lock.RLock()
    defer this.lock.RUnlock()

    val, ok := this.counters[counterName]
    return val, ok
}

// Len returns the current number of counters contained within this counter list.
func (this *List) Len() int {
    this.lock.RLock()
    this.lock.RUnlock()

    return len(this.counters)
}

func (this *List) MarshalJSON() ([]byte, error) {
    this.lock.RLock()
    this.lock.RUnlock()

    cl := make([]Counter, 0, len(this.counters))
    for k, _ := range this.counters {
        cl = append(cl, this.counters[k])
    }

    return json.Marshal(&cl)
}

// Remove attempts to remove a counter by the given name from the counter list.
func (this *List) Remove(counterName string) {
    this.lock.Lock()
    defer this.lock.Unlock()

    delete(this.counters, counterName)
}
