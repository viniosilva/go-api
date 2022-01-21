package cat

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniosilva/go-api/api/dto"
)

type CatResponse struct {
	ID       int    `json:"id"`
	Birthday string `json:"birthday"`
	Name     string `json:"name"`
}

type FindCatsResponse struct {
	Data []CatResponse `json:"data"`
}

// Cat FindCats godoc
// @Schemes
// @Description	Find cats route
// @Tags		cat
// @Produce		json
// @Success		200 {object} FindCatsResponse
// @Router		/api/v1/cats [get]
func (api *catApi) List(ctx *gin.Context) {
	res, err := api.app.List()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponseDto{Message: "Internal server error"})
		return
	}

	data := []CatResponse{}

	for _, cat := range res.Data {
		data = append(data, CatResponse{
			ID:       cat.ID,
			Name:     cat.Name,
			Birthday: cat.Birthday.Format("2006-01-02"),
		})
	}

	ctx.JSON(http.StatusOK, FindCatsResponse{Data: data})
}
