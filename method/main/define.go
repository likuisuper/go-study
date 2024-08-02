package main

import "fmt"

/*
	golang方法总是绑定对象实例，并隐式将实例作为第一实参。
		1、只能为当前包内命名类型定义方法
		2、参数receiver可任意命名，如方法中未使用可忽略参数名
		3、参数receiver类型可以是T或者*T，基类型不能是指针或接口
		4、不支持方法重载
		5、可用实例value或point调用方法，编译器自动转换
	一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
	接受者命名建议为类型的第一个小写字母
	普通函数和方法的区别：
		1、对于普通函数，接受者为值类型时，不能将指针类型的数据直接传递，反之亦然
		2、对于方法，接收者为值类型时，可以直接用指针类型的变量调用方法，反之亦然
*/

type User struct {
	Name  string
	Email string
}

//方法
func (u User) Notify() {
	//方法接受者是值类型，所以即使使用指针调用，也是对副本的操作
	fmt.Printf("%v:%v \n", u.Name, u.Email)
}

//指针类型方法
func (u *User) Notify2() {
	//方法接受者是指针类型，即使使用值类型调用，函数内部也是对指针的操作
	fmt.Printf("%v:%v \n", u.Name, u.Email)
}

//值类型调用和指针调用的区别，只需将方法还原成函数即可
type Data struct {
	x int
}

func (self Data) ValueTest() {
	fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest() {
	fmt.Printf("Pointer: %p\n", self)
}

func main1() {
	////值类型调用方法
	//u1 := User{"golang", "sfsdfd@golang.com"}
	//u1.Notify()
	//
	////指针类型调用方法
	//u2 := User{"go", "gl@go.com"}
	//u3 := &u2
	////由于Nofity方法接受者是值类型，所以这里会使用解引用指针调用，但是函数内部还是对副本的操作，而不是指针操作
	//u3.Notify()

	//值类型调用方法
	u1 := User{"golang", "sfsdfd@golang.com"}
	u1.Notify2()

	//指针类型调用方法
	u2 := User{"go", "gl@go.com"}
	u3 := &u2
	u3.Notify2()

	//值类型和指针调用区别测试
	//输出结果发现，无论是值变量d调用PointerTest还是指针变量p调用PointerTest，他们的结果都是一样，因为都是对同一指针的操作
	//而d.ValueTest()和p.d.ValueTest()结果不一样，因为函数内部是对副本的操作，不是对变量d和p的操作
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)

	d.ValueTest()
	d.PointerTest()

	p.ValueTest()
	p.PointerTest()
}
