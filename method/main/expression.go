package main

import "fmt"

/*
	golang表达式，分为两种
	1、method value：instance.method(args...)，需要绑定实例，所以不需要显示传参，会复制receiver
	2、method expression: <type>.func(instance, args...)，需要显示传参
*/

type EUser struct {
	id   int
	name string
}

func (u *EUser) Etest() {
	fmt.Printf("%p,%v\n", u, u)
}

func main3() {
	u := EUser{1, "tom"}
	u.Etest()

	mValue := u.Etest
	//隐式传递 receiver
	mValue()

	mExpression := (*EUser).Etest
	//显示传参
	mExpression(&u)
}
