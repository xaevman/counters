package counters

import (
	"time"
)

type PerfTimer struct {
	avg       Counter
	startTime time.Time
}

func NewPerfTimer(name string) *PerfTimer {
	return &PerfTimer{
		avg: NewRunningAvg(name),
	}
}

func (pt *PerfTimer) Counters() Counter {
	return pt.avg
}

func (pt *PerfTimer) StartSample() {
	pt.startTime = time.Now()
}

func (pt *PerfTimer) EndSample() {
	elapsed := time.Since(pt.startTime)
	pt.avg.Add(elapsed.Seconds())
}
