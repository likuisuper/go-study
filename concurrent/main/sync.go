package main

import (
	"fmt"
	"image"
	"strconv"
	"sync"
)

/** sync包常用组件
1、waitGroup
	waitGroup类似于java中的countdownLatch，add等同于构造函数传递的线程（协程），Done等同于countDown，Wait等同于await
2、Once
	确保某些操作在高并发场景下只执行一次，例如只加载一次配置文件，只关闭一次通道。
	它内部包含了一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，布尔值用来记录初始化是否完成。具体可查看once结构体
3、Map
	go中的map不是并发安全的，要使用并发安全的map，需要使用sync.Map,并且该Map不用初始化可直接使用
*/
var wg sync.WaitGroup
var icons map[string]image.Image
var m = make(map[string]int)
var m2 = sync.Map{}

func wgTest() {
	defer wg.Done()

	fmt.Println("Hello WaitGroup")
}

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main11() {
	//waitGroup测试
	//wg.Add(1)
	//go wgTest()
	//fmt.Println("main goroutine done!")
	//wg.Wait()

	//并发修改普通map测试，当并发多了之后会报错：concurrent map writes
	//wg2 := sync.WaitGroup{}
	//for i := 0; i < 20; i++ {
	//	wg2.Add(1)
	//	go func(n int) {
	//		key := strconv.Itoa(n)
	//		set(key, n)
	//		fmt.Printf("k=:%v,v=:%v\n", key, get(key))
	//		wg2.Done()
	//	}(i)
	//}
	//wg2.Wait()

	wg2 := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg2.Add(20)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)
			value, _ := m2.Load(key)
			fmt.Printf("k=:%v,v=:%v\n", key, value)
			wg2.Done()
		}(i)
	}
	wg2.Wait()
}
