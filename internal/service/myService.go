package service

import (
	"essemfly/go_base_app/internal/domain"
	"essemfly/go_base_app/internal/persistence/postgres"
	"fmt"
	"math/rand"
	"time"
)

type MyService struct {
	Repos *postgres.SQLRepository
}

func NewMyService(repos *postgres.SQLRepository) *MyService {
	return &MyService{Repos: repos}
}

func (s *MyService) WriteLog() {
	randomLogData := generateRandomLogData()

	s.Repos.LogRepository.WriteLog(&randomLogData)
}

func generateRandomLogData() domain.LogData {
	// 무작위 LogData 생성
	return domain.LogData{
		Topic:     "logResults",
		Message:   fmt.Sprintf("Random Message %d", rand.Intn(1000)),
		Score:     rand.Intn(100),
		CreatedAt: time.Now(),
	}
}
