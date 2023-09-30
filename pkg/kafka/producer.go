package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"

	"go01/pkg/models"
)

type HotdogProducer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewHotdogProducer(cfg Config) (*HotdogProducer, error) {
	brokers := []string{cfg.Broker}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	p, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return &HotdogProducer{p, cfg.Topic}, nil
}

func (p *HotdogProducer) ProduceHotdog(h *models.Hotdog) (string, error) {
	hotdogJson, err := json.Marshal(h)
	if err != nil {
		return "", fmt.Errorf("error serializing a hotdog: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.StringEncoder(hotdogJson),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return "", fmt.Errorf("error producing a hotdog: %w", err)
	}

	return fmt.Sprintf("message is stored in topic(%s)/partition(%d)/offset(%d)", p.topic, partition, offset), nil
}

func (p *HotdogProducer) Close() error {
	return p.producer.Close()
}
