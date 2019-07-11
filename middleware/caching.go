package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"sumcAPI/cache"
	"sumcAPI/handlers"
	"sumcAPI/workers"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Caching(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	input, err := handlers.ScheduleHandler.ParseInput(c)
	if err != nil {
		c.JSON(400, gin.H{"result": false, "error": err.Error()})
	}
	stop := input.Stop

	if scheduleLog, exists := workers.UpdateLog[stop]; exists {
		if scheduleLog.Generating {
			//Same schedule is currently being processed by another worker
			//Try the cache at each 3 seconds to retrieve it
			for i := 1; i <= 5; i++ {
				time.Sleep(3 * time.Second)
				schedule, e := cache.Get(stop)
				if e == nil {
					c.Header("Content-Type", "application/json; charset=utf-8")
					c.String(200, schedule)
					c.Abort()
					return
				}
			}
		}
	}

	//Return cached schedule if there is available
	schedule, e := cache.Get(stop)
	if e == nil {
		if _, exists := workers.UpdateLog[stop]; exists {
			workers.UpdateLog[stop].LastCalled = time.Now()
		}
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.String(200, schedule)
		c.Abort()
		return
	}

	//Add the schedule to the update log. The Job dispatcher will
	//pass it to the workers(through queue) which will update it
	workers.UpdateLog[stop] = &workers.ScheduleLog{time.Now(), time.Now(), true}

	//Proceed to generate new schedule
	c.Next()
	workers.UpdateLog[stop].Generating = false

	//Cache only succesfull responses with schedules
	if c.Writer.Status() < 400 {
		go cache.Set(stop, blw.body.String())
	}

}
