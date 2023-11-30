package service

import (
	"essemfly/go_base_app/internal/persistence"
)

type Services struct {
	MyService      *MyService
	AnotherService *AnotherService
	LogService     *LogService
}

func InitializeServices(p *persistence.Persistences) (*Services, error) {
	myService := NewMyService(p.SQLDatabase)
	anotherService := NewAnotherService(p.SQLDatabase)
	logService := NewLogService(p.KafkaProducer, p.KafkaConsumer, p.RedisClient)

	return &Services{
		MyService:      myService,
		AnotherService: anotherService,
		LogService:     logService,
	}, nil
}
