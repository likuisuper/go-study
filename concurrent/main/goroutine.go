package main

import (
	"fmt"
	"time"
)

/*
	goroutine类似于线程，但它是更轻量级的线程，即协程。
	一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数
*/

func hello() {
	fmt.Println("hello goroutine")
}

func main1() {
	//启动另外一个goroutine去执行hello函数
	//但是hello函数并不会输出。
	//程序启动时，go程序就会为main函数创建一个默认的goroutine
	//当main函数返回的时候goroutine就结束了，所有在main函数中启动的goroutine会一同结束
	go hello()
	fmt.Println("main goroutine done!")
	//阻塞main goroutine，等待hello goroutine结束
	//会先输出main goroutine done!，因为在创建新的goroutine的时候需要花费一些时间，而此时main goroutine是继续执行的
	time.Sleep(time.Second)
}
