package main

import (
	"github.com/gin-gonic/gin"
	"sumcAPI/handlers"
	"sumcAPI/middleware"
	"sumcAPI/services"
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
	handlers.ScheduleHandler = InitScheduleHandler()

	r := gin.Default()
	r.Use(middleware.Caching)
	r.GET("/stop/:number/", handlers.ScheduleHandler.Serve)
	r.Run() // listen and serve on 0.0.0.0:8080
}

/**
 * DI
 */
func InitScheduleHandler() handlers.Handler {
	provider := services.SumcProvider{}
	generator := services.NewGenerator(&provider)
	handler := handlers.NewHandler(&generator)
	services.Generator = generator //Exposing the generator so it can be used in middleware&caching

	return handler
}

/**
 * DI for testing (mock API)
 */
func InitializeMockSumcHandler() {
	//TODO
}
