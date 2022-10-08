package main

import (
	"fmt"
	"strings" // strings包，用于字符串查找、替换、比较等
)

func main() {
	// 关于字符串的操作
	useString()
}

func useString() {
	// 定义一个字符串：chengyunlai
	str := "chengyunlai"

	// 查找字符串
	println(strings.Contains(str, "cheng")) //true
	println(strings.Contains(str, ""))      //true
	println(strings.Contains(str, "x"))     // false

	println(strings.Index(str, "y")) //5
	println(strings.Index(str, "x")) //-1

	println(strings.ContainsAny(str, "zxcvbnm")) //true
	println(strings.ContainsAny(str, "qwrt"))    //false

	println(strings.ContainsRune(str, 97)) //true
	println(strings.ContainsRune(str, 98)) //false

	println(strings.Count(str, ""))  //12
	println(strings.Count(str, "n")) //2

	println(strings.LastIndex(str, "n"))

	println(strings.Cut(str, "yun")) //cheng lai true

	println(strings.Compare(str, "ahengyunlai")) //1
	println(strings.Compare(str, "chengyunlai")) //0
	println(strings.Compare(str, "dhengyunlai")) //-1

	fmt.Println("--")
	str2 := "cheng_yun_lai"
	split := strings.Split(str2, "_")
	for _, str := range split {
		fmt.Println(str)
	}
	/**
	cheng
	yun
	lai
	*/
}
