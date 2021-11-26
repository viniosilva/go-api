package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController interface {
	Configure(eng *gin.Engine)
	Health(ctx *gin.Context)
}

type HealthControllerImpl struct{}

func NewHealthController() HealthController {
	return &HealthControllerImpl{}
}

func (impl HealthControllerImpl) Configure(eng *gin.Engine) {
	healthGroup := eng.Group("/health")

	healthGroup.GET("/", impl.Health)
}

func (impl HealthControllerImpl) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
