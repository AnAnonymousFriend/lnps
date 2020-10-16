package models

import (
	mongoModel "LNPS/server/models/mongo"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var tableName string = "auth"

type auth struct {
	UserName  string	`bson:"dbname" ,json:"jsonname"`
	PassWord string
}

var authCollection  = NewAuthCollection()

func NewAuthCollection() *mongo.Collection {
	if mongoModel.MongoDB != nil {
		collection  := mongoModel.MongoDB.Collection(tableName)
		return  collection
	}else {
		println("DB 为空链接")
		return nil
	}

}

func Login(userName string,passWord string) (bool,error)  {
	var one auth

	var variableName = auth {userName, passWord}
	err := authCollection.FindOne(context.Background(), variableName).Decode(&one)

	if err !=nil {
		return false, err
	}
	return true,nil
}

func AddUser(userName string,passWord string)  bool{

	ash := auth{userName, passWord,}
	_, err := authCollection.InsertOne(context.TODO(), ash)

	if err != nil {
		println(err)
		return  false
	}
	return true
}