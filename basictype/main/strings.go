package main

import (
	"fmt"
	//别名s
	s "strings"
)

func main() {
	//	s1 := "hello" + "word"
	//
	//	//反引号表示多行字符串,其中定义的转义字符都无效，文本会原样输出
	//	s2 := `A "raw" string literal can include line breaks.`
	//
	//	fmt.Println(len(s1))
	//	fmt.Println(string(s1[0:5]))
	//	fmt.Println(s2)

	fmt.Println(s.Contains("test", "e"))
	fmt.Println(len("hello"))
	//输出 101
	fmt.Println("hello"[1])
	//输出 e
	fmt.Println(string("hello"[1]))
}
