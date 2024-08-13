package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var c redis.Conn

func init() {
	conn, err := redis.Dial("tcp", "192.168.100.16:39502", redis.DialPassword("eJVjdvBlqVYKJI9A"), redis.DialDatabase(3))
	if err != nil {
		fmt.Println("connect redis fail")
		return
	}
	fmt.Println("connect redis success")
	c = conn
}

func main() {
	defer c.Close()
	//string()
	//list()
	hash()
}

func string() {
	_, err := c.Do("SET", "go:test:", "hello")
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	r, err := redis.String(c.Do("GET", "go:test:"))
	if err != nil {
		fmt.Println("get error,", err)
		return
	}
	fmt.Println(r)
}

func expire() {
	_, err := c.Do("expire", "go:test:", 10)
	if err != nil {
		fmt.Println("expire error,", err)
		return
	}
}

func list() {
	_, err := c.Do("lpush", "book:list", "abc", "ceg", 300)
	if err != nil {
		fmt.Println("lpush error,", err)
		return
	}
	r, err := redis.String(c.Do("lpop", "book:list"))
	if err != nil {
		fmt.Println("rpop error,", err)
		return
	}
	fmt.Println(r)
}

func hash() {
	_, err := c.Do("hset", "books", "abc", 100)
	if err != nil {
		fmt.Println("hset error,", err)
		return
	}
	//获取指定key
	r, err := redis.Int(c.Do("hget", "books", "abc"))
	if err != nil {
		fmt.Println("hget error,", err)
		return
	}
	fmt.Println(r)
}
