package handler

import (
	"essemfly/go_base_app/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, services *service.Services) {
	appHandler := NewPingHandler(services)
	logsHandler := NewLogsHandler(services)

	r.GET("/ping", appHandler.Ping)
	r.POST("/logs", logsHandler.PostLogs)
	r.GET("/logs/analytics", logsHandler.GetAnalytics)

	return
}
