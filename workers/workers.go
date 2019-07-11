package workers

import (
	"encoding/json"
	"sumcAPI/cache"
	"sumcAPI/services"
	"time"
)

func Worker() {
	for range time.Tick(time.Second * 40) {
		busStop := <-Queue
		schedule := services.Generator.GenerateSchedule(busStop)
		bytes, _ := json.Marshal(schedule)
		cache.Set(busStop, string(bytes))
	}
}
