package handler_test

import (
	"essemfly/go_base_app/internal/handler"
	"essemfly/go_base_app/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/v3/assert"
)

func TestPing(t *testing.T) {
	r := gin.Default()
	handler.SetupRoutes(r, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}

func TestNewPingHandler(t *testing.T) {
	services := &service.Services{}

	handler := handler.NewPingHandler(services)
	assert.Assert(t, handler != nil)
}
