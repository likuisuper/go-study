package main

import "fmt"

/*
	继承是通过牺牲代码简洁性来换取可扩展性，而且这种可扩展性是通过侵入的方式实现的。
	go中并没有继承的概念，它所做的是通过嵌入字段的方式实现了类型之间的组合。
	只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段
*/

type Person struct {
	name string
	sex  string
	age  int
}

//自定义类型
type mystr string

type Student struct {
	Person
	id   int
	addr string
	//同名字段的情况
	name string
	mystr
}

//指针类型匿名字段
type Student2 struct {
	*Person
	id   int
	addr string
}

func main1() {
	//s1 := Student{Person{"lk", "man", 20}, 1, "hz"}
	//fmt.Println(s1)
	//
	////当不需要初始化所有属性时，可以显示声明需要初始化的字段，通过key value的方式
	//s2 := Student{Person: Person{"lk", "man", 20}}
	//fmt.Println(s2)
	//
	//s3 := Student{Person: Person{name: "lk"}}
	//fmt.Println(s3)

	//出现同名字段的情况
	var s Student
	//给自己字段赋值
	s.name = "lk"
	fmt.Println(s)

	//若给父类同名字段赋值
	s.Person.name = "ssss"
	fmt.Println(s)

	//所有内置类型和自定义类型都可以作为匿名字段使用
	s4 := Student{Person{"lk", "man", 20}, 1, "hz", "lk1", "ssss"}
	fmt.Println(s4)

	s5 := Student2{&Person{"lk", "man", 20}, 1, "hz"}
	//输出：{0xc000074480 1 hz}
	fmt.Println(s5)

}
