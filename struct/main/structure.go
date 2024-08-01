package main

import "fmt"

//结构体可以理解成java当中的类，只不过没有类的继承等面向对象的概念

type person struct {
	name string
	city string
	age  int8
}

type student struct {
	name string
	age  int
}

func main() {
	//1、基本实例化，先使用变量进行初始化，然后赋值进行实例化
	var p1 person
	p1.name = "likui"
	p1.city = "杭州"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)

	//2、使用new实例化
	var p2 = new(person)
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
	p2.name = "测试"
	p2.age = 18
	fmt.Printf("p2=%#v\n", p2)

	//3、取结构体地址实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("p3=%#v\n", p3)
	p3.name = "博客"
	p3.age = 18

	//4、使用键值对初始化
	p5 := person{
		name: "likui",
		city: "上海",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5)

	//也可以对结构体指针进行键值对初始化
	p6 := &person{
		name: "sss",
		city: "sss",
		age:  20,
	}
	fmt.Printf("p6=%#v\n", p6)

	//5、使用值的列表初始化，即不写键，直接写值
	//需要注意：1、必须初始化结构体的所有字段；2、初始值的填充顺序必须与字段在结构体中的声明顺序一样；3、该方式不能和减值初始化方式混用
	p7 := &person{
		"lll",
		"sss",
		20,
	}
	fmt.Printf("p7=%#v\n", p7)

	//面试题
	m := make(map[string]*student)
	stus := []student{
		{name: "a", age: 10},
		{name: "b", age: 20},
		{name: "c", age: 30},
	}
	fmt.Printf("stus:%p\n", stus)
	for _, stu := range stus {
		fmt.Println("std.name:", stu.name)
		//会发现&stu的地址都是同一个，因为stu是当前元素的副本，它存储在一个不同于原始切片元素的内存地址
		//它存储的是最后一次迭代中stu的值
		fmt.Printf("stu:%p\n", &stu)
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	fmt.Printf("m=%v\n", m)

	p := newPerson("lllk", "test", 29)
	fmt.Printf("%#v\n", p)
}

func anonyStruct() {
	var user struct {
		name string
		age  int
	}
	user.name = "likui"
	user.age = 18
	fmt.Printf("%#v\n", user)
}

//结构体内存布局
func structMemory() {
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{1, 2, 3, 4}

	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
}

/*
struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，
所以该构造函数返回的是结构体指针类型
*/
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
