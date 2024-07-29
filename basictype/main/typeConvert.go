package main

//go语言不允许隐式转换，且转换必须发生在两种相互兼容的类型之间
//func main() {
//	i := 90
//	f := float64(i)
//	u := uint(i)
//	//将等于字符串Z
//	s := string(rune(i))
//
//	fmt.Println(f, u, s)
//
//	//如何获取int字符串
//	//输出原始值，即90
//	s1 := strconv.Itoa(i)
//	fmt.Println(s1)
//}
