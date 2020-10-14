package models

import (
	mongoModel "LNPS/server/models/mongo"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var tableName string = "auth"

type auth struct {
	UserName  string	`bson:"dbname",json:"jsonname"`
	PassWord string
}


var MongoCollection *mongo.Collection

func Login(userName string,passWord string) (bool,error)  {
	var one auth
	var variable_name = auth {userName, passWord}
	collection  := mongoModel.MongoDB.Collection(tableName)
	err := collection.FindOne(context.Background(), variable_name).Decode(&one)

	if err !=nil {
		return false, err
	}
	return true,nil
}

func AddUser(userName string,passWord string)  bool{
	print(userName)
	print(passWord)
	ash := auth{userName, passWord,}
	collection  := mongoModel.MongoDB.Collection(tableName)
	_, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		println(err)
		return  false
	}
	return true
}