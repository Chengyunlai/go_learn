package main

import "fmt"

type person struct {
	name    string //描述姓名
	age     int    //描述年龄
	setName func(p *person, name string, age int)
}

func main() {

	// 定义1
	var people person
	fmt.Println(people) //{ 0}

	//通过类型.字段名的方式结构体的字段
	people.name = "chengyunlai"
	people.age = 18
	fmt.Println(people) // {chengyunlai 18}

	// 定义2 使用new返回指针的方式定义
	people2 := new(person)
	fmt.Println(people2)        //&{ 0}
	fmt.Println(*(people2))     //{ 0}
	fmt.Println((*people2).age) // 0
	(*people2).age = 18
	fmt.Println(people2.age)

	people.setName = func(p *person, name_ string, age_ int) {
		p.name = name_
		p.age = age_
		fmt.Printf("%p\n", p)
	}

	people.setName(&people, "张三", 15)
	fmt.Printf("%p\n", &people)
	fmt.Println(people)

	people_ := &people
	people_.setter("程云来", 18)
	fmt.Println(people_)

	people.setter("程云来", 18)
	fmt.Println(people)

	people3 := person{name: "chengyunlai", age: 21}
	fmt.Println(people3)

	people4 := new(person)
	people4.name = "张三a"
	fmt.Println(people4)

	p2 := person{name: "chengyunlai1", age: 221}
	//1 使用结构体指针
	var p *person
	p = &p2 //将p2 的地址赋给p
	fmt.Println(p)
	p.name = "好家伙" //修改p的值
	fmt.Println(p)
	fmt.Println(p2) //p2的值也被修改了
}

func (people *person) setter(name string, age int) {
	people.name = name
	people.age = age
}
