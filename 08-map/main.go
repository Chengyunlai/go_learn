package main

import "fmt"

func main() {
	//定义一个map,这时候没有初始化是不能直接使用的
	var province map[string]string

	//使用make进行创建，初始指定容量为2
	province = make(map[string]string, 2)

	// 创建方式二：new返回的是一个指针
	//pointMap := new(map[string]string)
	//province = *pointMap

	// 手动存值
	province["ZheJiang"] = "浙江"
	province["ShangHai"] = "上海"
	province["BeiJing"] = "北京"

	fmt.Println(province)             //map[BeiJing:北京 ShangHai:上海 ZheJiang:浙江]
	fmt.Println(province["ZheJiang"]) //浙江
	fmt.Println(province["ShangHai"]) //上海
	fmt.Println(province["BeiJing"])  //北京

	// 这里我们手动给它初始化了，也是可以的
	province_ := map[string]string{"JiangSu": "江苏", "SiChuan": "四川"}
	fmt.Println(province_)

	// 判断Map里是否包含...
	_, ok := province["ZheJiang"] //_只是一个占位符，没有其他作用，表示这个_值我们是不关心的
	fmt.Println(ok)               //true
	value, ok := province["ZheJiang"]
	fmt.Println(value)
	fmt.Println(ok)

	fmt.Println("遍历Map")
	// 遍历Map
	for key, value := range province {
		fmt.Println(key, value)
	}
	// 遍历Key
	for key, _ := range province {
		fmt.Println(key)
	}
	// 遍历value
	for _, value := range province {
		fmt.Println(value)
	}
	// 使用内置函数delete，通过key删除map元素
	delete(province, "ShangHai")
	fmt.Println(province)

	// Map本身是无序的
	sortMap := make(map[int]string, 10)
	fmt.Println(sortMap) //map[]

}
