package main

import "fmt"

/*
	函数与函数间需要交互数据才能体现并发执行函数的意义。
	go并发模型是csp，提倡通过通信共享内存而不是通过共享内存而实现通信（共享内存势必会带来加锁造成的性能问题）。
	goroutine是go程序并发的执行体，channel则是它们之间的连接。channel是可以让一个goroutine发送
	特定值到另一个goroutine的通信机制
*/

func main3() {
	//声明一个传递int类型的channel，channel是特殊的引用类型，不实例化默认是nil
	//var ch chan int
	//fmt.Println(ch)

	//1、发送操作，使用<-符号，channel在左边
	ch := make(chan int)
	//将10发送到ch中
	ch <- 10
	//2、接收操作，使用<-符号，channel在右边
	x := <-ch
	//会阻塞，因为make没有指定缓冲区容量，所以创建的就是无缓冲通道，必须有接收才能发送（需要goroutine接收）
	fmt.Println(x)
	//3、关闭，内置函数close
	close(ch)
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main4() {
	//无缓冲区通道，即同步通道，发送操作会一直阻塞知道有接收方，反之亦然
	ch := make(chan int)
	//main goroutine发送，必须是另外一个goroutine接收。先有接收方，才能发送数据
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")
}

func main5() {
	//有缓冲区的通道，只需要容量大于0
	ch := make(chan int, 1)
	ch <- 10
	a := <-ch
	//死锁，此时通道已经没数据了
	b := <-ch
	fmt.Println("发送成功", a, b)
}

//close。当管道不往里存值或取值的时候一定要关闭通道。通道是可以被垃圾回收机制回收的
func main6() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
		//1、对一个已关闭的通道再发送值就会导致panic
		//c <- 1
		//2、关闭一个已关闭的通道会导致panic
		//close(c)
	}()
	for {
		//对一个已关闭的通道进行接收会一直获取值直到通道为空
		//<-channel返回两个值，第一个是具体的数据，第二个是通道是否关闭
		if v, ok := <-c; ok {
			fmt.Println(v)
		} else {
			break
		}
	}
	fmt.Println("main结束")
}

//如何判断通道是否关闭，for range方式更常用
func main7() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	//发送数据到管道
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	//接收数据
	go func() {
		for {
			//通道关闭后再取值ok=false
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	//在主goroutine中从ch2中接收值打印
	for i := range ch2 {
		fmt.Println(i)
	}
}

//单向通道，限制通道只能进行发送或者接收
//只能接受发送通道
func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

//接受发送通道和接收通道
func square(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

//接受接收通道
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main8() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go square(ch2, ch1)
	//直接在main goroutine中打印
	printer(ch2)
}
