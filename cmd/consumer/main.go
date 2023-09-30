package main

import (
	"log"

	"go01/pkg/database"
	"go01/pkg/kafka"
	"go01/pkg/models"
)

func main() {
	db, err := database.OpenHotdogDatabase(database.Config{Filename: "./hotdogs.db", Table: "hotgogs"})
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Panic(err)
		}
	}()

	kafka.ListenForHotdogs(
		kafka.Config{Broker: "localhost:9092", Topic: "hotdogs"},
		func(h *models.Hotdog) error {
			log.Printf("received: %+v\n", h)
			if err := db.SaveHotdog(h); err != nil {
				log.Printf("failed to store in database: %s\n", err)
				return err
			}
			log.Printf("stored in database: %+v\n", h)
			return nil
		},
	)
}
