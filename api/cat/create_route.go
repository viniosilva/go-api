package cat

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniosilva/go-api/api/dto"
	"github.com/viniosilva/go-api/app/cat"
	"github.com/viniosilva/go-api/exception"
)

type CreateCatRequest struct {
	Birthday string `json:"birthday" example:"2000-01-01"`
	Name     string `json:"name" example:"Mimoso"`
}

type CreateCatResponse struct {
	ID       int    `json:"id"`
	Birthday string `json:"birthday"`
	Name     string `json:"name"`
}

// Cat CreateCat godoc
// @Schemes
// @Description	Create cat route
// @Tags		cat
// @Consume		json
// @Produce		json
// @Success		201	{object} CreateCatResponse
// @Failure		400 {object} dto.ErrorResponse
// @Param 		cat body CreateCatRequest true "Cat"
// @Router		/api/v1/cats [post]
func (api *catApi) Create(ctx *gin.Context) {
	var req CreateCatRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: err.Error()})
		return
	}

	res, err := api.app.Create(cat.CreateDto{
		Birthday: req.Birthday,
		Name:     req.Name,
	})

	if err != nil {
		switch err.(type) {
		case *exception.InvalidDateException:
			ctx.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, dto.ErrorResponseDto{Message: "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, CreateCatResponse{
		ID:       res.ID,
		Birthday: res.Birthday.Format("2006-01-02"),
		Name:     res.Name,
	})
}
