package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"strconv"
	"math"
)

var Queue chan int

func init() {
	Queue = make(chan int, 50)
	for i := 1; i <= 10; i++ {
		go worker()
	}
	go jobDispatcher()
}

func main() {
	r := gin.Default()
	//r.GET("/line", serveSchedule)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func generateSchedule(busStop int) (schedule map[int][]int) {
	response := callSumc(busStop)

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

	return
}
