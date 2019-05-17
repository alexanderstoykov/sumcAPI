package workers

import (
	"time"
	"sumcAPI/cache"
	"sumcAPI/handlers"
	"encoding/json"
)

func Worker() {
	for range time.Tick(time.Second * 40) {
		busStop := <-Queue
		schedule := handlers.GenerateSchedule(busStop)
		bytes, _ := json.Marshal(schedule)
		cache.Set(busStop, string(bytes))
	}
}
