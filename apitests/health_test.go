package apitests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"gosky/app/http/routers"
)

//检测是否健康
func TestHealthy(t *testing.T) {
	router := routers.NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	req.Header.Add("guid", "23i2393")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "ok", w.Body.String())
}
