package main

import (
	"github.com/gin-gonic/gin"
	"sumcAPI/middleware"
	"sumcAPI/handlers"
	"sumcAPI/workers"
)

func init() {
	workers.Queue = make(chan int, 50)
	for i := 1; i <= 10; i++ {
		go workers.Worker()
	}
	go workers.JobDispatcher()
}

func main() {
	r := gin.Default()
	r.Use(middleware.Caching)
	r.GET("/stop/:number/", handlers.Handler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
