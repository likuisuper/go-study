package main

import "fmt"

//全局定义数组
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var str = [5]string{3: "hello work", 4: "tom"}

//func main() {
//	//定义数组长度为3，未初始化的元素值为0，这里只初始化了2个
//	a := [3]int{1, 2}
//	fmt.Println(a)
//	//通过初始化值确定数组长度
//	primes := [...]int{2, 3, 4, 5, 6, 7, 9}
//	fmt.Println(len(primes))
//	fmt.Println(primes)
//	//左闭右开
//	fmt.Println(primes[0:3])
//
//	//使用索引号初始化元素
//	c := [5]int{2: 100, 4: 200}
//	fmt.Println(c)
//
//	d := [...]struct {
//		name string
//		age  int
//	}{
//		{"user1", 10},
//		{"user2", 18},
//	}
//	fmt.Println(arr0, arr1, str)
//	fmt.Println(d)
//
//	manyArray()
//
//	forArray()
//}

//全局
var arr01 [5][3]int
var arr02 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

/**
多维数组
*/
func manyArray() {
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println(arr01, arr02)
	fmt.Println(a, b)
}

/**
数组遍历
*/
func forArray() {
	var twoDimension [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimension[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoDimension)
}
