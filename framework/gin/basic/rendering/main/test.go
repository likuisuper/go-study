package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"log"
	"net/http"
	"time"
)

func main() {
	//resp()
	//redirect()
	async()
}

// 各种数据格式的响应
func resp() {
	r := gin.Default()
	//gin.H 是map[string]interface{}的一种快捷方式
	//json
	r.GET("/someJSON", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	//结构体响应
	r.GET("/moreJSON", func(context *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "lk"
		msg.Message = "hey"
		msg.Number = 123
		//注意msg.Name在json中变成了user
		context.JSON(http.StatusOK, msg)
	})

	//3.XML
	r.GET("someXML", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	//4.yaml
	r.GET("/someYMAL", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	//5.protobuf
	r.GET("/someProtoBuf", func(context *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		//protobuf的具体定义在testdata/protoexample/test.pb.go中
		//这里要使用指针接收，因为Test中包含了state，它是MessageState的，而这个结构体有DoNotCopy字段，它持有互斥锁
		//如果不使用指针的话，那么得到的就是对结构体的副本，这里会造成协程持有不同的锁，具体看：https://stackoverflow.com/questions/64183794/why-do-the-go-generated-protobuf-files-contain-mutex-locks
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		//数据在响应中变成了二进制数据
		//将输出protoexample.Test protobuf序列化了的数据

		context.ProtoBuf(http.StatusOK, data)
	})

	r.Run(":8080")
}

// 重定向
func redirect() {
	r := gin.Default()
	r.GET("/index", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://wwww.baidu.com")
	})
	r.Run()
}

// 同步、异步
// 当启动新的goroutine时，不能使用原始上下文，必须使用只读副本，不然会造成线程安全问题
func async() {
	r := gin.Default()
	//1、异步
	r.GET("/long_async", func(context *gin.Context) {
		//创建在goroutine中使用的副本
		cCp := context.Copy()
		go func() {
			//模拟长任务
			time.Sleep(5 * time.Second)
			//注意使用context的副本
			log.Println("异步执行：" + cCp.Request.URL.Path)
		}()
	})
	//2、同步
	r.GET("/long_sync", func(context *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + context.Request.URL.Path)
	})
	r.Run()
}
