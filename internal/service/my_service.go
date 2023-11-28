package service

import "essemfly/go_base_app/internal/db"

type MyService struct {
	db db.Database
}

func NewMyService(db db.Database) *MyService {
	return &MyService{db: db}
}

func (s *MyService) SomeMethod() string {
	return s.db.QuerySomething()
}
