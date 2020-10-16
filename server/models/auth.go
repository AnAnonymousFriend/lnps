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


/*func NewMongoCollection() *mongo.Collection  {

	if mongoModel.MongoDB != nil {
		collection  := mongoModel.MongoDB.Collection(tableName)
		return  collection
	}else {
		println("DB 为空链接")
		return nil
	}
}*/

func main()  {
	if mongoModel.MongoDB != nil {
		collection  := mongoModel.MongoDB.Collection(tableName)
		MongoCollection =  collection
	}else {
		println("DB 为空链接")
	}
}



func Login(userName string,passWord string) (bool,error)  {
	var one auth
	var variable_name = auth {userName, passWord}

	if MongoCollection != nil {
		 print("连接状态为：true")
	}else {
		println("连接状态为：False")
	}

	collection  := mongoModel.MongoDB.Collection(tableName)
	err := collection.FindOne(context.Background(), variable_name).Decode(&one)

	if err !=nil {
		return false, err
	}
	return true,nil
}

func AddUser(userName string,passWord string)  bool{
	
	ash := auth{userName, passWord,}
	collection  := mongoModel.MongoDB.Collection(tableName)
	_, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		println(err)
		return  false
	}
	return true
}