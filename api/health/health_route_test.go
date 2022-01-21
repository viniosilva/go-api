package health

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
		body   HealthResponse
	}{
		"should return status ok": {http.StatusOK, HealthResponse{Status: HealthStatusOK}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			// given
			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			api := NewApi()

			// when
			api.Health(ctx)
			var body HealthResponse
			json.Unmarshal(res.Body.Bytes(), &body)

			// then
			if res.Code != cs.status {
				t.Errorf("%s\nresult:\t\t%v\nexpected:\t%v", "StatusCode", res.Code, cs.body)
			}
			if body != cs.body {
				t.Errorf("%s\nresult:\t\t%v\nexpectedBody:\t%v", "Body", body, cs.body)
			}
		})
	}
}
