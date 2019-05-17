package handlers

import (
	"github.com/gin-gonic/gin"
	"time"
	"strconv"
	"math"
	"sumcAPI/services"
	"errors"
)

type Input struct {
	Stop int `form:"number"`
	Line int `form:"line"`
}

func ParseInput(c *gin.Context) (input Input, err error) {
	err = c.Bind(&input)
	stop , passed := c.Params.Get("number")
	if !passed {
		err = errors.New("Number was not passed")
		return
	}

	input.Stop, _ = strconv.Atoi(stop)
	return
}

func Handler(c *gin.Context) {
	input, err := ParseInput(c)
	if err != nil {
		c.JSON(400, gin.H{"result": false, "error": err.Error()})
	}
	schedule := GenerateSchedule(input.Stop)
	c.JSON(200, schedule)
}

func GenerateSchedule(busStop int) map[int][]int {
	response := services.CallSumc(busStop)
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
