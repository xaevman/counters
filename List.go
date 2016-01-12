package counters

import (
    "fmt"
    "sync"
)

type List struct {
    lock     sync.RWMutex
    counters map[string]Counter
}

func NewList() *List {
    newObj := &List {
        counters : make(map[string]Counter),
    }

    return newObj
}

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

func (this *List) Remove(counterName string) {
    this.lock.Lock()
    defer this.lock.Unlock()

    delete(this.counters, counterName)
}

func (this *List) Get(counterName string) (Counter, bool) {
    this.lock.RLock()
    defer this.lock.RUnlock()

    val, ok := this.counters[counterName]
    return val, ok
}

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
