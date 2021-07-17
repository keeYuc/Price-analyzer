package mongo

import (
	"context"
	"fmt"
	"server/config"

	"go.mongodb.org/mongo-driver/bson"
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

const (
	main  = "main"
	day_1 = "day_1"
	day_2 = "day_2"
	day_3 = "day_3"
	m60_1 = "day_1"
	m60_2 = "day_2"
	m60_3 = "day_3"
	m60_4 = "day_4"
	m60_5 = "day_5"
	m60_6 = "day_6"
)

func createIndex(m *mongo.Database) {
	c_main := mg.Collection(main)
	c_main.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{"id", 1}},
			Options: options.Index().SetUnique(true).SetName("id"),
		},
	})
	c_day1 := mg.Collection(day_1)
	c_day1.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{"id", 1}},
			Options: options.Index().SetName("id"),
		},
	})
	c_day2 := mg.Collection(day_2)
	c_day2.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{"id", 1}},
			Options: options.Index().SetName("id"),
		},
	})
	c_day3 := mg.Collection(day_3)
	c_day3.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{"id", 1}},
			Options: options.Index().SetName("id"),
		},
	})
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
	createIndex(mg)
	fmt.Println("Successfully connected and pinged.")
}
