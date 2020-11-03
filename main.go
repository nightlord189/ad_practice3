package main

import (
	"context"
	"fmt"
	"github.com/nightLord189/ad_practice3/config"
	"github.com/nightLord189/ad_practice3/db"
	"github.com/nightLord189/ad_practice3/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("start")
	conf := config.Load("config.json")
	client := db.Launch(conf)

	fmt.Println("dropping collection")
	client.Database("ad_practice3").Collection("almoblinvest").Drop(context.TODO())

	fmt.Println("inserting docs")
	invests := model.GetInvests()
	collection := client.Database("ad_practice3").Collection("almoblinvest")
	for i := 0; i < len(invests); i++ {
		_, err := collection.InsertOne(context.TODO(), invests[i])
		if err != nil {
			panic("Insert err: " + err.Error())
		}
	}

	fmt.Println("creating index")
	indexName, err := client.Database("ad_practice3").Collection("almoblinvest").Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"id": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		panic("err index: " + err.Error())
	}
	fmt.Println(indexName)
}
