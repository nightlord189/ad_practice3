package db

import (
	"context"
	"time"

	"github.com/nightLord189/ad_practice3/config"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Launch(cfg *config.Config) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.ConnectionString))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	//defer client.Disconnect(ctx)
	return client
}
