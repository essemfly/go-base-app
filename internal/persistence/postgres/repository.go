package postgres

import "essemfly/go_base_app/internal/domain"

type SQLRepository struct {
	LogRepository          LogRepository
	LogAnalyticsRepository LogAnalyticsRepository
}

type LogRepository interface {
	WriteLog(log *domain.LogData) error
	ListLogs() []*domain.LogData
}

type LogAnalyticsRepository interface {
	ListLogAnalysisResults() []*domain.LogAnalysisResult
}
