package managers

import (
	"context"
	"project/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type ChatManager struct {
	collection *mongo.Collection
}

func NewChatManager(client *mongo.Client, dbName, collectionName string) *ChatManager {
	collection := client.Database(dbName).Collection(collectionName)
	return &ChatManager{collection: collection}
}

func (m *ChatManager) CreateMessage(message *models.MessageCReq) error {
	message.SentAt = time.Now()
	_, err := m.collection.InsertOne(ctx, message)
	return err
}

func (m *ChatManager) GetMessages(matchID string) (*models.MessageGARes, error) {
	filter := bson.M{
		"match_id": matchID,
	}
	opts := options.Find().SetSort(bson.D{{Key: "sent_at", Value: 1}}) // Sort by SentAt in ascending order

	cursor, err := m.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages models.MessageGARes
	for cursor.Next(ctx) {
		var message models.MessageCReq
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages.Messages = append(messages.Messages, &message)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &messages, nil
}

