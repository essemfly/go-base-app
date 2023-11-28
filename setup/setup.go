package setup

import (
	"essemfly/go_base_app/config"
	"essemfly/go_base_app/internal/db"
	"essemfly/go_base_app/internal/service"
)

type Services struct {
	MyService      *service.MyService
	AnotherService *service.AnotherService
}

func InitializeServices(cfg config.Config) (*Services, error) {
	database, err := db.NewSQLDatabase(cfg)
	if err != nil {
		return nil, err
	}

	myService := service.NewMyService(database)
	anotherService := service.NewAnotherService(database)

	return &Services{
		MyService:      myService,
		AnotherService: anotherService,
	}, nil
}
