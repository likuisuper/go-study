package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

func main() {
	send()
}

func send() {
	config := &sarama.Config{}
	config.Version = sarama.V2_0_0_0
	config.ClientID = "javagroup"
	//-1：发送完数据需要leader和所有follower都确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	//分区策略为轮询
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	//是否返回成功的消息
	config.Producer.Return.Successes = true
	//使用同步生产者，必须配置errors为true
	config.Producer.Return.Errors = true
	config.Producer.MaxMessageBytes = 1024
	config.Producer.Timeout = 5 * time.Second
	//必须大于0，一次发送请求的数量
	config.Net.MaxOpenRequests = 5
	config.Net.DialTimeout = 1 * time.Second
	config.Net.ReadTimeout = 5 * time.Second
	config.Net.WriteTimeout = 5 * time.Second
	config.Net.KeepAlive = 10 * time.Second
	config.Admin.Timeout = 10 * time.Second
	msg := &sarama.ProducerMessage{}
	msg.Topic = "go_test"
	msg.Value = sarama.StringEncoder("this is a test log")
	//连接，这里使用的是同步生产者，会阻塞当前线程，所以吞吐量较低，但是可以立即直到消息是否发送成功
	producer, err := sarama.NewSyncProducer([]string{"192.168.110.41:9092,192.168.110.42:9092,192.168.110.43:9092"}, config)
	if err != nil {
		fmt.Println("producer closed,err", err)
		return
	}
	defer producer.Close()
	//发送消息

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed,err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", partition, offset)
}
