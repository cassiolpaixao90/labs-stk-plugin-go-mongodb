package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

var (
	mongoClient *mongo.Client
	once        sync.Once
)

type MongoDBConfig struct {
	URI    string
	Logger interface{}
}

func NewMongoConnection(cfg *MongoDBConfig) *MongoDBConfig {
	return &MongoDBConfig{URI: cfg.URI, Logger: cfg.Logger}
}

func (m *MongoDBConfig) GetMongoClient() *mongo.Client {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientOptions := options.Client().ApplyURI(m.URI)

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			return
		}

		mongoClient = client
	})

	return mongoClient
}
