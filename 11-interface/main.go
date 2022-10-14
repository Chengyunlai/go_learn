package main

import "fmt"

type Dog struct {
	name string
}

type Cat struct {
	name string
}

// 定义一个Mover接口，里面设置了一个move方法
type Mover interface {
	move()
}

func (cat Cat) move() {
	fmt.Println("猫移动了")
}

//func (dog Dog) move() {
//	fmt.Println("狗移动了")
//}

// 此时实现Mover接口的是*dog类型，所以不能给x传入dog类型的wangcai，此时x只能存储*dog类型的值。
func (dog *Dog) move() {
	fmt.Println("狗会动")
}

func main() {
	var animal Mover
	cat := Cat{}
	animal = cat
	animal.move()
	dog := Dog{}
	animal = &dog
	animal.move()

	fmt.Println("---")
	var x Mover
	var wangcai = Dog{} // 旺财是dog类型
	x = &wangcai        // x可以接收dog类型
	x.move()
	var fugui = &Dog{} // 富贵是*dog类型
	x = fugui          // x可以接收*dog类型
	x.move()

	//使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。
	//因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui。

	//总结:
	//若某个类型实现接口的方法中定义的接收者如果是指类型的话，那么使用多态思想时直接传入类型和指针类型都可以实现

	//若某个类型实现接口的方法是指针接收者,那么使用多态思想时直接传入的类型只能是指针类型

	fmt.Println("--空接口--")
	// 定义一个空接口x
	var x2 interface{}
	fmt.Println(x2)
	s := "pprof.cn"
	x2 = s
	fmt.Printf("type:%T value:%v\n", x2, x2) //type:string value:pprof.cn
	i := 100
	x2 = i
	fmt.Printf("type:%T value:%v\n", x2, x2) //type:string value:pprof.cn
	b := true
	x2 = b
	fmt.Printf("type:%T value:%v\n", x2, x2) //type:string value:pprof.cn

	// 断言
	b2, ok := x2.(string)
	//b2, ok := x2.(string)

	fmt.Println(b2)
	fmt.Println(ok)
}
