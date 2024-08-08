package kafka

import (
	"encoding/json"
	"log"
	"project/models"
	"project/storage/managers"
	// "google.golang.org/protobuf/encoding/protojson"
)

func MessageCreateHandler(chatRepo *managers.ChatManager) func(message []byte) {
	return func(message []byte) {
		var req models.MessageCReq
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		err := chatRepo.CreateMessage(&req)
		if err != nil {
			log.Printf("Cannot create message via Kafka: %v", err)
			return
		}
		log.Printf("Created message for chat with id: %+v",req.MatchID)
	}
}
