package services

import (
	"encoding/json"
	"math"
	"strconv"
	"sumcAPI/interfaces"
	"time"
)

type SumcGenerator struct {
	provider interfaces.ScheduleProvider
}

var Generator SumcGenerator

func NewGenerator(provider interfaces.ScheduleProvider) SumcGenerator {
	return SumcGenerator{provider: provider}
}

func (this *SumcGenerator) GenerateSchedule(busStop int) map[int][]int {
	bytes := this.provider.CallAPI(busStop)
	response := SumcResponse{}
	json.Unmarshal(bytes, &response)

	schedule := make(map[int][]int)
	now := time.Now().UTC().Add(3 * time.Hour)
	ymd := now.Format("2006-01-02")

	for _, line := range response.Lines {
		lineNumber, _ := strconv.Atoi(line.Name)
		for _, arrival := range line.Arrivals {
			if _, ok := schedule[lineNumber]; !ok {
				schedule[lineNumber] = []int{}
			}
			tt, _ := time.Parse("2006-01-02 15:04:05", ymd+" "+arrival.Time)
			minutes := math.Round(tt.Sub(now).Minutes())
			schedule[lineNumber] = append(schedule[lineNumber], int(minutes))
		}
	}

	return schedule
}
