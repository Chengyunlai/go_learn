package main

import "fmt"

func main() {
	type person struct {
		name    string //描述姓名
		age     int    //描述年龄
		setName func(p *person, name string, age int)
	}
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
}
