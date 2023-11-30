package service

import (
	"encoding/json"
	"essemfly/go_base_app/internal/domain"
	"essemfly/go_base_app/internal/persistence/kafka"
	"essemfly/go_base_app/internal/persistence/redis"
	"log"
	"time"
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

func (ls *LogService) ConsumeLogs() *domain.LogAnalysisResult {
	msgs := ls.consumer.ConsumeMessages(LOG_TOPIC_NAME)

	totalScore := 0
	for _, msg := range msgs {
		var logData domain.LogData
		err := json.Unmarshal([]byte(msg["value"]), &logData)
		if err != nil {
			log.Println("err", err)
			continue
		}
		totalScore += logData.Score
	}

	return &domain.LogAnalysisResult{
		Topic:     LOG_TOPIC_NAME,
		AvgScore:  float64(totalScore) / float64(len(msgs)),
		NumLogs:   len(msgs),
		CreatedAt: time.Now(),
	}
}

func (ls *LogService) GetLatestLogAnalytics() (*domain.LogAnalysisResult, error) {
	results, err := ls.redisCLient.GetLatestLogResult()
	if err != nil {
		return nil, err
	}

	var logResult domain.LogAnalysisResult
	err = json.Unmarshal(results, &logResult)
	if err != nil {
		return nil, err
	}

	return &logResult, nil
}

func (ls *LogService) SaveLogAnalytics(logResult *domain.LogAnalysisResult) error {
	jsonData, err := json.Marshal(logResult)
	if err != nil {
		return err
	}
	score := float64(logResult.CreatedAt.Unix())

	return ls.redisCLient.SaveLogAnalytics(score, jsonData)
}
