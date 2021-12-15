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

func TestAppCatAppFindCats(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cases := map[string]struct {
		res       ListCatsDto
		buildMock func(mock *repository.MockCatRepository)
	}{
		"should return cat list when cats exists": {
			res: ListCatsDto{Data: []CatDto{{ID: 1, Name: "Mimo", Birthday: "2020-11-20"}}},
			buildMock: func(mock *repository.MockCatRepository) {
				mockCats := []model.Cat{{ID: 1, Name: "Mimo", Birthday: time.Date(2020, 11, 20, 0, 0, 0, 0, time.UTC)}}
				mock.EXPECT().Find().Return(mockCats)
			},
		},
		"should return empty cat list when cats not exists": {
			res: ListCatsDto{Data: []CatDto{}},
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

			mockCatApp := repository.NewMockCatRepository(ctrl)
			cs.buildMock(mockCatApp)

			app := NewCatApp(mockCatApp)

			// when
			res := app.FindCats()

			// then
			if !reflect.DeepEqual(res, cs.res) {
				t.Errorf("%s\nresult: %v\nexpected: %v", "Cats", res, cs.res)
			}
		})
	}
}

func TestAppCatAppCreateCat(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cases := map[string]struct {
		cmd        CatCmd
		res        CatDto
		errMessage string
		buildMock  func(mock *repository.MockCatRepository)
	}{
		"should create a cat when command is valid": {
			cmd: CatCmd{Birthday: "2021-10-15", Name: "Lola"},
			res: CatDto{ID: 1, Name: "Lola", Birthday: "2021-10-15"},
			buildMock: func(mock *repository.MockCatRepository) {
				mockCat := &model.Cat{ID: 1, Name: "Lola", Birthday: time.Date(2021, 10, 15, 0, 0, 0, 0, time.UTC)}
				mock.EXPECT().Create(gomock.Any()).Return(mockCat)
			},
		},
		"should throw invalidCmd error when command is invalid": {
			cmd:        CatCmd{Birthday: "2021-10", Name: "Lola"},
			errMessage: "Invalid command",
			buildMock:  func(mock *repository.MockCatRepository) {},
		},
	}
	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			// given
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCatApp := repository.NewMockCatRepository(ctrl)
			cs.buildMock(mockCatApp)

			app := NewCatApp(mockCatApp)

			// when
			res, err := app.CreateCat(cs.cmd)

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
