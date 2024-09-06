package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	//apiParam()
	//urlParam()
	group()
}

// api参数可以通过context的param方法来获取
func apiParam() {
	engine := gin.Default()
	engine.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		//截取action中多出来的/
		action = strings.Trim(action, "/")
		context.String(http.StatusOK, name+" is "+action)
	})
	engine.Run()
}

// url参数可以通过DefaultQuery或Query方法获取
// DefaultQuery若参数不存在，返回默认值，Query若不存在，返回空串
func urlParam() {
	engine := gin.Default()
	engine.GET("/user", func(context *gin.Context) {
		//指定默认值
		name := context.DefaultQuery("name", "lk")
		context.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	engine.Run()
}

// group用了管理一些相同的url
func group() {
	r := gin.Default()
	//路由组1，处理get请求
	v1 := r.Group("/v1")
	//{}是书写规范
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
	r.Run()
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "lk")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "cxylk")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
