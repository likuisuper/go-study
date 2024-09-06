package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	//test()
	test1()
}

func test() {

	r := gin.Default()
	//使用全局中间件
	//r.Use(MiddleWare())
	{
		//局部中间件，作用于单个路由
		r.GET("/ce", MiddleWare(), func(context *gin.Context) {
			//取值
			value, _ := context.Get("name")
			log.Println("name:", value)
			context.JSON(http.StatusOK, gin.H{"name": value})
		})
	}
	r.Run()
}

func test1() {
	r := gin.Default()
	r.Use(myTime)
	group := r.Group("/shopping")
	{
		group.GET("/index", shopIndexHandler)
		group.GET("/home", shopHomeHandler)
	}
	r.Run()
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}
