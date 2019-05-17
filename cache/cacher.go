package cache

import (
	"sync"
	"errors"
)

var Mutex = &sync.Mutex{}
var schedules = map[int]map[int][]int{}

func Get(busStop int) (map[int][]int, error) {
	if _, ok := schedules[busStop]; !ok {
		return nil, errors.New("Bus stop was not found")
	}
	return schedules[busStop], nil
}

func Set(busStop int, schedule map[int][]int) {
	schedules[busStop] = schedule
}

func Remove(busStop int) {
	delete(schedules, busStop)
}
