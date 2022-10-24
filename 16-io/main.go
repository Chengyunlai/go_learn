package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 注意：这里默认指的路径是指整个工程的根目录，见password的位置
	// 打开只读文件

	fmt.Println("打开文件", os.O_RDWR)
	open, err := os.Open("./password.txt") // 返回两个值，第一个是文件对象的指针*File，第二个是判断文件是否存在，如果文件不存在则报错 err != nil
	if err != nil {
		fmt.Println("文件打开失败，请查询原因")
		fmt.Println(err)
	}

	fmt.Println(open)
	defer open.Close()

	// 操作文件：读取内容

	//var buf [64]byte // 注意这里一定要指定数组
	//n, err := open.Read(buf[:])
	//if err != nil {
	//	fmt.Println("文件读取失败，请查询原因")
	//	fmt.Println(err)
	//}

	// 读文件使用这种思路
	//var content []byte
	//for {
	//	n, err := open.Read(buf[:])
	//	if err == io.EOF {
	//		// 读取结束
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("read file err ", err)
	//		return
	//	}
	//	content = append(content, buf[:n]...)
	//}
	//fmt.Println(string(content))
	//read(open) // 读文件是按照文件指针的。

	// 操作文件：写内容
	//n, err := open.WriteString("123") //无法写入，因为这是个只读文件
	//fmt.Println(n)
	//fmt.Println(err) //write ./password.txt: Access is denied.
	read(open)

	file, err2 := os.OpenFile("./password.txt", os.O_RDONLY, 0777)
	fmt.Println(file)
	fmt.Println(err2)
	file.WriteString("aaa")

}

func read(open *os.File) {
	//file := copy(*open)
	fmt.Println("读取文件")
	var buf [64]byte
	var content []byte
	for {
		n, err := open.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))

	// 读写文件
	//os.OpenFile("./password")
}

//func copy(dest os.File) os.File {
//	temOpen := new(os.File)
//	*temOpen = dest
//	return *temOpen
//}
