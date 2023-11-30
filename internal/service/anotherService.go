package service

import "essemfly/go_base_app/internal/persistence/database"

type AnotherService struct {
	db database.DB
}

func NewAnotherService(db database.DB) *AnotherService {
	return &AnotherService{db: db}
}

func (s *AnotherService) SomeMethod() string {
	return s.db.QuerySomething()
}
