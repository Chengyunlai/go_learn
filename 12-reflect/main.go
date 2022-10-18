package main

import (
	"fmt"
	"reflect"
)

type mystruct struct {
	Name string
	Sex  int
	Age  int `json:"age"`
}

func main() {
	a := 36
	atype := reflect.TypeOf(a)
	fmt.Println(atype)        // 输出类型
	fmt.Println(atype.Name()) //获取类型名称为int
	fmt.Println(atype.Kind()) //获取归属的类型
	avalue := reflect.ValueOf(a)
	fmt.Println(avalue.Int()) //获取具体的数值

	fmt.Println("结构体")
	x := mystruct{}
	x.Name = "chengyunlai"
	typeofmystruct := reflect.TypeOf(x)
	fmt.Println(reflect.ValueOf(x))
	fmt.Println(typeofmystruct)
	fmt.Println(typeofmystruct.Name()) //获取反射类型对象  mystruct
	fmt.Println(typeofmystruct.Kind()) //获取反射类型种类  struct

	pointTypeofmystruct := reflect.TypeOf(&x)
	fmt.Println(pointTypeofmystruct)
	fmt.Println(pointTypeofmystruct.Elem().Name())
	fmt.Println(pointTypeofmystruct.Elem().Kind())
	fmt.Println("--")
	fmt.Println(reflect.ValueOf(x))

	fieldnum := typeofmystruct.NumField() //获取结构体成员字段的数量

	fmt.Println("--")
	for i := 0; i < fieldnum; i++ {
		fieldname := typeofmystruct.Field(i) //索引对应的字段信息
		fmt.Println(fieldname)
	}

	fmt.Println("--")
	for i := 0; i < fieldnum; i++ {
		name, err := typeofmystruct.FieldByName("Name")
		fmt.Println(name)

		fmt.Println(err)
	}

	fmt.Println("---")
	h := haojiahuo{"好家伙", 20}
	fmt.Println(h)
	hofvalue := reflect.ValueOf(h)             //获取结构体的reflect.Value对象。
	for i := 0; i < hofvalue.NumField(); i++ { //循环结构体内字段的数量
		//获取结构体内索引为i的字段值
		fmt.Println(hofvalue.Field(i).Interface())
	}
	fmt.Println(hofvalue.Field(1).Interface()) //获取结构体内索引为1的字段的类型
}

type haojiahuo struct {
	Name string
	Age  int
}
