package counters

import (
    "fmt"
    "sync"
)

type Int struct {
    lock sync.RWMutex
    name string
    val  int64
}

func NewInt(name string) *Int {
    newObj := &Int {
        name : name,
    }

    return newObj
}

func (this *Int) Add(val interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.val += val.(int64)
}

func (this *Int) Get() string {
    this.lock.RLock()
    defer this.lock.RUnlock()

    return fmt.Sprintf("%d", this.val)
}

func (this *Int) GetRaw() interface{} {
    this.lock.RLock()
    defer this.lock.RUnlock()

    return this.val
}

func (this *Int) Name() string {
    return this.name
}

func (this *Int) Set(val interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.val = val.(int64)
}
