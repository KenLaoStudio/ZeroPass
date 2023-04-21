package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client *mongo.Client
}

func NewMongoClient(connectionString string) (*MongoClient, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return &MongoClient{
		Client: client,
	}, nil
}

func (mc *MongoClient) InsertDocument(database, collection string, document interface{}) (*mongo.InsertOneResult, error) {
	coll := mc.Client.Database(database).Collection(collection)
	result, err := coll.InsertOne(context.Background(), document)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// 類似地，實現其他 CRUD 操作
