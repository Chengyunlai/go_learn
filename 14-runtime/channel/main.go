package main

import (
	"fmt"
	"sync"
)

/*
*
Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过【通信】共享内存而不是通过共享内存而实现通信。
如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。
*/

var wg sync.WaitGroup

func main() {
	var ch1 chan int   // 声明一个传递整型的通道
	var ch2 chan bool  // 声明一个传递布尔型的通道
	var ch3 chan []int // 声明一个传递int切片的通道

	fmt.Println(ch1) // nil 空指针
	fmt.Println(ch2) // nil 空指针
	fmt.Println(ch3) // nil 空指针

	// 需要使用make对通道进行初始化
	ch1 = make(chan int)
	ch2 = make(chan bool)
	ch3 = make(chan []int)

	fmt.Println(ch1) // 地址值
	fmt.Println(ch2) // 地址值
	fmt.Println(ch3) // 地址值

	fmt.Println("操作通道")

	wg.Add(1)
	go func() {
		x := <-ch1
		fmt.Println("正在读取内容")
		fmt.Println(x)
		fmt.Println("读取内容结束")
		close(ch1)
		defer wg.Done()
	}()

	ch1 <- 10
	wg.Wait()
}
