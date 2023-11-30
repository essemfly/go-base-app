package persistence

import (
	"essemfly/go_base_app/config"
	"essemfly/go_base_app/internal/persistence/database"
	"essemfly/go_base_app/internal/persistence/kafka"
	"essemfly/go_base_app/internal/persistence/redis"
)

type Persistences struct {
	KafkaProducer *kafka.KafkaProducer
	KafkaConsumer *kafka.KafkaConsumer
	RedisClient   *redis.Redis
	SQLDatabase   *database.SQLDatabase
}

func InitializePersistence(cfg config.Config) (*Persistences, error) {
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

	return &Persistences{
		KafkaProducer: kafkaProducer,
		KafkaConsumer: kafkaConsumer,
		RedisClient:   redisClient,
		SQLDatabase:   db,
	}, nil

}
