package service

import (
	"essemfly/go_base_app/internal/domain"
	"essemfly/go_base_app/internal/persistence/postgres"
)

type AnotherService struct {
	Repos *postgres.SQLRepository
}

func NewAnotherService(repos *postgres.SQLRepository) *AnotherService {
	return &AnotherService{Repos: repos}
}

func (s *AnotherService) ListLogs() []*domain.LogData {
	return s.Repos.LogRepository.ListLogs()
}
