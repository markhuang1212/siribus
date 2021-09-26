package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markhuang1212/siribus/internal/router"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	r := router.NewRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())

}
