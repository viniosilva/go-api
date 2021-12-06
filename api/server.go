package api

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/viniosilva/go-api/docs"
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

// @title           Swagger Go API
// @version         1.0
// @description     For studies.
// @BasePath		/api/v1
func (impl ServerImpl) Start(host string) {
	router := gin.Default()
	group := router.Group("/api")
	impl.health.Configure(group)
	impl.cat.Configure(group)

	docs.SwaggerInfo.Host = host
	group.GET("/swagger/*a", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(host)
}
