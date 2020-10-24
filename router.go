package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"test-time-tracker/data"
	"test-time-tracker/models"
	"test-time-tracker/setup"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Ping test
	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	setup.RouterTracker(r)
	setup.RouterUser(r)
	return r
}

func main() {
	defer data.Shutdown()
	for i := models.Absent; i <= models.OOO; i++ {
		fmt.Println(i)
	}
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":8080")
}
