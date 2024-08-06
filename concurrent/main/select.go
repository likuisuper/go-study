package main

import (
	"fmt"
	"time"
)

/*
	多路复用，监听多个管道，直到其中一个channel ready。没有数据则阻塞
*/

func main10() {
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//for {
	//	data, ok := <-ch1
	//	data2, ok2 := <-ch2
	//}
	//使用select同时响应多个通道操作代替上面的循环
	//select类似于switch语句
	//select {
	//case <-ch1:
	//	fmt.Println("如果ch1成功读取到数据，则进行该case处理语句")
	//case <-ch2:
	//	fmt.Println("如果ch2成功读取到数据，则进行该case处理语句")
	//default:
	//	//如果上面都没有成功，则进入default处理数据
	//}

	output1 := make(chan string)
	output2 := make(chan string)
	go test1(output1)
	go test2(output2)
	//select监控，有一个管道ready就是执行对象的case语句，比如这里只会执行s2
	//当把test1和test2中的睡眠时间去掉，则两个管道都ready的情况下，会随机选择一个执行
	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
}

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}
