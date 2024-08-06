package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	互斥锁
		能够保证同时只有一个goroutine访问共享资源。多个goroutine同时等待一个锁时，唤醒的策略是随机的。
		需要注意，互斥锁虽然可以保证临界区中代码的串行执行，但却不能保证这些代码执行的原子性，因为goroutine
		中断的时间有很多，即使这些语句在临界区之内
	读写锁
		读锁和读锁不互斥，读锁和写锁互斥，写锁和读锁、写锁都互斥
	条件变量
		用于协调想要访问共享资源的那些线程，条件变量是基于锁实现的，所以使用条件变量之前必须获取对应的锁，
		只有先锁定条件变量基于的互斥锁，才能调用wait方法，为什么呢？看一下wait做了什么就知道了：
			1、把调用它的goroutine(也就是当前的goroutine)加入到当前条件变量的通知队列中
			2、解锁当前的条件变量基于的那个互斥锁，也就是释放锁资源，所以在调用wait之前必须先锁定互斥锁，否则引发panic
			3、goroutine处于等待，阻塞在调用wait方法上，直到被通知唤醒
			4、如果被唤醒，那么唤醒之后重新锁定当前条件变量基于的互斥锁
*/

var x int64
var wg1 sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func add() {
	for i := 0; i < 5000; i++ {
		//加锁,重复加锁会造成死锁
		//将加锁和解锁放在循环内部，那么每次循环迭代时，只有一个goroutine尝试获取锁，
		//而在外部时，所有goroutine都会在同一时刻尝试获取锁，这样可以减少竞争，提升性能
		lock.Lock()
		x = x + 1
		//解锁
		lock.Unlock()
		//重复解锁导致panic
		//lock.Unlock()
	}
	wg1.Done()
}

//写锁
func write() {
	rwlock.Lock()
	x = x + 1
	//模拟写操作耗时
	time.Sleep(5 * time.Millisecond)
	rwlock.Unlock()
	wg1.Done()
}

//读锁
func read() {
	rwlock.RLock()
	time.Sleep(2 * time.Millisecond)
	println(x)
	rwlock.RUnlock()
	wg1.Done()
}

//条件变量，模拟向信箱发信
var mailbox uint8
var lock2 sync.RWMutex

//Mutex和RWMutex都拥有lock和unlock方法，只不过它们都是指针方法，所以这两个类型的指针类型才是sync.Locker接口的实现类型
var sendCond = sync.NewCond(&lock2)
var recvCond = sync.NewCond(lock2.RLocker())

func send() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		//这里的lock是持有锁
		lock2.Lock()
		for mailbox == 1 {
			sendCond.Wait()
		}
		mailbox = 1
		fmt.Println("发送邮件：", mailbox)
		lock2.Unlock()
		recvCond.Signal()
	}
}

func receive() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		lock2.RLock()
		for mailbox == 0 {
			recvCond.Wait()
		}
		mailbox = 0
		fmt.Println("获取邮件：", mailbox)
		lock2.RUnlock()
		sendCond.Signal()
	}
}

func main12() {
	//wg1.Add(2)
	//go add()
	//go add()
	//wg1.Wait()
	////多次运行结果不一样
	//fmt.Println(x)

	//读写锁测试
	//for i := 0; i < 10; i++ {
	//	wg1.Add(1)
	//	go write()
	//}
	//
	//for i := 0; i < 1000; i++ {
	//	wg1.Add(1)
	//	go read()
	//}
	//wg1.Wait()

	//信号量测试
	go send()
	go receive()

	time.Sleep(20 * time.Second)

}
