package handler

import (
	"essemfly/go_base_app/internal/service"

	"github.com/gin-gonic/gin"
)

type PingHandler struct {
	Services *service.Services
}

func (h *PingHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func NewPingHandler(services *service.Services) *PingHandler {
	return &PingHandler{
		Services: services,
	}
}
