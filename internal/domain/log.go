package domain

import "time"

type Log struct {
	Topic     string
	Message   string
	Score     int
	CreatedAt time.Time
}
