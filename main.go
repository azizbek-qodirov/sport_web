package main

import (
	"log"
	"project/api"
	"project/config"
	"project/kafka"
	"project/storage"
)

func main() {
	cf := config.Load()
	dbs, err := storage.NewStorage(cf)
	if err != nil {
		log.Fatal(err)
	}
	defer dbs.Close()

	kcm := kafka.NewKafkaConsumerManager()
	if err := kcm.RegisterConsumer([]string{cf.KAFKA_BROKER}, "message-create", "message-create-id", kafka.MessageCreateHandler(dbs.Chat)); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            log.Printf("Consumer for topic'message-create' already exists")
        } else {
            log.Fatalf("Error registering 'message-create': %v", err)
        }
	}

	r := api.NewRouter(dbs)
	if err := r.Run(cf.GATEWAY_PORT); err != nil {
		log.Fatal(err)
	}
}
