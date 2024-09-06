package main

import (
	"basic/example/model"
	"basic/example/router"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var addr = flag.String("addr", "127.0.0.1:8080", "server addr")
var seelogConfig = flag.String("log", "example/conf/seelog.xml", "seelog config")
var mysqlPath = flag.String("mysql", "example/conf/mysql.json", "mysql config")

func init() {
	logger, err := seelog.LoggerFromConfigAsFile(*seelogConfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	seelog.ReplaceLogger(logger)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	h2s := &http2.Server{}
	g := gin.Default()
	router.SetRouter(g)
	model.InitDb(*mysqlPath)
	server := &http.Server{
		Handler:        h2c.NewHandler(g, h2s),
		Addr:           *addr,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	seelog.Info("server run:", *addr)
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGHUP)
	seelog.Info("服务启动")
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			seelog.Error(err.Error())
			seelog.Flush()
			os.Exit(0)
		}
	}()
	//退出应用
	<-quitChan
	seelog.Info("服务退出")
	_ = server.Close()
}
