package models

import (
	"LNPS/server/models/mongo"
	"LNPS/server/pkg/setting"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

var tableName string = "auth"

type auth struct {
	UserName  string	`bson:"dbname",json:"jsonname"`
	PassWord string
}


var MongoDb *mongo.Database

func init()  {
	fmt.Printf("%v \n", setting.DatabaseSetting.Type)
	println("参数为："+setting.DatabaseSetting.Type)
	var db,err = mongoModel.ConnectToDB("mongodb://localhost:27017","lnps",10,10)
	if err!=nil {
		println(err)
	}
	MongoDb = db
	println("调用Init")

}



func Login(userName string,passWord string) (bool,error)  {
	var one auth
	variable_name := auth {userName, passWord}
	collection  := MongoDb.Collection(tableName)
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
	collection  := MongoDb.Collection(tableName)
	_, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		println(err)
		return  false
	}
	return true
}