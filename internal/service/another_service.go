package service

import "essemfly/go_base_app/internal/db"

type AnotherService struct {
	db db.Database
}

func NewAnotherService(db db.Database) *AnotherService {
	return &AnotherService{db: db}
}

func (s *AnotherService) SomeMethod() string {
	return s.db.QuerySomething()
}
