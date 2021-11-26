package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniosilva/go-api/app"
)

type CatController interface {
	Configure(eng *gin.Engine)
	FindCats(ctx *gin.Context)
}

type CatControllerImpl struct {
	catApp app.CatApp
}

func NewCatController(catApp app.CatApp) CatController {
	return &CatControllerImpl{catApp: catApp}
}

func (impl CatControllerImpl) Configure(eng *gin.Engine) {
	catGroup := eng.Group("/cats")

	catGroup.GET("/", impl.FindCats)
}

func (impl CatControllerImpl) FindCats(ctx *gin.Context) {
	res := impl.catApp.FindCats()

	ctx.JSON(http.StatusOK, res)
}
