package service

import (
	"essemfly/go_base_app/config"
	"essemfly/go_base_app/internal/db"
)

type Services struct {
	MyService      *MyService
	AnotherService *AnotherService
}

func InitializeServices(cfg config.Config) (*Services, error) {
	database, err := db.NewSQLDatabase(cfg)
	if err != nil {
		return nil, err
	}

	myService := NewMyService(database)
	anotherService := NewAnotherService(database)

	return &Services{
		MyService:      myService,
		AnotherService: anotherService,
	}, nil
}
