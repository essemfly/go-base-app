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

func (kc *KafkaConsumer) ConsumeMessages(topic string) {
	partitionConsumer, err := kc.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Partition consumer 생성 실패: %s", err.Error())
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		log.Printf("메시지 수신: Key: %s, Value: %s", string(message.Key), string(message.Value))
	}
}
