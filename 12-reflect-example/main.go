package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Dog struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// 这里例子是模拟最底层的封装结构体，假设前端传入了一段json字符串，我们需要将它转为对应的struct
func main() {
	personJson := `{"name":"程云来","age":23}`
	dogJson := `{"name":"小黑","color":"黑棕色"}`

	//var person Person
	//var dog Dog
	//
	//if err := json.Unmarshal([]byte(personJson), &person); err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(person)
	//
	//if err := json.Unmarshal([]byte(dogJson), &dog); err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(dog)

	person := Person{}.toStruct(personJson)
	dog := Dog{}.toStruct(dogJson)
	fmt.Println(person)
	fmt.Println(dog)

}

// 发现封装json的代码重复性太高了，进行改造，我们可以考虑封装一个结构体的方法，通过它接收一个json字符串后然后对其进行封装。
// 1.定义接口
type JsonToStruct interface {
	toStruct()
}

// 2.实现接口，实现接口中应该将具体的封装操作通过反射的方式抽取出来，避免重复的代码
func (Person) toStruct(jsonStr string) Person {
	instance := newInstance(Person{}, jsonStr)
	person := instance.(*Person)
	return *person
}

func (Dog) toStruct(jsonStr string) Dog {
	instance := newInstance(Dog{}, jsonStr)
	dog := instance.(*Dog)
	return *dog
}

// 3.将封装的语句抽取封装
func newInstance(x interface{}, jsonStr string) interface{} {
	tp := reflect.TypeOf(x)
	instance := reflect.New(tp).Interface()
	if err := json.Unmarshal([]byte(jsonStr), &instance); err != nil {
		fmt.Println(err)
	}
	return instance
}

// 4.优化
type GetType interface {
	toStruct()
}
