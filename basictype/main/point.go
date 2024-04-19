package main

import "fmt"

//func main() {
//a := 3
////取变量a的地址，将指针保存到b中，此时b的类型就是一个int类型指针
//b := &a
//fmt.Println(&a)
////%p:指针占位符,%T:类型
////输出结果：b:0xc00001a098 type:*int
//fmt.Printf("b:%p type:%T\n", b, b)
////变量b存放的是a的地址，&b获取的是b的地址
//fmt.Println(&b)
//
////指针取值，根据指针取内存取值，此时取的就是指针b保存的值3
//c := *b
//fmt.Printf("type of c:%T\n", c)
////%v:相应值的默认格式
//fmt.Printf("value of c:%v\n", c)
//
////总结，&取出地址，*根据地址取出地址指向的值
////1、对变量进行取地址（&）可以获得这个变量的指针变量
////2、指针变量的值是指针地址
////3、对指针变量进行取值（*）操作，可以获取指针变量指向的原变量的值
//
////指针传值,go中的函数传参都是值拷贝
//a1 := 10
//modify1(a1)
////还是10
//fmt.Println(a1)
//modify2(&a1)
//fmt.Println(a1)
//
////空指针,当一个指针被定义后没有分配到任何变量时，它的值为nil
//nullPoint()

//	newTest()
//
//	makeTest()
//
//	test()
//}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func nullPoint() {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%s\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空")
	}

}

func newTest() {
	//指针是引用类型，需要初始化后才会拥有内存空间，才可以给它赋值
	var a *int
	//程序崩溃，因为未初始化
	//*a = 100
	fmt.Println(a)
	//fmt.Println(*a)

	//使用new函数初始化，该函数不常用，它返回的是一个类型的指针，并且该指针对应的值为该类型的零值
	a = new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) //*int
	fmt.Printf("%T\n", b) //*bool

	fmt.Println(*a) //0
	fmt.Println(*b) //false

	*a = 10
	fmt.Println(*a)
}

func makeTest() {
	//make也用于分配内存，但它只用于map、slice以及chan的内存创建
	//并且返回就是类型本身，而不是他们的指针类型，因为这三种类型就是引用类型（引用类型意味着必须要初始化，而该函数就是完成初始化工作的）
	var b map[string]int
	b = make(map[string]int)
	b["测试"] = 100
	fmt.Println(b)
}

/**
通过指针修改变量值
*/
func test() {
	var a int
	fmt.Println(&a)
	var ptr *int
	ptr = &a
	*ptr = 2000
	fmt.Println(a)
}
