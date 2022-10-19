package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var chan1 chan int = make(chan int)
	i := 0

	go func() {
		for {
			v, ok := <-chan1
			fmt.Println(v)
			fmt.Println(ok)
		}
	}()

	for {
		chan1 <- i
		i++
		// 切换资源
		if i == 10 {
			break
		}
	}
	time.Sleep(time.Second * 5)
}
