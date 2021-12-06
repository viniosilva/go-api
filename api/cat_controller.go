package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	catApp "github.com/viniosilva/go-api/app/cat"
)

type CatController interface {
	Configure(eng *gin.RouterGroup)
	FindCats(ctx *gin.Context)
}

type CatControllerImpl struct {
	catApp catApp.CatApp
}

func NewCatController(catApp catApp.CatApp) CatController {
	return &CatControllerImpl{catApp: catApp}
}

func (impl CatControllerImpl) Configure(eng *gin.RouterGroup) {
	catGroup := eng.Group("/v1/cats")

	catGroup.GET("/", impl.FindCats)
}

// Cat FindCats godoc
// @Schemes
// @Description	Find cats route
// @Tags		cat
// @Produce		json
// @Success		200 {object} catApp.ListCatsDto
// @Router		/api/v1/cats [get]
func (impl CatControllerImpl) FindCats(ctx *gin.Context) {
	res := impl.catApp.FindCats()

	ctx.JSON(http.StatusOK, res)
}
