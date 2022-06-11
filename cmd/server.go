package main

import (
	"github.com/iwanzhang/tools/http"
	"github.com/iwanzhang/tools/mongo"
)

func main() {

	mongo.InitClient("localhost", 27017)
	http.Server()
}
