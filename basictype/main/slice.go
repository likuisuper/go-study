package main

import "fmt"

//切片
//func main() {
//	//createSlice()
//	//initSlice()
//
//	//makeForSlice()
//
//	//structSlice()
//
//	//appendSlice()
//
//	//resize()
//
//	//copyslice()
//
//	//traverse()
//}

//创建切片各种方式
func createSlice() {
	//声明切片，切片跟数组不一样，定义数组必须指明长度
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("非空")
	}
	// :=
	s2 := []int{}

	// 3.make
	var s3 []int = make([]int, 1)
	fmt.Println(s1, s2, s3)

	//4.初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	//5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	//前包后不包
	s6 = arr[1:4]
	fmt.Println(s6)
}

//切片初始化
//全局
var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

//从切片s的索引位置start到end处所获得的切片
//var slice0 []int=arr[start:end]
var slice0 []int = arr[2:8]

//从切片s的索引位置0到end处所获得的切片
//var slice1 []int=arr[:end]
var slice1 []int = arr[0:6] //可以简写为arr[:6]
//从切片s的索引位置start到len(s)-1处所获得的切片
//var slice2 []int = arr[start:]
var slice2 []int = arr[5:10] //可以简写为arr[5:]
//从切片索引s索引0到len(s)-1处所获得的切片
//var slice3 []int = arr[:]
var slice3 []int = arr[0:len(arr)] //可以简写为arr[:]
var slice4 = arr[:len(arr)-1]      //相当于去掉最后一个元素

func initSlice() {
	fmt.Printf("全局变量：arr %v\n", arr)
	fmt.Printf("全局变量：slice0 %v\n", slice0)
	fmt.Printf("全局变量：slice1 %v\n", slice1)
	fmt.Printf("全局变量：slice2 %v\n", slice2)
	fmt.Printf("全局变量：slice3 %v\n", slice3)
	fmt.Printf("全局变量：slice4 %v\n", slice4)

	//局部变量
	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	var slice5 []int = arr2[2:8]
	fmt.Printf("局部变量：slice5 %v\n", slice5)

}

//通过make创建切片
var slice6 []int = make([]int, 10)
var slice7 = make([]int, 10)

//长度和容量都为10
var slice8 = make([]int, 10, 10)

func makeForSlice() {
	fmt.Printf("make全局slice6 : %v\n", slice6)
	fmt.Printf("make全局slice7 : %v\n", slice7)
	fmt.Printf("make全局slice8 : %v\n", slice8)

	fmt.Println("--------------------")

	slice9 := make([]int, 10)
	slice10 := make([]int, 10, 10)
	fmt.Printf("make局部slice9 : %v\n", slice9)
	fmt.Printf("make局部slice10 : %v\n", slice10)
	fmt.Println(slice10, len(slice10), cap(slice10))

	fmt.Println("-----------------")
	//通过指针修改切片中元素的值
	s := []int{0, 1, 2, 3}
	p := &s[2]
	*p += 100
	fmt.Println(s)

	//[][]T，是指元素类型为[]T
	data := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33},
	}
	fmt.Println(data)

}

func structSlice() {
	d := [5]struct {
		x int
	}{}
	s := d[:]

	d[1].x = 10
	d[2].x = 20

	fmt.Println(d)
	fmt.Println(s)
	//&d[0]的地址就是&d的地址
	fmt.Printf("%p,%p,%p\n", &d, &d[0], &d[1])
}

func appendSlice() {
	var a = []int{1, 2, 3}
	var b = []int{4, 5, 6}
	c := append(a, b...)
	fmt.Printf("slice c : %v\n", c)
	d := append(c, 7)
	fmt.Printf("slice d: %v\n", d)
	e := append(d, 8, 9)
	fmt.Printf("slice e : %v\n", e)
}

func resize() {
	data := [...]int{0, 1, 2, 3, 4, 10, 0}
	//data[start:high:cap]
	//s的长度为2，容量为3
	s := data[:2:3]
	//追加两个，即容量为4，超过原来容量3，会重新分配
	s = append(s, 100, 200)

	fmt.Println(s, data)
	//可以发现，两个数组地址不一样，如果上面只追加一个值，那么两个数组的地址是一样的
	fmt.Println(&s[0], &data[0])

	var a = []int{1, 3, 4, 5}
	b := a[1:2]
	//切片b的内存地址指向切片a中从索引1开始的地方，所以容量是3
	fmt.Printf("slice b :%v ,len(b) :%v, cap(b) :%v\n", b, len(b), cap(b))
	//因为b只有一个元素[3]，所以切片c将包含从b索引0开始的所有元素[3,4,5]
	//c := b[0:1]
	c := b[0:3]
	fmt.Printf("slice c :%v ,len(c) :%v, cap(c) :%v\n", c, len(c), cap(c))
}

func copyslice() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(data))
	s1 := data[8:] //8,9
	s2 := data[:5] //0,1,2,3,4
	fmt.Println(s1)
	fmt.Println(s2)
	//copy用于在两个slice间复制数据
	copy(s2, s1)
	fmt.Println(s2)
}

func traverse() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//复制整个切片
	slice := data[:]
	for index, value := range slice {
		fmt.Printf("index :%v, value :%v\n", index, value)
	}
}
