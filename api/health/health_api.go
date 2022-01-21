package health

import (
	"github.com/gin-gonic/gin"
)

type HealthApi interface {
	Configure(eng *gin.RouterGroup)
	Health(ctx *gin.Context)
}

type healthApi struct{}

func NewApi() HealthApi {
	return &healthApi{}
}

func (api *healthApi) Configure(eng *gin.RouterGroup) {
	healthGroup := eng.Group("/health")

	healthGroup.GET("/", api.Health)
}
