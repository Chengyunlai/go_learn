package main

import "fmt"

type Person struct {
	name string
	dog  Dog
	cat  Cat
}

type Animal struct {
	name string
	age  int
}

type Dog struct {
	//name string
	//age  int
	Animal
}

type Cat struct {
	//name string
	//age  int
	Animal
}

func main() {
	//person := Person{name: "老程", dog: Dog{name: "小黑", age: 3}, cat: Cat{name: "发财", age: 5}}
	//fmt.Println(person) //{老程 {小黑 3} {发财 5}}
	person := Person{name: "老程", dog: Dog{Animal{name: "小黑", age: 3}}, cat: Cat{Animal{name: "发财", age: 5}}}
	fmt.Println(person)          //{老程 {{小黑 3}} {{发财 5}}}
	fmt.Println(person.dog.name) //小黑

	person.feedDog(person.dog)
	person.feedCat(person.cat)
	person.feed(person.dog)
	person.feed(person.cat)
}

func (*Person) feedCat(cat Cat) {
	fmt.Println(cat.name)
	fmt.Println("吃东西")
}

func (*Person) feedDog(dog Dog) {
	fmt.Println(dog.name)
	fmt.Println("吃东西")
}

func (*Person) feed(Animal interface{}) {
	switch v := Animal.(type) {
	case Dog:
		fmt.Println(v.name)
		fmt.Println("吃骨头")
	case Cat:
		fmt.Println(v.name)
		fmt.Println("吃鱼")
	}
}
