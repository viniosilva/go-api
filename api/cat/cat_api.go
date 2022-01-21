package cat

import (
	"github.com/gin-gonic/gin"
	"github.com/viniosilva/go-api/app/cat"
)

type CatApi interface {
	Configure(eng *gin.RouterGroup)
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
}

type catApi struct {
	app cat.CatApp
}

func NewApi(app cat.CatApp) CatApi {
	return &catApi{app: app}
}

func (api *catApi) Configure(eng *gin.RouterGroup) {

	catGroup := eng.Group("/v1/cats")

	catGroup.POST("/", api.Create)
	catGroup.GET("/", api.List)
}
