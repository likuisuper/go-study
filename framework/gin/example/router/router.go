package router

import (
	"basic/example/controller"
	_ "basic/example/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRouter(r *gin.Engine) {
	//设置跨域中间件
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	person := &controller.Person{}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("/v1")
	{
		v1.GET("/person/:id", person.GetById)
		v1.GET("/person/", person.GetAll)
		v1.POST("/person", person.Add)
		v1.PUT("/person/:id", person.Update)
		v1.DELETE("/person/:id", person.Del)
	}
}
