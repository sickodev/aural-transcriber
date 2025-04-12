package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sickodev/aural-transcriber/app/controller"
)

func main() {
	fmt.Printf("Hello World")

	r := gin.Default()

	api := r.Group("/api")
	{
		healthController := controller.NewHealthController()
		api.GET("/health", healthController.HealthCheck)
	}

	r.Run(":8000")
}
