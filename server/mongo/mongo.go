package mongo

import (
	"context"
	"fmt"
	"server/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mg *mongo.Database

func GetMg() *mongo.Database {
	return mg
}

func init() {
	initMongo()
}

func initMongo() {
	// uri := "mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Get().Mongo.Uri))
	if err != nil {
		panic(err)
	}
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	mg = client.Database(config.Get().Mongo.Database)
	fmt.Println("Successfully connected and pinged.")
}
