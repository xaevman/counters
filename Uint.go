package counters

import (
    "fmt"
    "sync"
)

type Uint struct {
    lock sync.RWMutex
    name string
    val  uint64
}

func NewUint(name string) *Uint {
    newObj := &Uint {
        name : name,
    }

    return newObj
}

func (this *Uint) Add(val interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.val += val.(uint64)
}

func (this *Uint) Get() string {
    this.lock.RLock()
    defer this.lock.RUnlock()

    return fmt.Sprintf("%d", this.val)
}

func (this *Uint) GetRaw() interface{} {
    this.lock.RLock()
    defer this.lock.RUnlock()

    return this.val
}

func (this *Uint) Name() string {
    return this.name
}

func (this *Uint) Set(val interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.val = val.(uint64)
}
