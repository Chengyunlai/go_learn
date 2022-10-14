package main

import (
	"fmt"
)

// 定义结构体
type Person struct {
	name string
	age  int
}

func main() {
	// 定义方法一:通过var的方式定义并初始化
	var people Person = Person{name: "var", age: 1}
	fmt.Println(people) //{ 0}

	// 定义方法2:通过 := 的方式，
	people2 := Person{}
	fmt.Println(people2)

	// 定义方式2:通过 := 的方式并赋值
	p0 := Person{
		"张三",
		18,
	}
	fmt.Println(p0)

	p1 := new(Person)
	fmt.Println(p1)

	p2 := newPerson("Chengyunlai", 18)
	fmt.Println(p2)

}

// 使用函数来实例化结构体
func newPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}
