package service

import "essemfly/go_base_app/internal/persistence/database"

type MyService struct {
	db database.DB
}

func NewMyService(db database.DB) *MyService {
	return &MyService{db: db}
}

func (s *MyService) SomeMethod() string {
	return s.db.QuerySomething()
}
