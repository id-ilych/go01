package main

import (
	"log"

	"go01/pkg/http"
	"go01/pkg/kafka"
	"go01/pkg/models"
)

func main() {
	producer, err := kafka.NewHotdogProducer(kafka.Config{Broker: "localhost:9092", Topic: "hotdogs"})
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Panic(err)
		}
	}()

	http.ListenForHotdogs(
		http.Config{Addr: ":8080", Route: "/hotdogs"},
		func(h *models.Hotdog) bool {
			log.Printf("received: %+v\n", h)
			res, err := producer.ProduceHotdog(h)
			if err != nil {
				log.Printf("%s\n", err)
				return false
			}

			log.Println(res)
			return true
		},
	)
}
