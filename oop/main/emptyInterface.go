package main

import "fmt"

/*
	空接口，即没有定义任何方法的接口。因此任何类型都实现了空接口，可以类比为java的object
*/

//1、空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	//定义一个空接口x
	var x interface{}
	s := "sdfdf"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", i, i)
	b := true
	fmt.Printf("type:%T value:%v\n", b, b)

	//2、空接口作为map的值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "lk"
	studentInfo["age"] = 20
	studentInfo["married"] = false
	fmt.Println(studentInfo)

	//3、如何获取空接口的类型和值？使用类型断言：x.(T)
	var m interface{}
	m = "sdfsdfsd"
	//x.(T)会返回两个值，第一个参数是x转换为T类型后的变量，第二个参数是一个布尔值，若为true则表示断言成功，false断言失败
	v, ok := m.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}

	//使用switch替换if
	justifyType(m)
}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is string,value is %v\n", v)
	case int:
		fmt.Printf("x is int,value is %v\n", v)
	case bool:
		fmt.Printf("x is bool,value is %v\n", v)
	default:
		fmt.Println("unsupported type")
	}
}
