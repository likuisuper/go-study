package main

import "fmt"

/*
	函数特点：
		1、支持不定变参
		2、支持多返回值
		3、支持匿名函数和闭包
		4、函数也是类型，可以赋值给变量
		5、不支持嵌套，一个包不能有两个名字一样的函数
		6、不支持重载
		7、不支持默认参数
*/

//类型相同的相邻参数，参数类型可以合并，多返回值必须用括号
func test(x, y int, s string) (int, string) {
	n := x + y
	return n, fmt.Sprintf(s, n)
}

//匿名函数，即可以不用定义方法名称
func test1(fn func() int) int {
	return fn()
}

//定义函数类型
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

//func main() {
//	//fmt.Println(test(1, 3, "sss"))
//
//	s1 := test1(func() int {
//		return 100
//	})
//	//就相当于java中的匿名内部类，由于format还定义了参数，所以后面要传递参数
//	s2 := format(func(s string, x, y int) string {
//		return fmt.Sprintf(s, x, y)
//	}, "%d,%d", 10, 20)
//
//	println(s1)
//	println(s2)
//
//}
