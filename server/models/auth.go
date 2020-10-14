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

func main()  {
	collection  := mongoModel.MongoDB.Collection(tableName)
	MongoCollection = collection
}

/*func init()  {
	setting.Setup()
	var mongoHost = setting.MongoDBSetting.Host
	var DbName = setting.MongoDBSetting.DbName

	var db,err = mongoModel.ConnectToDB(mongoHost,DbName,10,10)
	if err!=nil {
		println(err)
	}
	MongoDb = db
}*/


func Login(userName string,passWord string) (bool,error)  {
	var one auth
	variable_name := auth {userName, passWord}
	//collection  := MongoDb.Collection(tableName)
	err := MongoCollection.FindOne(context.Background(), variable_name).Decode(&one)

	if err !=nil {
		return false, err
	}
	return true,nil
}

func AddUser(userName string,passWord string)  bool{
	print(userName)
	print(passWord)
	ash := auth{userName, passWord,}
	//collection  := MongoDb.Collection(tableName)
	_, err := MongoCollection.InsertOne(context.TODO(), ash)
	if err != nil {
		println(err)
		return  false
	}
	return true
}