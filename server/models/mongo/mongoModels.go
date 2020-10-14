package mongoModel

import (
	"LNPS/server/pkg/setting"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)


var MongoDB *mongo.Database

func Setup(){

		var mongoHost = setting.MongoDBSetting.Host
		print("参数为")
		var DbName = setting.MongoDBSetting.DbName

		var db,err = ConnectToDB(mongoHost,DbName,10,10)
		if err!=nil {
			println(err)
		}
		MongoDB = db
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

