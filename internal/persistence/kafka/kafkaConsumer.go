package kafka

import (
	"essemfly/go_base_app/config"
	"log"

	"github.com/IBM/sarama"
)

type KafkaConsumer struct {
	consumer sarama.Consumer
}

func NewKafkaConsumer(cfg config.Config) (*KafkaConsumer, error) {
	brokers := []string{cfg.KafkaURL}
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumer: consumer,
	}, nil
}

func (kc *KafkaConsumer) ConsumeMessages(topic string) []map[string]string {
	partitionConsumer, err := kc.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Partition consumer 생성 실패: %s", err.Error())
	}
	defer partitionConsumer.Close()

	var messages []map[string]string

	msgs := partitionConsumer.Messages()
	for msg := range msgs {
		messages = append(messages, map[string]string{
			"key":   string(msg.Key),
			"value": string(msg.Value),
		})
	}

	return messages
}
