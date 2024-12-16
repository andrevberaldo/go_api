package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func NewHealthController() HealthController {
	return HealthController{}
}

func InitializeHealthController(server *gin.Engine) {
	HealthController := NewHealthController()
	server.GET("/health", HealthController.CheckHealth)
}

func (hc *HealthController) CheckHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "up",
	})
}
