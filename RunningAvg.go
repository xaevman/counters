package counters

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

const MAX_TRACKING_SEC = 30

type RunningAvg struct {
	lock       sync.RWMutex
	name       string
	lastUpdate time.Time
	sum        KahanSum
	counter    uint64
}

func NewRunningAvg(name string) *RunningAvg {
	newObj := &RunningAvg{
		name: name,
	}

	newObj.reset(0.0)

	return newObj
}

func (ra *RunningAvg) Add(val interface{}) {
	ra.lock.Lock()
	defer ra.lock.Unlock()

	if time.Since(ra.lastUpdate).Seconds() > MAX_TRACKING_SEC {
		ra.reset(ra.getVal())
	}

	ra.sum.Add(val.(float64))
	ra.counter++
}

func (ra *RunningAvg) Get() string {
	ra.lock.RLock()
	defer ra.lock.RUnlock()

	val := ra.getVal()

	return fmt.Sprintf("%f", val)
}

func (ra *RunningAvg) GetRaw() interface{} {
	ra.lock.RLock()
	defer ra.lock.RUnlock()

	return ra.getVal()
}

func (ra *RunningAvg) Name() string {
	return ra.name
}

func (ra *RunningAvg) Set(val interface{}) {
	ra.lock.Lock()
	defer ra.lock.Unlock()

	ra.reset(val.(float64))
}

func (ra *RunningAvg) MarshalJSON() ([]byte, error) {
	ra.lock.RLock()
	defer ra.lock.RUnlock()

	var buffer bytes.Buffer

	buffer.WriteString("{")
	buffer.WriteString(fmt.Sprintf("\"key\" : \"%s\",", ra.name))
	buffer.WriteString(fmt.Sprintf("\"value\" : %.8f", ra.getVal()))
	buffer.WriteString("}")

	return buffer.Bytes(), nil
}

func (ra *RunningAvg) getVal() float64 {
	if ra.counter == 0 {
		return 0.0
	} else {
		return ra.sum.Sum() / float64(ra.counter)
	}
}

func (ra *RunningAvg) reset(initialVal float64) {
	ra.sum.Reset()

	if initialVal > 0.0 {
		ra.counter = 1
	} else {
		ra.counter = 0
		ra.sum.Add(initialVal)
	}

	ra.lastUpdate = time.Now()
}
