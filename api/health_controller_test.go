package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestApiHealthControllerHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cases := map[string]struct {
		status int
		body   map[string]interface{}
	}{
		"should return status ok": {http.StatusOK, map[string]interface{}{
			"status": "OK",
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			// given
			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			controller := NewHealthController()

			// when
			controller.Health(ctx)
			body := res.Body.String()
			jsonb, _ := json.Marshal(cs.body)
			expectedBody := string(jsonb)

			// then
			if res.Code != cs.status {
				t.Errorf("%s\nresult: %v\nexpected: %v", "StatusCode", res.Code, cs.body)
			}
			if body != expectedBody {
				t.Errorf("%s\nresult: %s\nexpectedBody: %s", "Body", body, expectedBody)
			}
		})
	}
}
