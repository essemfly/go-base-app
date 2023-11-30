package service

import (
	"essemfly/go_base_app/internal/persistence/kafka"
	"essemfly/go_base_app/internal/persistence/redis"
)

const LOG_TOPIC_NAME = "logs"

type LogService struct {
	producer    *kafka.KafkaProducer
	consumer    *kafka.KafkaConsumer
	redisCLient *redis.Redis
}

func NewLogService(producer *kafka.KafkaProducer, consumer *kafka.KafkaConsumer, redisClient *redis.Redis) *LogService {
	return &LogService{
		producer:    producer,
		consumer:    consumer,
		redisCLient: redisClient,
	}
}

func (ls *LogService) CreateLog(log string) error {
	return ls.producer.SendMessage(LOG_TOPIC_NAME, log)
}
