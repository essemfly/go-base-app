package domain

import "time"

type LogData struct {
	Topic     string
	Message   string
	Score     int
	CreatedAt time.Time
}

type LogAnalysisResult struct {
	Topic     string
	AvgScore  float64
	NumLogs   int
	CreatedAt time.Time
}
