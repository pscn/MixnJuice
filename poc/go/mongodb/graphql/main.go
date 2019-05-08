package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Flavor struct {
	UUID    string
	name    string
	vendor  string
	density float64
}

func Connect(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client
}

func main() {
	client := Connect("mongodb://localhost:27017")
	collection := client.Database("gusta").Collection("flavor")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, Flavor{
		name:    "pi",
		density: 3.14159,
		vendor:  "math",
	})
	if err != nil {
		panic(err)
	}
	id := res.InsertedID
	fmt.Printf("ID=%d\n", id)
}
