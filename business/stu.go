package business

import (
	"github.com/iwanzhang/tools/global"
	"go.mongodb.org/mongo-driver/mongo"
)

type Stu struct {
	Age  int    `bson:"age"`
	Name string `bson:"name"`
}

func (s *Stu) GetSession() *mongo.Collection {

	return global.MongoClient.Database("test").Collection("stu")

}
