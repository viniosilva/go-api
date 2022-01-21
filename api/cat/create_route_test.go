package cat

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/viniosilva/go-api/app/cat"
)

func TestApiCatCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cases := map[string]struct {
		payload   string
		status    int
		body      CreateCatResponse
		buildMock func(mock *cat.MockCatApp)
	}{
		"should create a cat when json is valid": {
			payload: `{"birthday": "2000-01-01", "name": "Lola"}`,
			status:  http.StatusCreated,
			body:    CreateCatResponse{ID: 1, Name: "Mimo", Birthday: "2020-11-20"},
			buildMock: func(mock *cat.MockCatApp) {
				mockCats := cat.CreateResult{ID: 1, Name: "Mimo", Birthday: time.Date(2020, 11, 20, 0, 0, 0, 0, time.UTC)}
				mock.EXPECT().Create(gomock.Any()).Return(mockCats, nil)
			},
		},
		"should throw BadRequest exception when json is invalid": {
			payload:   `{"name": 1}`,
			status:    http.StatusBadRequest,
			buildMock: func(mock *cat.MockCatApp) {},
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
			ctx.Request = httptest.NewRequest("POST", "/api/v1/cats", strings.NewReader(cs.payload))

			// when
			api.Create(ctx)

			var body CreateCatResponse
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
