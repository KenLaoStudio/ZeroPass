package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
		log.Printf(err.Error())
		return nil, err
	}

	return result, nil
}

func (mc *MongoClient) FindOne(database, collection string, filter bson.M, result interface{}) error {
	coll := mc.Client.Database(database).Collection(collection)
	err := coll.FindOne(context.Background(), filter).Decode(result)
	return err
}

func (mc *MongoClient) FindAll(database, collection string, filter bson.M, results interface{}) error {
	coll := mc.Client.Database(database).Collection(collection)
	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		return err
	}
	err = cursor.All(context.Background(), results)
	return err
}

func (mc *MongoClient) UpdateOne(database, collection string, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	coll := mc.Client.Database(database).Collection(collection)
	result, err := coll.UpdateOne(context.Background(), filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	return result, nil
}
