package db

import (
	"context"
	"film36exp/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

const (
	Database       = "film36exp"
	CollectionFilm = "films"
	CollectionPic  = "Pics"
)

func SetClint(c *mongo.Client) {
	// can be more explicitly
	if client == nil {
		client = c // <--- NOT THREAD SAFE
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	return client.Database(Database).Collection(collectionName)
}

func Create(collectionName string, item interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(Database).Collection(collectionName)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	return collection.InsertOne(ctx, item)
}

func Delete(collectionName string, _id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": bson.M{"$eq": _id}}
	collection := client.Database(Database).Collection(collectionName)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	return collection.DeleteOne(ctx, filter)
}

func Update(collectionName string, _id primitive.ObjectID, item interface{}) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": bson.M{"$eq": _id}}
	update := bson.M{"$set": item}
	collection := client.Database(Database).Collection(collectionName)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	return collection.UpdateMany(ctx, filter, update)
}

func FindOne(collectionName string, _id primitive.ObjectID) (r *mongo.SingleResult) {
	collection := client.Database(Database).Collection(collectionName)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	return collection.FindOne(ctx, model.Film{ID: _id})
}
