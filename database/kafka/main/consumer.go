package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

func receive() {
	consumer, err := sarama.NewConsumer([]string{"192.168.110.41:9092,192.168.110.42:9092,192.168.110.43:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer:%v\n", err)
		return
	}
	//根据topic获取到所有分区
	partitions, err := consumer.Partitions("go_test")
	if err != nil {
		fmt.Printf("fail to get list of partitions:err%v\n", err)
		return
	}
	fmt.Println(partitions)
	//遍历所有分区
	for partition := range partitions {
		//每个分区创建一个消费者
		consumePartition, err := consumer.ConsumePartition("go_test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer consumePartition.AsyncClose()
		//异步从每个分区消费
		go func(partitionConsumer sarama.PartitionConsumer) {
			for message := range consumePartition.Messages() {
				fmt.Printf("partition:%d offset:%d key:%v value:%v", message.Partition, message.Offset, message.Key, message.Value)
			}
		}(consumePartition)
	}
}
