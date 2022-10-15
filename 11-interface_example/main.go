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

type Eater interface {
	eat()
}

func main() {
	//person := Person{name: "老程", dog: Dog{name: "小黑", age: 3}, cat: Cat{name: "发财", age: 5}}
	//fmt.Println(person) //{老程 {小黑 3} {发财 5}}
	person := Person{name: "老程", dog: Dog{Animal{name: "小黑", age: 3}}, cat: Cat{Animal{name: "发财", age: 5}}}
	fmt.Println(person)          //{老程 {{小黑 3}} {{发财 5}}}
	fmt.Println(person.dog.name) //小黑

	person.feedDog(person.dog)
	person.feedCat(person.cat)
	var x Eater
	x = person.dog
	person.feed(x)
	x = person.cat
	person.feed(x)

	//var x interface{}
	//x = person.dog
	//fmt.Println(x)

}

func (Cat) eat() {
	fmt.Println("吃鱼")
}

func (Dog) eat() {
	fmt.Println("吃骨头")
}

func (*Person) feedCat(cat Cat) {
	fmt.Println(cat.name)
	//fmt.Println("吃东西")
	cat.eat()
}

func (*Person) feedDog(dog Dog) {
	fmt.Println(dog.name)
	//fmt.Println("吃东西")
	dog.eat()
}

func (*Person) feed(Animal Eater) {
	//switch v := Animal.(type) {
	//case Dog:
	//	fmt.Println("--")
	//	fmt.Println(v)
	//	fmt.Println(v.name)
	//	//fmt.Println("吃骨头")
	//	v.eat()
	//case Cat:
	//	fmt.Println("--")
	//	fmt.Println(v)
	//	fmt.Println(v.name)
	//	//fmt.Println("吃鱼")
	//	v.eat()
	//}
	Animal.eat()
}
