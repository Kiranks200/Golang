package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("demoDB").Collection("info_colletions")

	newDocument := bson.D{{"name", "D"}, {"age", 34}, {"city", "egg Island"}}
	insertResult, err := collection.InsertOne(context.TODO(), newDocument)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a document:", insertResult.InsertedID)

	document2 := bson.D{{"name", "N"}, {"age", 40}, {"city", "leaf"}}
	insertResult2, err := collection.InsertOne(context.TODO(), document2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted another document: ", insertResult2.InsertedID)

	document3 := bson.D{{"name", "B"}, {"age", 40}, {"city", "karakura"}}
	res, _ := collection.InsertOne(context.TODO(), document3)
	fmt.Println(res)

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
}
