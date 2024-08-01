package main

import "fmt"
import "errors"

//1、利用recover处理panic指令，defer必须放在panic之前定义，
func exception1() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) //将interface{}转换为具体类型
		}
	}()

	fmt.Println("defer之后")

	panic("panic error")

	fmt.Println("panic之后")
}

//延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获
func exception2() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}

//多个defer逆序执行，即先进后出
//recover只有在延迟内直接调用才会终止错误，否则总是返回nil，任何未捕获的错误都会沿堆栈向外传递
func exception3() {
	defer func() {
		fmt.Println(recover()) //有效
	}()

	defer recover() //无效，属于延迟语句

	defer fmt.Println(recover()) //无效，属于延迟语句

	defer func() {
		func() {
			println("defer inner")
			recover() //延迟函数内部调用，不是直接调用，无效
		}()
	}()

	panic("test panic")
}

/*
	errors.New和fmt.Errorf函数用于创建实现error接口的错误对象。
	通过判断错误对象实例来确定具体错误类型
*/
var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}

//实现类似try catch的异常处理
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()

	//执行业务逻辑
	fun()
}

func main() {
	//结果：defer之后 panic error
	//exception1()

	//exception2()

	//结果：defer inner,<nil>,test panic
	//exception3()

	//----错误类型begin
	//defer func() {
	//	fmt.Println(recover())
	//}()
	//
	//switch z, err := div(10, 0); err {
	//case nil:
	//	println(z)
	//case ErrDivByZero:
	//	panic(err)
	//}
	//----错误类型end

	//try catch
	Try(func() {
		fmt.Println("这是业务逻辑")
		panic("抛出业务异常")
	}, func(err interface{}) {
		fmt.Println(err)
	})

}
