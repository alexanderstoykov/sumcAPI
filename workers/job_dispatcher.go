package main

import (
	"time"
	"sumcAPI/cache"
)

var UpdateLog = map[int]scheduleLog{}

type scheduleLog struct {
	lastCalled time.Time
	lastCached time.Time
	generating bool
}

/**
 * Passing bus stop numbers to channel (queue), which
 * will be read by multiple workers for generating schedule
 */
func jobDispatcher() {
	for range time.Tick(5 * time.Second) {
		for busStop, scheduleLog := range UpdateLog {
			//If the schedule was not called more than 30 minutes
			//remove it from the log&cache and stop updating it
			sinceLastCall := time.Now().Sub(scheduleLog.lastCalled).Seconds()
			if sinceLastCall > 1800 {
				cache.Mutex.Lock()
				delete(UpdateLog, busStop)
				cache.Remove(busStop)
				cache.Mutex.Unlock()

				continue
			}

			//If it was cached less than 30 seconds ago
			//skip regenerating the cache
			sinceLastCache := time.Now().Sub(scheduleLog.lastCached).Seconds()
			if sinceLastCache < 30 {
				continue
			}

			//Pass the bus stop to the workers to generate schedule
			Queue <- busStop
		}
	}
}
