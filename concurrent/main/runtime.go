package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	os线程和goroutine的关系：
		1、一个os线程对应用户态多个goroutine
		2、go程序可以同时使用多个os线程
		3、goroutine和os线程是多对多的关系，即m:n，即调度m个goroutine到n个os线程
*/

func main2() {
	//1、runtime.GoSched:让出CPU时间片，重新等待安排任务
	//go func(s string) {
	//	for i := 0; i < 2; i++ {
	//		fmt.Println(s)
	//	}
	//}("word")
	//
	////主协程
	//for i := 0; i < 2; i++ {
	//	//切一下，再次分配任务
	//	runtime.Gosched()
	//	fmt.Println("hello")
	//}

	//2、runtime.Goexit
	//输出结果为 B.defer A.defer，Goexit之后的defer语句和代码都不会执行
	//go func() {
	//	defer fmt.Println("A.defer")
	//	func() {
	//		defer fmt.Println("B.defer")
	//		//结束协程
	//		runtime.Goexit()
	//		defer fmt.Println("C.defer")
	//		fmt.Println("B")
	//	}()
	//	fmt.Println("A")
	//}()
	//
	//for {
	//
	//}

	//3、runtime.GOMAXPROCS:设置当前程序并发时占用的CPU逻辑核心数，1.5之后默认使用全部的CPU逻辑核心数
	//设置1个核心，此时是做完一个任务再做另外一个任务，设置为2，则两个任务并行执行
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Minute)
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
