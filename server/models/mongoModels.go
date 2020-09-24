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

var MongoDb *mongo.Database

func MongoInit() {
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
		MongoDb = client.Database(dbName)
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
}

func InsertOne(tableName string,obj interface{})  (bool,error){
	collection  := MongoDb.Collection(tableName)
	_, err := collection.InsertOne(context.TODO(), obj)
	if err != nil {
		return true, err
	}
	println("添加条数据失败:",err)
	return false,err
}

func DeleteOne(tableName string,obj interface{}) (bool,error) {
	collection  := MongoDb.Collection(tableName)
	_, err := collection.DeleteOne(context.TODO(), obj)
	if err != nil {
		return true, err
	}
	println("添加条数据失败:",err)
	return false,err

}

func DeleteAll(tableName string) {
	collection  := MongoDb.Collection(tableName)
	collection.Drop(context.Background())
}