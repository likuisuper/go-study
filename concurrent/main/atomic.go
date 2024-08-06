package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
	加锁会涉及内核态的上下文切换，会比较耗时、代价较高。
	针对基本数据类型可以直接使用原子操作保证并发安全。它在用户态就可以完成，因此性能比加锁更好。
	Mutex不能保证操作的原子性，这和java中的锁是有区别的
*/

var i int64
var l sync.Mutex
var w sync.WaitGroup

//普通版加函数
func normalAdd() {
	i++
	w.Done()
}

//互斥锁版加函数
func mutexAdd() {
	l.Lock()
	i++
	l.Unlock()
	w.Done()
}

//原子版加函数
func atomicAdd() {
	//为什么原子类中这些方法的操作对象都是指针类型的呢？
	//如果是值类型，那么操作的就是副本，这样修改副本的值就和原值无关了，无法实现原子类的语义
	atomic.AddInt64(&i, 1)
	w.Done()
}

func main13() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		w.Add(1)
		//耗时3Ms左右
		//go normalAdd()
		//比上面耗时高
		//go mutexAdd()
		//比互斥锁耗时低
		go atomicAdd()
	}
	w.Wait()
	end := time.Now()
	fmt.Println(i)
	fmt.Println(end.Sub(start))
}
