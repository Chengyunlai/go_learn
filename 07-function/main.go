package main

import (
	"fmt"
)

func main() {
	//fmt.Println("你好" + 5)
	fmt.Println("--匿名函数--")
	//声明一个变量f  类型是一个函数类型
	var f func()
	//将自定义函数myfunc 赋给变量f,记住go语言是包管理的
	f = myfunc
	//调用变量f 相当于调用函数myfunc()
	f()

	fmt.Println("--匿名函数作为参数传递--")
	res2 := oper(20, 12, add)
	fmt.Println(add)
	fmt.Println(res2)

	// 匿名函数作为回调函数直接写入参数中，这里是直接将一个匿名函数传过去
	res3 := oper(2, 4, func(a, b int) int {
		return a + b
	})
	fmt.Println(res3)

	fmt.Println("--匿名函数作为返回值传递--")

	res0 := closure()
	fmt.Println(res0) //0x49a880  返回内层函数函数体地址
	r10 := res0()     //执行closure函数返回的匿名函数
	fmt.Println(r10)  //1
	r20 := res0()
	fmt.Println(r20) //2
	//普通的函数应该返回1，而这里存在闭包结构所以返回2 。
	//一个外层函数当中有内层函数，这个内层函数会操作外层函数的局部变量,并且外层函数把内层函数作为返回值,则这里内层函数和外层函数的局部变量,统称为闭包结构。这个外层函数的局部变量的生命周期会发生改变，不会随着外层函数的结束而销毁。
	//所以上面打印的r2 是累计到2 。

	res20 := closure() //再次调用则产生新的闭包结构 局部变量则新定义的
	fmt.Println(res20)
	r30 := res20()
	fmt.Println(r30)

	fmt.Println("--指针--")

	//定义一个变量
	a := 2
	fmt.Printf("变量A的地址为%p", &a) //通过%p占位符, &符号获取变量的内存地址。
	//变量A的地址为0xc000072090

	//创建一个指针
	// 指针的声明 通过 *T 表示T类型的指针
	var i *int     //int类型的指针
	fmt.Println(i) // < nil >空指针
	fmt.Println(f)

	//因为指针存储的变量的地址 所以指针存储值
	i = &a
	fmt.Println(i)  //i存储a的内存地址0xc000072090
	fmt.Println(*i) //i存储这个指针存储的变量的数值2
	*i = 100        // 此时因为i存放的是a的地址，这时候对i操作就是对a操作，所以*i = 100 的时候，a的值就是100
	fmt.Println(*i) //100
	fmt.Println(a)  //100通过指针操作 直接操作的是指针所对应的数值

	fmt.Println("--函数作为参数--")
	res := alu(10.0, 20.0, mult)
	fmt.Println(res)

	res = func(a int, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(res) //打印匿名函数返回值

	res = alu(10, 20, func(a int, b int) int {
		return a * b
	})
	fmt.Println(res)

	speed := fast_shoes("张三")
	for i := 0; i < 10; i++ {
		fmt.Println(speed(10)) // 20
	}
	// fmt.Println(speed(10)) // 20
	// fmt.Println(speed(10))
	// fmt.Println(speed(10))
	// fmt.Println(speed(10))

}

// 定义一个闭包结构的函数 返回一个匿名函数,匿名函数作为返回值的时候，如果用到了外部的局部变量，当返回值还在使用的时候，局部变量不会被释放
func closure() func() int { //外层函数
	//定义局部变量a
	a := 0 //外层函数的局部变量
	//定义内层匿名函数 并直接返回
	return func() int { //内层函数
		a++ //在匿名函数中将变量自增。内层函数用到了外层函数的局部变量，此变量不会随着外层函数的结束销毁
		return a
	}
}

// 定义的是一个普通的函数
func add(a, b int) int {
	return a + b
}

// oper就叫做高阶函数
// fun 函数作为参数传递则fun在这里叫做回调函数
func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun) //20 12 0x49a810A   第三个打印的是传入的函数体内存地址
	res := fun(a, b)       //fun 在这里作为回调函数 程序执行到此之后才完成调用
	return res

}

// 自定义函数
func myfunc() {
	fmt.Println("myfunc...")

	// 匿名函数
	func() {
		fmt.Println("匿名函数")
	}()

	func(a int, b int) {
		fmt.Println(a, b)
	}(1, 2) //括号内为匿名函数的实参

	res := func(a int, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(res) //打印匿名函数返回值

	ress := func() float64 {
		return 10.0
	}() //输出10.0

	fmt.Println(ress) //打印10.0

	resss := func() float32 {
		return 11.0
	} // 只是定义了一个函数

	fmt.Println(resss()) //11.0
}

// 首先定义一个函数，准备将它作为参数传递。该函数是一个乘法函数，接收两个float的数，然后将他们相乘并返回
func mult(a, b int) int {
	return a * b
}

// 定义一个函数用来接收函数，并调用这个函数（哈哈套娃行为）
func alu(a, b int, fun func(int, int) int) int {
	res := fun(a, b)
	return res
}

func fast_shoes(name string) func(_speed int) int {
	fmt.Println(name + "获得飞速鞋子一双")
	speed := 10
	return func(_speed int) int {
		_speed += speed
		speed++
		return _speed
	}
}
