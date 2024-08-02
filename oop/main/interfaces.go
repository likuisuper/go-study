package main

import "fmt"

/*
	接口定义了一组规范，在go中，接口是一种类型（很重要），一种抽象类型。
	接口在命名时，一般会在接口名后面加er，如writer
	当方法名和接口名首字母都是大写时，这个方法可以被接口所在的包之外的代码访问，如同java的public
*/

//1、实现接口的条件：一个对象只要全部实现了接口中的方法，那么就实现了这个接口。
//也就是说，方法就是一个需要实现的方法列表
type Sayer interface {
	say()
}

type cat struct{}

type dog struct{}

//dog实现Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

//cat实现Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

//3、值接收者和指针接收者实现接口的区别
type Mover interface {
	move()
}
type dog2 struct{}

//值接收者实现接口
//func (d dog2) move() {
//	fmt.Println("狗会动")
//}
//指针接收者实现接口
func (d *dog2) move() {
	fmt.Println("狗会动")
}

func main2() {
	//2、接口类型变量能够储存所有实现了该接口的实例
	//var x Sayer
	//a := cat{}
	//b := dog{}
	//x = a //可以把cat实例直接赋值给x
	//x.say()
	//x = b //可以把dog实例直接赋值给x
	//x.say()

	//var x Mover
	//wangcai := dog2{} //旺财是dog2类型
	//x = wangcai //如果是值接收者实现接口，x可以接收dog2类型，如果是指针接收者实现接口，这里会编译报错
	//fugui := &dog2{} //富贵是*dog2类型
	//x = fugui //x可以接收*dog类型
	//x.move()
	//可以发现，使用值接收者实现接口，不管是结构体还是结构体指针类型的变量都可以赋值给改接口变量
	//因为go中有对指针类型变量求值的语法糖
}
