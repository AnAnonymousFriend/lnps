package models

import (
	"context"
	"time"
	"fmt"
	"LNPS/server/pkg/setting"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoCollection *mongo.Collection


func init(tableName string) {
	uri := setting.MongoSetting.Host
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		dbName := setting.MongoSetting.DbName
		MongoCollection = client.Database(dbName).Collection(tableName)
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

}

func Add(tableName string)  {

}