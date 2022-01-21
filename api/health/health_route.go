package health

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

// Health godoc
// @Schemes
// @Description	Healthcheck route
// @Tags		health
// @Produce		json
// @Success		200 {object} HealthResponse
// @Failure		503 {object} HealthResponse
// @Router		/api/health [get]
func (api *healthApi) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, HealthResponse{Status: HealthStatusOK})
}
