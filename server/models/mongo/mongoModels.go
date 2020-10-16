package mongoModel

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var MongoDB = NewMongoDB()


func NewMongoDB() *mongo.Database{
		/*var mongoHost = setting.MongoDBSetting.Host
		var DbName = setting.MongoDBSetting.DbName
		*/
	var mongoHost = "mongodb://localhost:27017"
	var DbName = "lnps"
		var db,err = ConnectToDB(mongoHost,DbName,10,10)
		if err!=nil {
			println(err)
		}
		return db
}

func ConnectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, err
	}

	return client.Database(name), nil
}

