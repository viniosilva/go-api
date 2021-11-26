package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/viniosilva/go-api/app"
	"github.com/viniosilva/go-api/model"
)

func TestApiCatControllerFindCats(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cases := map[string]struct {
		status    int
		body      []map[string]interface{}
		buildMock func(mock *app.MockCatApp)
	}{
		"should return cat list when cats exists": {
			status: http.StatusOK,
			body: []map[string]interface{}{
				{"ID": 1.0, "Name": "Mimo", "Birthday": "2020-11-20T00:00:00Z"},
			},
			buildMock: func(mock *app.MockCatApp) {
				mockCats := []model.Cat{{ID: 1, Name: "Mimo", Birthday: time.Date(2020, 11, 20, 0, 0, 0, 0, time.UTC)}}
				mock.EXPECT().FindCats().Return(mockCats)
			},
		},
		"should return empty cat list when cats not exists": {
			status: http.StatusOK,
			body:   []map[string]interface{}{},
			buildMock: func(mock *app.MockCatApp) {
				mock.EXPECT().FindCats().Return([]model.Cat{})
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			// given
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCatApp := app.NewMockCatApp(ctrl)
			cs.buildMock(mockCatApp)

			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			controller := NewCatController(mockCatApp)

			// when
			controller.FindCats(ctx)

			var body []map[string]interface{}
			json.Unmarshal(res.Body.Bytes(), &body)

			// then
			if res.Code != cs.status {
				t.Errorf("%s\nresult: %v\nexpected: %v", "StatusCode", res.Code, cs.body)
			}
			if !reflect.DeepEqual(body, cs.body) {
				t.Errorf("%s\nresult: %s\nexpectedBody: %s", "Body", body, cs.body)
			}
		})
	}
}
