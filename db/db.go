package db

import "go.mongodb.org/mongo-driver/mongo"

type MongoDB struct {
	Collection *mongo.Collection
}

var Client *mongo.Client

func Films() *mongo.Collection {
	return Client.Database("film36exp").Collection("films")
}
