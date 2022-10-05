package main

import(
	"fmt"
)
	


func main() {
	// 定义一个数组(定长)，不给出值的情况下，默认根据类型的初始化
	var array [2]int
	fmt.Println(array) //[0 0]

	// 默认下标从0开始
	array[0] = 1
	fmt.Println(array) //[1 0]
	fmt.Println(array[1]) //0

	// 根据给的值定义数组的长度
	var array2 []int = []int{1}
	fmt.Println(array2) //[1]

	// 重要的是声明后面的值类型是个数组，不然程序会觉得它是个方法。
	var array3 = []int{1,2}
	fmt.Println(array3) //[1 2]

	// 当然我们可以使用简易定义的方式 :=
	array4 := []int{3,4}
	fmt.Println(array4) //[3 4]

	// 获取数组的长度
	fmt.Println(len(array3)) //2

	// 特殊用法，跳跃式指定值
	array5 := [3]int{0:0,2:2}
	fmt.Println(array5) //[0,0,2]

	// 遍历数组-for循环
	for i := 0; i < len(array5); i++ {
    fmt.Println(array5[i])
		// 0
		// 0
		// 2
  }

	for index,value := range array5{
		fmt.Println(index,value)
		// 0 0
		// 1 0
		// 2 2
	}

	// 切片

}