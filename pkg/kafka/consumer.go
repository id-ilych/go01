package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"

	"go01/pkg/models"
)

func ListenForHotdogs(cfg Config, receiver func(h *models.Hotdog) error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	log.Printf("connecting to broker: %s", cfg.Broker)
	consumer, err := sarama.NewConsumer([]string{cfg.Broker}, config)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Panic(err)
		}
	}()

	log.Printf("starging '%s' topic consumer", cfg.Topic)
	partitionConsumer, err := consumer.ConsumePartition(cfg.Topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Panic(err)
		}
	}()

	log.Printf("ready to receive messages")
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("received message: %s\n", string(msg.Value))

			var hotdog models.Hotdog
			err := json.Unmarshal(msg.Value, &hotdog)
			if err != nil {
				log.Panic(fmt.Errorf("failed to parse hotdog: %w", err))
			}

			if err := receiver(&hotdog); err != nil {
				log.Panic(err)
			}

		case err := <-partitionConsumer.Errors():
			log.Printf("received error: %v\n", err.Err)
		default:
			continue
		}
	}
}
