package kafka

import (
	"books-management-system/config"
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	Producer *kafka.Producer
}

func NewKafkaProducer() (*Producer, error) {
	kafkaConfig := config.AppConfig.Kafka
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaConfig.Broker,
	})
	if err != nil {
		return nil, err
	}
	return &Producer{Producer: p}, nil
}

func (p *Producer) Publish(topic, eventType string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          jsonData,
		Key:            []byte(eventType),
	}, nil)

	if err != nil {
		return err
	}

	log.Printf("Kafka Event Published: %s -> %s", eventType, string(jsonData))
	return nil
}
