package main

import (
	"time"
	"sumcAPI/cache"
)

func worker() {
	for range time.Tick(time.Second * 40) {
		busStop := <-Queue
		schedule := generateSchedule(busStop)
		cache.Set(busStop, schedule)
	}
}
