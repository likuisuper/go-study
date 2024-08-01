package main

import "fmt"

/*
	go中有两种传递方式：
		1、值传递，即调用函数时将实际参数复制一份传递到函数中，在函数中对参数进行修改，不会影响到实际参数
		2、引用传递，调用函数时将实际参数的地址传递到函数中，在函数中对参数进行修改，会影响到实际参数
	默认使用值传递
	注意1：
		无论是值传递还是引用传递，传递给函数的都是变量的副本，不过值传递是值的拷贝，引用传递是地址的拷贝，
		一般来说，地址拷贝更高效，而值拷贝取决于拷贝的对象大小，对象越大则性能越低
	注意2：
		map,slice,chan,指针,interface默认是引用传递
*/

func swap1(x, y int) {
	var temp int
	temp = x
	x = y
	y = temp
}

//引用传递
func swap2(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

/*
	可变参数，即参数不固定，但后面的类型固定
	golang可变参数本质上是一个slice，只能有一个，且必须是最后一个
	函数赋值时，如果传递的是数组或者切片，参数后面必须加上...
*/
func variable1(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}

//func main() {
//	x := 7
//	swap1(x, 3)
//	//输出结果仍为7
//	println(x)
//
//	a, b := 1, 2
//	swap2(&a, &b)
//	//输出结果为2,1
//	println(a)
//	println(b)
//
//	println(variable1("sum:%d", 1, 2, 3))
//
//	s := []int{1, 2, 3}
//	//s后面必须跟上...，否则编译报错
//	res := variable1("sum2:%d", s...)
//	println(res)
//}
