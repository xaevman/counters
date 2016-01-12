package counters

import (
    "fmt"
    "sync"
)

type Float struct {
    lock sync.RWMutex
    name string
    val  float64
}

func NewFloat(name string) *Float {
    newObj := &Float {
        name : name,
    }

    return newObj
}

func (this *Float) Add(val interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.val += val.(float64)
}

func (this *Float) Get() string {
    this.lock.RLock()
    defer this.lock.RUnlock()

    return fmt.Sprintf("%f", this.val)
}

func (this *Float) GetRaw() interface{} {
    this.lock.RLock()
    defer this.lock.RUnlock()

    return this.val
}

func (this *Float) Name() string {
    return this.name
}

func (this *Float) Set(val interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.val = val.(float64)
}
