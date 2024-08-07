package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"project/config"
	"project/storage/managers"
	"time"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	MongoClient *mongo.Client
	PgsqlClient *sql.DB
	Chat        *managers.ChatManager
	Match       *managers.MatchManager
}

var ctx = context.Background()

func NewStorage(cf *config.Config) (*Storage, error) {
	// ################# PGSQL CONNECTION #################
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cf.DB_HOST, cf.DB_USER, cf.DB_NAME, cf.DB_PASSWORD, cf.DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to the database postgres!!!")

	// ################# MONGODB CONNECTION #################
	clientOptions := options.Client().ApplyURI(cf.MONGO_URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to the database mongodb!!!")

	return &Storage{
		MongoClient: client,
		Chat:        managers.NewChatManager(client, cf.MONGO_DB_NAME, cf.MONGO_CHAT_COLLECTION_NAME),
		Match:       managers.NewMatchManager(db),
	}, nil
}

func (s *Storage) Close() error {
	var errs []error

	if s.MongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := s.MongoClient.Disconnect(ctx); err != nil {
			errs = append(errs, fmt.Errorf("error disconnecting MongoDB: %w", err))
		}
	}

	if s.PgsqlClient != nil {
		if err := s.PgsqlClient.Close(); err != nil {
			errs = append(errs, fmt.Errorf("error closing PostgreSQL: %w", err))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
