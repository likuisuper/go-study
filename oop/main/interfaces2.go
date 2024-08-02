package main

import "fmt"

/*
	一个类型可以实现多个接口，多个接口彼此之间独立，不知道对方的实现
*/
type Sayer2 interface {
	say()
}
type Mover2 interface {
	move()
}

type dog3 struct {
	name string
}

type car struct {
	brand string
}

//麻烦的地方在于，如果要实现的接口方法很多，只能写很多的方法一一实现，而不能同时实现这些方法
func (d dog3) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}
func (d dog3) move() {
	fmt.Printf("%s会动\n", d.name)
}

//2、多个类型实现同一个接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

//3、接口嵌套
type animal interface {
	Sayer2
	Mover2
}
type cat2 struct {
	name string
}

func (c cat2) say() {
	fmt.Println("喵喵喵")
}
func (c cat2) move() {
	fmt.Println("猫会动")
}

func main3() {
	//var x Sayer2
	//var y Mover2
	//var a = dog3{name: "旺财"}
	////必须要赋值给接口，不然接口没有实现
	//x = a
	//y = a
	//x.say()
	//y.move()
	//
	//var b = car{
	//	brand: "法拉利",
	//}
	//
	//y = b
	//y.move()

	//嵌套接口
	var x animal
	x = cat2{name: "花花"}
	x.say()
	x.move()
}
