package main

import "fmt"

type UserTest struct {
	id   string
	name string
}

type Manager struct {
	//匿名字段
	UserTest
	title string
}

func (self *UserTest) ToString() string {
	return fmt.Sprintf("UserTest: %p,%v", self, self)
}

//定义同名方法，即可实现"ovveride"，通过左边的箭头可以
func (self *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p,%v", self, self)
}

func main2() {
	m := Manager{UserTest{"1", "tom"}, "admin"}
	fmt.Println(m.ToString())
	fmt.Println(m.UserTest.ToString())
}
