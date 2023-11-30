package service

import (
	"essemfly/go_base_app/config"
	"essemfly/go_base_app/internal/persistence/database"
	"essemfly/go_base_app/internal/persistence/kafka"
	"essemfly/go_base_app/internal/persistence/redis"
)

type Services struct {
	MyService      *MyService
	AnotherService *AnotherService
	LogService     *LogService
}

func InitializeServices(cfg config.Config) (*Services, error) {
	db, err := database.NewSQLDatabase(cfg)
	if err != nil {
		return nil, err
	}

	kafkaProducer, err := kafka.NewKafkaProducer(cfg)
	if err != nil {
		return nil, err
	}

	kafkaConsumer, err := kafka.NewKafkaConsumer(cfg)
	if err != nil {
		return nil, err
	}

	redisClient, err := redis.NewRedisClient(cfg)
	if err != nil {
		return nil, err
	}

	myService := NewMyService(db)
	anotherService := NewAnotherService(db)
	logService := NewLogService(kafkaProducer, kafkaConsumer, redisClient)

	return &Services{
		MyService:      myService,
		AnotherService: anotherService,
		LogService:     logService,
	}, nil
}
