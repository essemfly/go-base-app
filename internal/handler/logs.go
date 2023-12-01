package handler

import (
	"encoding/json"
	"essemfly/go_base_app/internal/domain"
	"essemfly/go_base_app/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogHandler struct {
	Services *service.Services
}

func (h *LogHandler) PostLogs(c *gin.Context) {
	var logData domain.LogData
	if err := c.BindJSON(&logData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(logData)
	if err != nil {
		log.Fatalf("JSON 직렬화 실패: %s", err)
		return
	}

	if err := h.Services.LogService.CreateLog(string(jsonData)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "로그 처리 완료"})
}

func (h *LogHandler) GetAnalytics(c *gin.Context) {
	result, err := h.Services.LogService.GetLatestLogAnalytics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func (h *LogHandler) ListLogsIntoPostgres(c *gin.Context) {
	logs := h.Services.AnotherService.ListLogs()
	c.JSON(200, logs)
}

func (h *LogHandler) PostLogsIntoPostgres(c *gin.Context) {
	h.Services.MyService.WriteLog()
	c.JSON(http.StatusOK, gin.H{"message": "로그 처리 완료"})
}

func NewLogsHandler(services *service.Services) *LogHandler {
	return &LogHandler{
		Services: services,
	}
}
