package counters

import (
    "fmt"
    "sync"
    "time"
)

type Timer struct {
    lock      sync.RWMutex
    name      string
    startTime time.Time
}

func NewTimer(name string) *Timer {
    newObj := &Timer {
        name : name,
    }

    return newObj
}

func (this *Timer) Add(val interface{}) {}

func (this *Timer) Get() string {
    this.lock.RLock()
    defer this.lock.RUnlock()

    elapsed := time.Since(this.startTime).Seconds()

    return fmt.Sprintf("%f", elapsed)
}

func (this *Timer) GetRaw() interface{} {
    this.lock.RLock()
    defer this.lock.RUnlock()

    return time.Since(this.startTime).Seconds()
}

func (this *Timer) Name() string {
    return this.name
}

func (this *Timer) Set(val interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.startTime = time.Now()
}
