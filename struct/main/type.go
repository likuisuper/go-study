package main

//类型定义，定义一个全新的类型
type NewInt int

//类型别名，本质上别名和type都是同一个类型，编译完成时不会有MyInt类型
type MyInt = int

//func main() {
//	var a NewInt
//	var b MyInt
//	fmt.Printf("type of a:%T\n", a) //NewInt
//	fmt.Printf("type of b:%T\n", b) //int
//}
