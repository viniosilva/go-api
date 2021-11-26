package app

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
		cats      []model.Cat
		buildMock func(mock *repository.MockCatRepository)
	}{
		"should return cat list when cats exists": {
			cats: []model.Cat{{ID: 1, Name: "Mimo", Birthday: time.Date(2020, 11, 20, 0, 0, 0, 0, time.UTC)}},
			buildMock: func(mock *repository.MockCatRepository) {
				mockCats := []model.Cat{{ID: 1, Name: "Mimo", Birthday: time.Date(2020, 11, 20, 0, 0, 0, 0, time.UTC)}}
				mock.EXPECT().FindCats().Return(mockCats)
			},
		},
		"should return empty cat list when cats not exists": {
			cats: []model.Cat{},
			buildMock: func(mock *repository.MockCatRepository) {
				mock.EXPECT().FindCats().Return([]model.Cat{})
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
			if !reflect.DeepEqual(res, cs.cats) {
				t.Errorf("%s\nresult: %v\nexpected: %v", "Cats", res, cs.cats)
			}
		})
	}
}
