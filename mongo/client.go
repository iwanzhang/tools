package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/iwanzhang/tools/global"
)

func InitClient(ip string, port int) {

	uri := fmt.Sprintf("mongodb://%s:%d", ip, port)
	clientOption := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	global.MongoClient = client
}
