package main

import "fmt"

func main() {
	//ifTest()
	//switchTest()
	forTest()
}

func ifTest() {
	a := 10

	if a > 20 {
		fmt.Println(">")
	} else if a < 20 {
		fmt.Println("<")
	} else {
		fmt.Println("=")
	}

	x := 0
	//初始化的变量可以在语句块中使用
	if n := "abc"; x > 0 {
		println(n[2])
	} else if x < 0 {
		println(n[1])
	} else {
		println(string(n[0]))
	}
}

func switchTest() {
	x := 42.0
	switch x {
	case 0:
	case 1, 2:
		fmt.Println("multiple matchs")
	case 42:
		fmt.Println("reached")
	case 43:
		fmt.Println("unreached")
	default:
		fmt.Println("optional")
	}
}

func forTest() {
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
}
