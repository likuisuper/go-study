package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		//最初的连接数量
		MaxIdle: 16,
		//最大连接数量，不确定可以0，表示自动定义，按需调整
		MaxActive: 0,
		//连接关闭时间300秒，300秒不使用自动关闭
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "192.168.100.16:39502", redis.DialPassword("eJVjdvBlqVYKJI9A"), redis.DialDatabase(3))
		},
	}
}

func main() {
	//从连接池获取一个连接
	c := pool.Get()
	//运行结束，将连接放回连接池
	defer c.Close()

	_, err := c.Do("set", "abc", 200)
	if err != nil {
		fmt.Println("set error,", err)
		return
	}
	r, err := redis.Int(c.Do("get", "abc"))
	if err != nil {
		fmt.Println("get error,", err)
		return
	}
	fmt.Println(r)

	//关闭连接池
	pool.Close()
}
