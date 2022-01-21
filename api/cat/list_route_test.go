package cat

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/viniosilva/go-api/app/cat"
)

func TestApiCatList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cases := map[string]struct {
		status    int
		body      FindCatsResponse
		buildMock func(mock *cat.MockCatApp)
	}{
		"should return cat list when cats exists": {
			status: http.StatusOK,
			body:   FindCatsResponse{Data: []CatResponse{{ID: 1, Name: "Mimo", Birthday: "2020-11-20"}}},
			buildMock: func(mock *cat.MockCatApp) {
				mockCats := cat.ListResult{Data: []cat.CatResult{{ID: 1, Name: "Mimo", Birthday: time.Date(2020, 11, 20, 0, 0, 0, 0, time.UTC)}}}
				mock.EXPECT().List().Return(mockCats, nil)
			},
		},
		"should return empty cat list when cats not exists": {
			status: http.StatusOK,
			body:   FindCatsResponse{Data: []CatResponse{}},
			buildMock: func(mock *cat.MockCatApp) {
				mock.EXPECT().List().Return(cat.ListResult{Data: []cat.CatResult{}}, nil)
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			// given
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCatApp := cat.NewMockCatApp(ctrl)
			cs.buildMock(mockCatApp)

			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			api := NewApi(mockCatApp)

			// when
			api.List(ctx)

			var body FindCatsResponse
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
