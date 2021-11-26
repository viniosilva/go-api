package api

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	Start(host string)
}

type ServerImpl struct {
	health HealthController
	cat    CatController
}

func NewServer(health HealthController, cat CatController) Server {
	return &ServerImpl{health: health, cat: cat}
}

func (impl ServerImpl) Start(host string) {
	router := gin.Default()

	impl.health.Configure(router)
	impl.cat.Configure(router)

	router.Run(host)
}
