package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthStatus string

const (
	HealthStatusOK                  HealthStatus = "OK"
	HealthStatusInternalServerError HealthStatus = "Internal server error"
)

type HealthResponse struct {
	Status HealthStatus `json:"status"`
}

type HealthController interface {
	Configure(eng *gin.RouterGroup)
	Health(ctx *gin.Context)
}

type HealthControllerImpl struct{}

func NewHealthController() HealthController {
	return &HealthControllerImpl{}
}

func (impl HealthControllerImpl) Configure(eng *gin.RouterGroup) {
	healthGroup := eng.Group("/health")

	healthGroup.GET("/", impl.Health)
}

// Health godoc
// @Schemes
// @Description	Healthcheck route
// @Tags		health
// @Produce		json
// @Success		200 {object} HealthResponse
// @Failure		503 {object} HealthResponse
// @Router		/api/health [get]
func (impl HealthControllerImpl) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, HealthResponse{Status: HealthStatusOK})
}
