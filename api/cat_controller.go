package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	catApp "github.com/viniosilva/go-api/app/cat"
)

type CatController interface {
	Configure(eng *gin.RouterGroup)
	CreateCat(ctx *gin.Context)
	FindCats(ctx *gin.Context)
}

type catControllerImpl struct {
	catApp catApp.CatApp
}

func NewCatController(catApp catApp.CatApp) CatController {
	return &catControllerImpl{catApp: catApp}
}

func (impl catControllerImpl) Configure(eng *gin.RouterGroup) {
	catGroup := eng.Group("/v1/cats")

	catGroup.POST("/", impl.CreateCat)
	catGroup.GET("/", impl.FindCats)
}

// Cat CreateCat godoc
// @Schemes
// @Description	Create cat route
// @Tags		cat
// @Consume		json
// @Produce		json
// @Success		201	{object} catApp.CatDto
// @Failure		400 {object} ErrorResponse
// @Param 		cat body catApp.CatCmd true "Cat"
// @Router		/api/v1/cats [post]
func (impl catControllerImpl) CreateCat(ctx *gin.Context) {
	var cmd catApp.CatCmd
	err := ctx.BindJSON(&cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	res, err := impl.catApp.CreateCat(cmd)

	if err != nil {
		switch err.(type) {
		case *catApp.InvalidCmdError:
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

// Cat FindCats godoc
// @Schemes
// @Description	Find cats route
// @Tags		cat
// @Produce		json
// @Success		200 {object} catApp.ListCatsDto
// @Router		/api/v1/cats [get]
func (impl catControllerImpl) FindCats(ctx *gin.Context) {
	res := impl.catApp.FindCats()

	ctx.JSON(http.StatusOK, res)
}
