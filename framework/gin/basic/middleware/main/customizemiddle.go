package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func MiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		log.Println("自定义中间件开始执行")
		//设置变量到context的key中，可以通过get()获取
		context.Set("name", "lk")
		status := context.Writer.Status()
		//执行函数
		context.Next()
		//中间件执行完后续的一些事情
		log.Println("自定义中间件执行完毕", status)
		end := time.Since(start)
		log.Println("执行时间：", end)
	}
}

func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	//统计时间
	end := time.Since(start)
	log.Println("程序用时：", end)
}
