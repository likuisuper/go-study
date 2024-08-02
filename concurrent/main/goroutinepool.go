package main

import (
	"fmt"
	"math/rand"
)

/*
	本质上是生产者消费者模型
	可以有效控制goroutine数量，防止暴涨
	需求：随机生成数字，计算该数字的各个位数之和
*/

type Job struct {
	Id int
	//需要计算的随机数
	RandNum int
}

type Result struct {
	//因为在创建管道的时候使用了指针，所以这里也定义成指针类型，如果创建管道的时候没有使用指针，这里不需要定义成指针类型
	job *Job
	//求和
	sum int
}

func main() {
	//1、job管道
	//为什么要使用chan *job而不是chan job呢？
	//chan *job是一个指向job结构体实例的指针类型的通道，向这种通道发送数据时，实际是发送指向job结构体的指针
	//而chan job是job结构体类型的通道，它直接存储job结构体的实例，向这种通道发送数据时，实际上是复制job结构体的实例
	//所以为了数据量造成复制间的开销，这里使用指针类型
	jobChan := make(chan *Job, 128)
	//2、结果管道
	resultChan := make(chan *Result, 128)
	//3、创建工作池
	createPool(64, jobChan, resultChan)
	//4、开启打印协程
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(resultChan)

	var id int
	//循环创建job，输入到管道
	for {
		id++
		//生成随机数
		rNum := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: rNum,
		}
		jobChan <- job
	}

}

//创建工作池，num表示开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			//遍历job管道所有数据，进行相加
			for job := range jobChan {
				rNum := job.RandNum
				var sum int
				for rNum != 0 {
					temp := rNum % 10
					sum += temp
					rNum /= 10
				}
				r := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}
