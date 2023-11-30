package handler

import (
	"essemfly/go_base_app/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, services *service.Services) {
	appHandler := NewPingHandler(services)

	r.GET("/ping", appHandler.Ping)

	return
}
