package main

import "fmt"

//常量
//func main() {
//	definition()
//
//}

func definition() {
	const pi = 3
	const e = 2

	const (
		a = 3
		b = 2
	)

	const (
		n1 = 100
		//省略了值则表示和上面一行的值相同
		n2
		n3
	)

	fmt.Println(n1, n2, n3)
}

func iotaTest() {
	const (
		//iota可以理解为const语句块中的行索引
		n1 = iota //0
		n2
		n3
		n4
	)

	const (
		a1 = iota
		a2
		//使用_跳过某些值
		_
		a4
	)

	const (
		//iota声明中间插队
		b1 = iota
		b2 = 100
		b3 = iota
		b4
	)
}
