package cat

import (
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/viniosilva/go-api/model"
	"github.com/viniosilva/go-api/repository"
)

func TestCatAppCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cases := map[string]struct {
		dto        CreateDto
		res        CreateResult
		errMessage string
		buildMock  func(mock *repository.MockCatRepository)
	}{
		"should create a cat when command is valid": {
			dto: CreateDto{Birthday: "2021-10-15", Name: "Lola"},
			res: CreateResult{ID: 1, Name: "Lola", Birthday: time.Date(2021, 10, 15, 0, 0, 0, 0, time.UTC)},
			buildMock: func(mock *repository.MockCatRepository) {
				mockCat := &model.Cat{ID: 1, Name: "Lola", Birthday: time.Date(2021, 10, 15, 0, 0, 0, 0, time.UTC)}
				mock.EXPECT().Create(gomock.Any()).Return(mockCat)
			},
		},
		"should throw invalidCmd error when command is invalid": {
			dto:        CreateDto{Birthday: "2021-10", Name: "Lola"},
			errMessage: "Invalid command",
			buildMock:  func(mock *repository.MockCatRepository) {},
		},
	}
	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			// given
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCatRepository := repository.NewMockCatRepository(ctrl)
			cs.buildMock(mockCatRepository)

			app := NewApp(mockCatRepository)

			// when
			res, err := app.Create(cs.dto)

			// then
			if err != nil && err.Error() == cs.errMessage {
				t.Errorf("Err\nresult: %v\nexpected: %v", err.Error(), cs.errMessage)
			}
			if !reflect.DeepEqual(res, cs.res) {
				t.Errorf("Cat\nresult: %v\nexpected: %v", res, cs.res)
			}
		})
	}
}
