package main

import "fmt"

/**
defer延迟调用，被延迟执行的是defer函数，而不是defer语句
特性：
	1、defer是先进后出，因为后面的资源会依赖前面的资源
		原理：defer语句每次执行的时候，go会把它携带的defer函数及其参数值另行存储到一个链表中，相当于一个栈
	2、defer语句中的变量，在defer声明时就决定了
	3、defer函数调用会在它即将结束执行的那一刻执行
用途：
	1、关闭文件句柄
	2、释放锁资源
	3、释放数据库连接
*/

func defer1() {
	var whatever [5]struct{}
	//输出4,3,2,1,0
	for i := range whatever {
		//defer语句延迟执行的函数会在包含defer语句的函数执行结束时才被调用，但是它们被调用时会使用当前的值,
		//而i是一个变量，不是当前循环的值
		defer fmt.Println(i)
	}
}

//defer碰上闭包，下面是一个闭包函数，因为匿名函数内部使用到了外部函数的变量
//由于闭包用到的变量i在执行的时候已经变成4，所以输出全部是4
func defer2() {
	var whaterver [5]struct{}
	for i := range whaterver {
		//将i的值传递给匿名函数，以便在defer执行时捕获当前i的值，这样可以确保defer语句使用的是当前循环的值
		defer func() {
			fmt.Println(i)
		}()
	}
}

//func main() {
//	//defer1()
//	//defer2()
//
//	ts := []Test{{"a"}, {"b"}, {"c"}}
//	for _, t := range ts {
//		//输出结果：c closed c closed c closed
//		//同样是循环，为什么输出的是同一个值而不是上面的逆序输出呢？
//		//在结构体中说过，这里的t它只是当前元素的副本，它存储的是最后一次迭代中t的值，即为c
//		//由于defer会延迟执行函数直到包含它的函数执行完毕，因此在循环结束时，t指向的是最后一个Test实例，即为c
//		//fmt.Println(t.name)
//		//fmt.Printf("stu:%p\n", &t)
//		//defer t.Close()
//
//		//将t赋值给t2，每次迭代都会创建一个新的变量t2
//		//由于t2是一个新的变量，所以在循环结束时，每个Close方法实际上引用的是不同的t2实例，分别是{"c"},{"b"},{"a"}
//		t2 := t
//		fmt.Printf("stu:%p\n", &t2)
//		defer t2.Close()
//	}
//}

//defer f.close
type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, "closed")
}

//defer与return
func foo() (i int) {
	i = 0
	defer func() {
		fmt.Println(i)
	}()

	//具名返回值，执行return 2的时候已经将i的值重新赋值为2，所以defer函数中输出的是2而不是0
	//显示return返回前，会先修改命名返回参数
	//执行顺序 (i=i+2)->(call defer)->return
	return 2
}
