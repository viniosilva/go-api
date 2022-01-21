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

func TestCatAppList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cases := map[string]struct {
		errMessage string
		res        ListResult
		buildMock  func(mock *repository.MockCatRepository)
	}{
		"should return cat list when cats exists": {
			res: ListResult{Data: []CatResult{{ID: 1, Name: "Mimo", Birthday: time.Date(2020, 11, 20, 0, 0, 0, 0, time.UTC)}}},
			buildMock: func(mock *repository.MockCatRepository) {
				mockCats := []model.Cat{{ID: 1, Name: "Mimo", Birthday: time.Date(2020, 11, 20, 0, 0, 0, 0, time.UTC)}}
				mock.EXPECT().Find().Return(mockCats)
			},
		},
		"should return empty cat list when cats not exists": {
			res: ListResult{Data: []CatResult{}},
			buildMock: func(mock *repository.MockCatRepository) {
				mock.EXPECT().Find().Return([]model.Cat{})
			},
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
			res, err := app.List()

			// then
			if err != nil && err.Error() == cs.errMessage {
				t.Errorf("Err\nresult: %v\nexpected: %v", err.Error(), cs.errMessage)
			}

			if !reflect.DeepEqual(res, cs.res) {
				t.Errorf("%s\nresult: %v\nexpected: %v", "Cats", res, cs.res)
			}
		})
	}
}
