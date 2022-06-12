package main

import (
	"fmt"
	"github.com/iwanzhang/tools/global"
	"github.com/iwanzhang/tools/mongo"
)

func main() {

	//localhost 27017
	mongo.InitClient("10.211.55.12", 27017)
	fmt.Printf("+%v", global.MongoClient)
}
