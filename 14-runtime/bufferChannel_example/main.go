package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var chan1 chan int = make(chan int, 10)
	i := 0

	wg.Add(1)
	go func() {
		//for {
		//	v, ok := <-chan1
		//	if !ok {
		//		break
		//	} else {
		//		fmt.Println(v)
		//	}
		//}
		for v := range chan1 { // 优雅的从通道循环取值
			fmt.Println(v)
		}
		defer wg.Done()
	}()

	for {
		if i == 10 {
			// 停止往通道发内容,原因是退出之后子程序还没接收完毕
			close(chan1)
			break
		} else {
			chan1 <- i
			i++
		}
	}

	wg.Wait()
}
