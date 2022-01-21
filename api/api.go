package api

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/viniosilva/go-api/api/cat"
	"github.com/viniosilva/go-api/api/health"
	"github.com/viniosilva/go-api/docs"
)

type Api interface {
	Start(host string)
}

type api struct {
	health health.HealthApi
	cat    cat.CatApi
}

func NewApi(health health.HealthApi, cat cat.CatApi) Api {
	return &api{health: health, cat: cat}
}

// @title           Swagger Go API
// @version         1.0
// @description     For studies.
// @BasePath		/api
func (api *api) Start(host string) {
	router := gin.Default()
	group := router.Group("/api")
	api.health.Configure(group)
	api.cat.Configure(group)

	docs.SwaggerInfo.Host = host
	group.GET("/swagger/*a", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(host)
}
