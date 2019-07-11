package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"sumcAPI/interfaces"
)

type Input struct {
	Stop int `form:"number"`
	Line int `form:"line"`
}

type Handler struct {
	ScheduleGenerator interfaces.ScheduleGenerator
}

var ScheduleHandler Handler

func NewHandler(generator interfaces.ScheduleGenerator) Handler {

	return Handler{
		ScheduleGenerator: generator,
	}
}

func (*Handler) ParseInput(c *gin.Context) (input Input, err error) {
	err = c.Bind(&input)
	stop, passed := c.Params.Get("number")
	if !passed {
		err = errors.New("Number was not passed")
		return
	}

	input.Stop, _ = strconv.Atoi(stop)
	return
}

func (this *Handler) Serve(c *gin.Context) {
	input, err := this.ParseInput(c)
	if err != nil {
		c.JSON(400, gin.H{"result": false, "error": err.Error()})
	}
	schedule := this.ScheduleGenerator.GenerateSchedule(input.Stop)
	c.JSON(200, schedule)
}
