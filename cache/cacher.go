package cache

import (
	"sync"
	"errors"
)

var Mutex = &sync.Mutex{}
var schedules = map[int]string{}

func Get(busStop int) (string, error) {
	if _, ok := schedules[busStop]; !ok {
		return "", errors.New("Bus stop was not found")
	}
	return schedules[busStop], nil
}

func Set(busStop int, schedule string) {
	schedules[busStop] = schedule
}

func Remove(busStop int) {
	delete(schedules, busStop)
}
