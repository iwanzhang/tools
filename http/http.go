package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/iwanzhang/tools/business"
	"go.mongodb.org/mongo-driver/bson"
)

func Server() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/insert", insert)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func insert(c *gin.Context) {

	var stu business.Stu
	//
	//stu.Name = "iwanzhang"
	//stu.Age = 25
	//
	//res, err := stu.GetSession().InsertOne(context.TODO(), stu)
	//fmt.Println(res)
	//if err != nil {
	//	panic(err)
	//}

	filter := bson.D{{"name", "iwan"}} //精确搜索

	var result business.Stu
	err := stu.GetSession().FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"message": result,
	})

}
