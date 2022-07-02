package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	//localhost 27017
	//mongo.InitClient("10.211.55.12", 27017)
	//fmt.Printf("+%v", global.MongoClient)

	query := []bson.M{
		{"$match": bson.M{
			"company": "xx",
			"biz":     "yy",
		}},
	}
	//if req.SearchQuery != "" {
	query[0]["$match"].(bson.M)["$or"] = []bson.M{
		{"code": bson.RegEx{Pattern: fmt.Sprintf(".*%s.*", "req.SearchQuery"), Options: "im"}},
		{"raw_code": bson.RegEx{Pattern: fmt.Sprintf(".*%s.*", "req.SearchQuery"), Options: "im"}},
		{"name": bson.RegEx{Pattern: fmt.Sprintf(".*%s.*", "req.SearchQuery"), Options: "im"}},
	}
	//}

	// 统计查询
	countQ := query
	countQ = append(countQ, bson.M{"$group": bson.M{"_id": "code", "count": bson.M{"$sum": 1}}})

	fmt.Printf("%v", countQ)
}
