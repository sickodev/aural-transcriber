package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func NewHealthController() *HealthController{
	return &HealthController{}
}

func (controller *HealthController) HealthCheck(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"status":"OK",
		"message":"Service is running",
	});
}