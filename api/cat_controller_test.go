package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	catApp "github.com/viniosilva/go-api/app/cat"
)

func TestApiCatControllerFindCats(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cases := map[string]struct {
		status    int
		body      catApp.ListCatsDto
		buildMock func(mock *catApp.MockCatApp)
	}{
		"should return cat list when cats exists": {
			status: http.StatusOK,
			body:   catApp.ListCatsDto{Data: []catApp.CatDto{{ID: 1, Name: "Mimo", Birthday: "2020-11-20"}}},
			buildMock: func(mock *catApp.MockCatApp) {
				mockCats := catApp.ListCatsDto{Data: []catApp.CatDto{{ID: 1, Name: "Mimo", Birthday: "2020-11-20"}}}
				mock.EXPECT().FindCats().Return(mockCats)
			},
		},
		"should return empty cat list when cats not exists": {
			status: http.StatusOK,
			body:   catApp.ListCatsDto{Data: []catApp.CatDto{}},
			buildMock: func(mock *catApp.MockCatApp) {
				mock.EXPECT().FindCats().Return(catApp.ListCatsDto{Data: []catApp.CatDto{}})
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			// given
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCatApp := catApp.NewMockCatApp(ctrl)
			cs.buildMock(mockCatApp)

			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			controller := NewCatController(mockCatApp)

			// when
			controller.FindCats(ctx)

			var body catApp.ListCatsDto
			json.Unmarshal(res.Body.Bytes(), &body)

			// then
			if res.Code != cs.status {
				t.Errorf("%s\nresult:\t\t%v\nexpected:\t%v", "StatusCode", res.Code, cs.body)
			}
			if !reflect.DeepEqual(body, cs.body) {
				t.Errorf("%s\nresult:\t\t%v\nexpectedBody:\t%v", "Body", body, cs.body)
			}
		})
	}
}

func TestApiCatControllerCreateCat(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cases := map[string]struct {
		payload   string
		status    int
		body      catApp.CatDto
		buildMock func(mock *catApp.MockCatApp)
	}{
		"should create a cat when json is valid": {
			payload: `{"birthday": "2000-01-01", "name": "Lola"}`,
			status:  http.StatusCreated,
			body:    catApp.CatDto{ID: 1, Name: "Mimo", Birthday: "2020-11-20"},
			buildMock: func(mock *catApp.MockCatApp) {
				mockCats := catApp.CatDto{ID: 1, Name: "Mimo", Birthday: "2020-11-20"}
				mock.EXPECT().CreateCat(gomock.Any()).Return(mockCats, nil)
			},
		},
		"should throw BadRequest exception when json is invalid": {
			payload:   `{"name": 1}`,
			status:    http.StatusBadRequest,
			buildMock: func(mock *catApp.MockCatApp) {},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			// given
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCatApp := catApp.NewMockCatApp(ctrl)
			cs.buildMock(mockCatApp)

			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			controller := NewCatController(mockCatApp)
			ctx.Request = httptest.NewRequest("POST", "/api/v1/cats", strings.NewReader(cs.payload))

			// when
			controller.CreateCat(ctx)

			var body catApp.CatDto
			json.Unmarshal(res.Body.Bytes(), &body)

			// then
			if res.Code != cs.status {
				t.Errorf("%s\nresult:\t\t%v\nexpected:\t%v", "StatusCode", res.Code, cs.status)
			}
			if !reflect.DeepEqual(body, cs.body) {
				t.Errorf("%s\nresult:\t\t%v\nexpectedBody:\t%v", "Body", body, cs.body)
			}
		})
	}
}
