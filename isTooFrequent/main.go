package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type RateLimit interface {
	IsTooFrequent() bool
	NowTime() time.Time
}

type rateLimit struct {
	timestamps []time.Time
	timeWindow time.Duration
	callLimit int
	mu sync.Mutex
}
/*
	---
	
*/
func (r *rateLimit) IsTooFrequent() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	currentTime := r.NowTime()
	r.timestamps = append(r.timestamps, currentTime)
	threshold := currentTime.Add(-r.timeWindow)
	thresholdIndex := sort.Search(len(r.timestamps), func (i int) bool {
		return r.timestamps[i].After(threshold) || r.timestamps[i].Equal(threshold)
	})
	r.timestamps = r.timestamps[thresholdIndex:]
	return len(r.timestamps) > r.callLimit
}

func (r *rateLimit) NowTime() time.Time {
	return time.Now()
}

func NewRateLimit(timeWindow time.Duration, callLimit int) RateLimit {
	return &rateLimit{
		timestamps: []time.Time{},
		timeWindow: timeWindow,
		callLimit: callLimit,
	}
}

func main() {
	var wait sync.WaitGroup

	rateLimit := NewRateLimit(5*time.Minute, 10)
	for range 12 {
		wait.Add(1)
		go func() {
			defer wait.Done()
			if rateLimit.IsTooFrequent() {
				fmt.Println("too frequent")
			} else {
				fmt.Println("ok")
			}
		}()
	}

	wait.Wait()
}