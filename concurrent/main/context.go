package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	context具备手动、定时、超时发出取消信号、传值等功能，主要用于控制多个协程之间的协作，尤其是取消操作，具体查看Context接口。
	从使用功能上看，有4种实现好的context
	1、空context：不可取消，没有截止时间，主要用于context树的根节点
	2、可取消的context：用于发出取消信号，当取消的时候，它的子context也会取消
	3、可定时取消的context：多了一个定时的功能
	4、可超时取消的context：多个一个超时功能
	5、值context：用于存储一个key-value键值对
*/

//普通版监控狗，没办法让协程提前退出
func watchDog(name string) {
	//for select循环，一直监控
	for {
		select {
		default:
			fmt.Println(name, "正在监控...")
		}
		time.Sleep(1 * time.Second)
	}
}

//select+channel实现协程退出
//这种方式的缺点就是如果要同时取消多个协程或者定时取消多个协程，代码就会变得复杂难以维护
func watchDog2(name string, stopch chan bool) {
	//for select循环，一直监控
	for {
		select {
		case <-stopch:
			fmt.Println(name, "停止指令已收到，马上停止")
			//go的case语句自带break，所以这里加break是没有办法跳出整个for循环的
			return
		default:
			fmt.Println(name, "正在监控...")
		}
		time.Sleep(1 * time.Second)
	}
}

//使用context实现多个协程的取消
func watchDog3(name string, ctx context.Context) {
	//for select循环，一直监控
	for {
		select {
		//ctx.Done判断是否停止
		case <-ctx.Done():
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控...")
		}
		time.Sleep(1 * time.Second)
	}
}

//context传值，将context存储的值供其他协程使用
func getUser(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("【获取用户】", "协程退出")
			return
		default:
			userId := ctx.Value("userId")
			fmt.Println("【获取用户】", "用户id为：", userId)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	//var wg2 sync.WaitGroup
	//wg2.Add(1)
	//普通版测试
	//go func() {
	//	defer wg2.Done()
	//	watchDog("【监控狗1】")
	//}()
	//wg2.Wait()

	//select+channel测试
	//stopch := make(chan bool)
	//go func() {
	//	defer wg2.Done()
	//	watchDog2("【监控狗1】", stopch)
	//}()
	////先让监控狗监控5秒
	//time.Sleep(5 * time.Second)
	//stopch <- true
	//wg2.Wait()

	//context测试
	//context.WithCancel用于生成一个可以取消的Context，用于发送停止指令，
	//context.Background()用于生成一个空的context，一般作为整个context树的根节点
	//ctx, stop := context.WithCancel(context.Background())
	//go func() {
	//	defer wg2.Done()
	//	watchDog3("【监控狗1】", ctx)
	//}()
	////先让监控狗监控5秒
	//time.Sleep(5 * time.Second)
	////发出停止指令，stop()是WithCancel返回的取消函数
	//stop()
	//wg2.Wait()

	//取消多个协程
	wg3 := sync.WaitGroup{}
	wg3.Add(4)
	ctx, stop := context.WithCancel(context.Background())
	//withValue中的parent一定要传ctx，而不是context.Background()，否则stop()不会生效，getuser不会退出
	valCtx := context.WithValue(ctx, "userId", 2)
	go func() {
		defer wg3.Done()
		watchDog3("【监控狗1】", ctx)
	}()
	go func() {
		defer wg3.Done()
		watchDog3("【监控狗2】", ctx)
	}()
	go func() {
		defer wg3.Done()
		watchDog3("【监控狗3】", ctx)
	}()
	//context传值
	go func() {
		defer wg3.Done()
		getUser(valCtx)
	}()
	time.Sleep(5 * time.Second)
	stop()
	wg3.Wait()

}
